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

setup_environment
show_config "full"

#:[.'.]:>-------------------------------------
show_banner

rm -rf $HOME/.protoc
rm -f $(go env GOPATH)/bin/protoc-gen-go
rm -f $(go env GOPATH)/bin/protoc-gen-go-grpc
rm -f $HOME/go/bin/protoc-gen-go
rm -f $HOME/go/bin/protoc-gen-go-grpc
#:[.'.]:>-------------------------------------