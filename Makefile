.DEFAULT_GOAL := help
.PHONY: 
	help \
	commits-help \
	start \
	test \
	test-e2e \
	test-e2e-stress \
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
	@echo "['.']:> 🚀 MARKITOS-IT GO EVENTS COMMANDS"
	@echo "['.']:> =================================================="
	@printf "  \033[36m%-24s\033[0m %s\n" "help" "Show this interactive help menu"
	@printf "  \033[36m%-24s\033[0m %s\n" "commits-help" "Show help for commits"
	@printf "  \033[36m%-24s\033[0m %s\n" "build" "Generate the final application binary in the dist folder"
	@printf "  \033[36m%-24s\033[0m %s\n" "start" "Start the application locally"
	@printf "  \033[36m%-24s\033[0m %s\n" "test" "Run the application test suite"
	@printf "  \033[36m%-24s\033[0m %s\n" "test-e2e" "Run End-to-End tests via gRPC"
	@printf "  \033[36m%-24s\033[0m %s\n" "test-e2e-stress" "Run End-to-End stress tests via gRPC"
	@printf "  \033[36m%-24s\033[0m %s\n" "proto" "Generate code files from gRPC .proto files"
	@printf "  \033[36m%-24s\033[0m %s\n" "lint" "Analyze Go code with golangci-lint"
	@printf "  \033[36m%-24s\033[0m %s\n" "lint-fix" "Format Go code automatically (gofmt, goimports)"
	@printf "  \033[36m%-24s\033[0m %s\n" "tidy" "Clean and update Go dependencies (go mod tidy)"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-create" "Create the database in PostgreSQL"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-drop" "Completely drop the database"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-start" "Start the database environment"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-stop" "Stop the database environment"
	@printf "  \033[36m%-24s\033[0m %s\n" "db-seed" "Seed the database with test data via gRPC"
	@printf "  \033[36m%-24s\033[0m %s\n" "spanner-start" "Start the Spanner emulator environment"
	@printf "  \033[36m%-24s\033[0m %s\n" "spanner-stop" "Stop the Spanner emulator environment"
	@printf "  \033[36m%-24s\033[0m %s\n" "spanner-create" "Create the database in Spanner"
	@printf "  \033[36m%-24s\033[0m %s\n" "spanner-drop" "Completely drop the database in Spanner"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-start" "Start the MariaDB database environment"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-stop" "Stop the MariaDB database environment"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-create" "Create the database in MariaDB"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-drop" "Completely drop the database in MariaDB"
	@printf "  \033[36m%-24s\033[0m %s\n" "mariadb-seed" "Seed the MariaDB database with test data"
	@printf "  \033[36m%-24s\033[0m %s\n" "support-install-linter" "Install the golangci-lint tool"
	@printf "  \033[36m%-24s\033[0m %s\n" "support-uninstall-linter" "Uninstall the golangci-lint tool"
	@printf "  \033[36m%-24s\033[0m %s\n" "appsec-install" "Install security tools (Snyk, Gitleaks)"
	@printf "  \033[36m%-24s\033[0m %s\n" "appsec-uninstall" "Uninstall security tools"
	@printf "  \033[36m%-24s\033[0m %s\n" "appsec-test" "Run security tests (Snyk, Gitleaks)"
	@echo "['.']:> =================================================="
	@echo ""

commits-help:
	@echo ""
	@echo "📚 \033[1mQuick Guide to Conventional Commits!:\033[0m"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., feat: add login functionality) [1.0.0 -> 1.1.0]\033[0m\n" "feat:" "New feature (Increments MINOR version)"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., fix: resolve nil pointer in auth) [1.0.0 -> 1.0.1]\033[0m\n" "fix:" "Bug fix (Increments PATCH version)"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., docs: update api reference)\033[0m\n" "docs:" "Documentation changes"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., chore: update dependencies)\033[0m\n" "chore:" "Maintenance tasks or tools"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., refactor: simplify event loop)\033[0m\n" "refactor:" "Code changes (no fix/feat)"
	@echo "  --- (Use '!' for Breaking Changes, e.g., feat!:) ---"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., feat!: change auth protocol) [1.1.0 -> 2.0.0]\033[0m\n" "feat!:" "New feature with breaking changes"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., fix!: remove deprecated endpoint) [1.0.1 -> 2.0.0]\033[0m\n" "fix!:" "Bug fix with breaking changes"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., docs!: rewrite entire manual) [1.0.0 -> 2.0.0]\033[0m\n" "docs!:" "Documentation with breaking changes"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., chore!: drop support for node 14) [1.0.0 -> 2.0.0]\033[0m\n" "chore!:" "Maintenance with breaking changes"
	@printf "  \033[36m%-10s\033[0m %s \033[90m(e.g., refactor!: rename core interfaces) [1.0.0 -> 2.0.0]\033[0m\n" "refactor!:" "Refactoring with breaking changes"
	@echo ""

start:
	bash bin/app/start.sh

test:
	bash bin/app/test.sh

test-e2e:
	bash bin/app/test_e2e_grpc.sh

test-e2e-stress:
	bash bin/app/test-e2e-stress.sh

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