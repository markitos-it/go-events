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

CONTAINER_NAME="goevents-mariadb"
DB_ROOT_USER="root"
DB_ROOT_PASSWORD="admin"
DB_NEW="goevents"

docker exec -i ${CONTAINER_NAME} mariadb -u ${DB_ROOT_USER} -p${DB_ROOT_PASSWORD} -e "CREATE DATABASE IF NOT EXISTS ${DB_NEW};"
docker exec -i ${CONTAINER_NAME} mariadb -u ${DB_ROOT_USER} -p${DB_ROOT_PASSWORD} -e "CREATE USER IF NOT EXISTS '${DB_NEW}'@'%' IDENTIFIED BY '${DB_NEW}';"
docker exec -i ${CONTAINER_NAME} mariadb -u ${DB_ROOT_USER} -p${DB_ROOT_PASSWORD} -e "GRANT ALL PRIVILEGES ON ${DB_NEW}.* TO '${DB_NEW}'@'%'; FLUSH PRIVILEGES;"
#:[.'.]:>-------------------------------------