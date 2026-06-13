.DEFAULT_GOAL := help
.PHONY: 
	help \
	commits-help \
	start \
	test \
	test-e2e \
	proto \
	db-start \
	db-stop \
	db-create \
	db-drop \
	db-seed \
	spanner-start \
	spanner-stop \
	spanner-create \
	spanner-drop \
	mariadb-start \
	mariadb-stop \
	mariadb-create \
	mariadb-drop \
	mariadb-seed \
	format \
	lint \
	support-install-linter \
	support-uninstall-linter \
	build \
	appsec-install \
	appsec-uninstall \
	appsec-test \
	tidy

help:
	@echo ""
	@echo "['.']:> =================================================="
	@echo "['.']:> 🚀 MARKITOS-IT GO VENTS COMMANDS"
	@echo "['.']:> =================================================="
	@printf "  \033[36m%-24s\033[0m %s\n" "help" "Muestra este menú de ayuda interactivo"
	@printf "  \033[36m%-24s\033[0m %s\n" "commits-help" "Muestra ayuda para commits"
	@printf "  \033[36m%-24s\033[0m %s\n" "build" "Genera el binario final de la aplicación en la carpeta dist"
	@printf "  \033[36m%-24s\033[0m %s\n" "start" "Inicia la aplicación localmente"
	@printf "  \033[36m%-24s\033[0m %s\n" "test" "Ejecuta la suite de pruebas de la aplicación"
	@printf "  \033[36m%-24s\033[0m %s\n" "test-e2e" "Ejecuta las pruebas End-to-End mediante gRPC"
	@printf "  \033[36m%-24s\033[0m %s\n" "proto" "Genera los archivos de código a partir de los de gRPC .proto"
	@printf "  \033[36m%-24s\033[0m %s\n" "lint" "Analiza el código Go con golangci-lint"
	@printf "  \033[36m%-24s\033[0m %s\n" "lint-fix" "Formatea el código Go automáticamente (gofmt, goimports)"
	@printf "  \033[36m%-24s\033[0m %s\n" "tidy" "Limpia y actualiza las dependencias de Go (go mod tidy)"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-create" "Crea la base de datos en PostgreSQL"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-drop" "Elimina (drop) la base de datos por completo"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-start" "Inicia el entorno de la base de datos"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-stop" "Detiene el entorno de la base de datos"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-seed" "Puebla la base de datos con datos de prueba a través de gRPC"
	@printf "  \033[36m%-24s\033[0m %s\n" "spanner-start" "Inicia el entorno del emulador de Spanner"
	@printf "  \033[36m%-24s\033[0m %s\n" "spanner-stop" "Detiene el entorno del emulador de Spanner"
	@printf "  \033[36m%-24s\033[0m %s\n" "spanner-create" "Crea la base de datos en Spanner"
	@printf "  \033[36m%-24s\033[0m %s\n" "spanner-drop" "Elimina (drop) la base de datos en Spanner por completo"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-start" "Inicia el entorno de la base de datos MariaDB"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-stop" "Detiene el entorno de la base de datos MariaDB"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-create" "Crea la base de datos en MariaDB"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-drop" "Elimina (drop) la base de datos en MariaDB por completo"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-seed" "Puebla la base de datos MariaDB con datos de prueba"
	@printf "  \033[36m%-24s\033[0m %s\n" "support-install-linter" "Instala la herramienta golangci-lint"
	@printf "  \033[36m%-24s\033[0m %s\n" "support-uninstall-linter" "Desinstala la herramienta golangci-lint"
	@printf "  \033[36m%-24s\033[0m %s\n" "appsec-install" "Instala herramientas de seguridad (Snyk, Gitleaks)"
	@printf "  \033[36m%-24s\033[0m %s\n" "appsec-uninstall" "Desinstala las herramientas de seguridad"
	@printf "  \033[36m%-24s\033[0m %s\n" "appsec-test" "Ejecuta las pruebas de seguridad (Snyk, Gitleaks)"
	@echo "['.']:> =================================================="
	@echo ""

commits-help:
	@echo ""
	@echo "📚 \033[1mGuía rápida de Conventional Commits!:\033[0m"
	@printf "  \033[36m%-10s\033[0m %s\n" "feat:" "Nueva característica (Sube versión MINOR)"
	@printf "  \033[36m%-10s\033[0m %s\n" "fix:" "Corrección de un error (Sube versión PATCH)"
	@printf "  \033[36m%-10s\033[0m %s\n" "docs:" "Cambios en la documentación"
	@printf "  \033[36m%-10s\033[0m %s\n" "chore:" "Tareas de mantenimiento o herramientas (ej. Docker, Makefile)"
	@printf "  \033[36m%-10s\033[0m %s\n" "refactor:" "Cambios en el código que no fijan errores ni añaden funciones"
	@echo "  --- (Usa '!' para Breaking Changes, ej: feat!:) ---"
	@printf "  \033[36m%-10s\033[0m %s\n" "feat!:" "Nueva característica con cambios importantes"
	@printf "  \033[36m%-10s\033[0m %s\n" "fix!:" "Corrección de error con cambios importantes"
	@printf "  \033[36m%-10s\033[0m %s\n" "docs!:" "Documentación con cambios importantes"
	@printf "  \033[36m%-10s\033[0m %s\n" "chore!:" "Mantenimiento con cambios importantes"
	@printf "  \033[36m%-10s\033[0m %s\n" "refactor!:" "Refactorización con cambios importantes"
	@echo ""

start:
	bash bin/app/start.sh

test:
	bash bin/app/test.sh

test-e2e:
	bash bin/app/test_e2e_grpc.sh

test-verbose:
	bash bin/app/test-verbose.sh

proto:
	bash bin/app/proto.sh

db-start:
	bash bin/database/postgres/start.sh

db-stop:
	bash bin/database/postgres/stop.sh

db-create:
	bash bin/database/postgres/create.sh

db-drop:
	bash bin/database/postgres/drop.sh

db-seed:
	bash bin/database/postgres/seed.sh

mariadb-start:
	bash bin/database/mariadb/start.sh

mariadb-stop:
	bash bin/database/mariadb/stop.sh

mariadb-create:
	bash bin/database/mariadb/create.sh

mariadb-drop:
	bash bin/database/mariadb/drop.sh

mariadb-seed:
	bash bin/database/mariadb/seed.sh

spanner-start:
	bash bin/database/spanner/start.sh

spanner-stop:
	bash bin/database/spanner/stop.sh

spanner-create:
	bash bin/database/spanner/create.sh

spanner-drop:
	bash bin/database/spanner/drop.sh

spanner-seed:
	bash bin/database/spanner/seed.sh

lint:
	bash bin/code/lint.sh

lint-fix:
	bash bin/code/lint-fix.sh

support-install-linter:
	bash bin/support/install-linter.sh

support-uninstall-linter:
	bash bin/support/uninstall-linter.sh

build:
	bash bin/app/build.sh

clean:
	bash bin/app/clean.sh

appsec-test:
	bash bin/appsec/test.sh

appsec-install:
	bash bin/appsec/install.sh

appsec-uninstall:
	bash bin/appsec/uninstall.sh

appsec-pre-commit:
	bash bin/appsec/pre-commit.sh

appsec-pre-push:
	bash bin/appsec/pre-push.sh

tidy:
	go mod tidy