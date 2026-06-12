#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'
echo "🔧 Compiling proto files from $(pwd)"
ENVIRONMENT_FILE="bin/shared/environment.sh"
source $ENVIRONMENT_FILE

function log_info() {
    echo "[INFO] $*"
}
function log_error() {
    echo "[ERROR] $*" >&2
}

setup_environment
show_config "full"

#:[.'.]:>-------------------------------------
show_banner

echo "['.']:> 🔨 Building..."
mkdir -p dist
rm -f dist/govent
go build -o dist/govent cmd/app/main.go
echo "['.']:> ✅ Binario generado en dist/govent"
#:[.'.]:>-------------------------------------

log_info "Compiled successfully."