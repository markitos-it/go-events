#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
ENVIRONMENT_FILE="bin/shared/environment.sh"
source "$ENVIRONMENT_FILE"

# Colores para la consola
GREEN='\033[0;32m'
RED='\033[0;31m'
RESET='\033[0;0m'

log_info() { echo -e "[INFO] $*"; }
log_error() { echo -e "[ERROR] $*" >&2; }

# Función auxiliar para pintar OK/KO
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
log_info "Iniciando análisis rápido..."
echo "--------------------------------------------------"

GLOBAL_EXIT=0

# 1. Validación de Gitleaks instalado
if ! gitleaks version > /dev/null 2>&1; then
    print_result "Gitleaks (Instalación)" 1 && GLOBAL_EXIT=1
    log_error "Gitleaks no está instalado en el sistema."
fi

# 2. Ejecución de GolangCI-Lint
if golangci-lint run > /dev/null 2>&1; then
    print_result "GolangCI-Lint" 0
else
    print_result "GolangCI-Lint" 1 && GLOBAL_EXIT=1
    log_error "Errores de linter detectados en Go. Ejecuta 'golangci-lint run' para verlos."
fi

# 3. Ejecución de Gitleaks
# Nota: Si este script corre en el pre-commit, recuerda cambiar 'detect --source .' por 'protect --staged' para que vuele.
if gitleaks detect --source . --no-git > /dev/null 2>&1; then
    print_result "Gitleaks (Secretos)" 0
else
    print_result "Gitleaks (Secretos)" 1 && GLOBAL_EXIT=1
    log_error "Gitleaks detectó posibles credenciales o tokens expuestos."
fi

echo "--------------------------------------------------"
if [[ $GLOBAL_EXIT -eq 0 ]]; then
    log_info "Análisis finalizado con éxito."
    exit 0
else
    log_error "El análisis terminó con errores. Revisa los [KO] anteriores."
    exit 1
fi