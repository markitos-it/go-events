#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
ENVIRONMENT_FILE="bin/shared/environment.sh"
source "$ENVIRONMENT_FILE"

# Colores para la consola
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0;3m' # No Color
RESET='\033[0;0m'

log_info() { echo -e "[INFO] $*"; }
log_error() { echo -e "[ERROR] $*" >&2; }

# Función auxiliar para formatear el resultado
print_result() {
    local name=$1
    local status=$2
    if [[ $status -eq 0 ]]; then
        echo -e "  - $name: ${GREEN}[OK]${RESET}"
    else
        echo -e "  - $name: ${RED}[KO]${RESET}"
        return 1
    fi
}

setup_environment
show_config "full"

echo "--------------------------------------------------"
log_info "Iniciando análisis de seguridad..."
echo "--------------------------------------------------"

# 1. Verificaciones rápidas de herramientas
snyk auth "$SNYK_TOKEN" > /dev/null 2>&1

# Variable para rastrear si algo falla al final
GLOBAL_EXIT=0

# 2. Ejecución de Scans (Silenciosos a menos que fallen)

# --- SNYK CODE ---
if snyk code test --severity-threshold=medium --include-ignores > /dev/null 2>&1; then
    print_result "Snyk Code (SAST)" 0
else
    print_result "Snyk Code (SAST)" 1 && GLOBAL_EXIT=1
    log_error "Snyk Code encontró vulnerabilidades de severidad Media/Alta."
fi

# --- SNYK SCA ---
if snyk test --all-projects --severity-threshold=medium --include-ignores > /dev/null 2>&1; then
    print_result "Snyk SCA (Dependencias)" 0
else
    print_result "Snyk SCA (Dependencias)" 1 && GLOBAL_EXIT=1
    log_error "Snyk SCA encontró vulnerabilidades en tus librerías."
fi

# --- SNYK IAC ---
if IAC_OUTPUT=$(snyk iac test --severity-threshold=high 2>&1); then
    print_result "Snyk IaC" 0
else
    if echo "$IAC_OUTPUT" | grep -qE "Could not find any valid IaC files|SNYK-CLI-0012|monthly limit"; then
        echo -e "  - Snyk IaC: [OMITIDO] (No hay archivos o límite alcanzado)"
    else
        print_result "Snyk IaC" 1 && GLOBAL_EXIT=1
        echo "$IAC_OUTPUT" >&2
        log_error "Snyk IaC encontró problemas de configuración."
    fi
fi

# --- GITLEAKS ---
# Nota: Si es para pre-commit, recuerda cambiar 'detect --source .' por 'protect --staged'
if gitleaks detect --source . --no-git > /dev/null 2>&1; then
    print_result "Gitleaks (Secretos)" 0
else
    print_result "Gitleaks (Secretos)" 1 && GLOBAL_EXIT=1
    log_error "Gitleaks detectó posibles credenciales/tokens expuestos."
fi

# --- GOLANGCI-LINT ---
if golangci-lint run > /dev/null 2>&1; then
    print_result "GolangCI-Lint" 0
else
    print_result "GolangCI-Lint" 1 && GLOBAL_EXIT=1
    log_error "Errores de linter detectados en Go. Ejecuta 'golangci-lint run' para verlos."
fi

echo "--------------------------------------------------"
if [[ $GLOBAL_EXIT -eq 0 ]]; then
    log_info "Análisis finalizado con éxito."
    exit 0
else
    log_error "El análisis terminó con errores. Revisa los [KO] anteriores."
    exit 1
fi