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

log_info "Removing database and associated user"

CONTAINER_NAME="goevents-mariadb"
DB_ROOT_USER="root"
DB_ROOT_PASSWORD="admin"
DATABASE_NEW_SERVICE="goevents"

docker exec -i ${CONTAINER_NAME} mariadb -u ${DB_ROOT_USER} -p${DB_ROOT_PASSWORD} -e "DROP DATABASE IF EXISTS ${DATABASE_NEW_SERVICE};"
docker exec -i ${CONTAINER_NAME} mariadb -u ${DB_ROOT_USER} -p${DB_ROOT_PASSWORD} -e "DROP USER IF EXISTS '${DATABASE_NEW_SERVICE}'@'%';"

log_info "Removal process completed"
#:[.'.]:>-------------------------------------