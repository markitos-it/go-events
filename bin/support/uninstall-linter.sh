#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
ENVIRONMENT_FILE="bin/shared/environment.sh"
source $ENVIRONMENT_FILE

function log_info() {
    echo "[INFO] $*"
}
function log_error() {
    echo "[ERROR] $*" >&2
}
echo "['.']:> =================================================="
echo "['.']:> 🗑️  DESINSTALANDO GOLANGCI-LINT"
echo "['.']:> =================================================="

setup_environment
show_config "full"

#:[.'.]:>-------------------------------------
show_banner

GOPATH_BIN="$(go env GOPATH)/bin"
if [ -f "$GOPATH_BIN/golangci-lint" ]; then
    rm -f "$GOPATH_BIN/golangci-lint"
    echo "['.']:> ✅ golangci-lint ha sido eliminado de $GOPATH_BIN"
else
    echo "['.']:> ℹ️  golangci-lint no estaba instalado en $GOPATH_BIN"
fi
#:[.'.]:>-------------------------------------