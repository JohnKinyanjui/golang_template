# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

MIGRATION_DIR = "/internal/db/migrations"

ifeq ($(OS),Windows_NT)
	MAIN_PATH = /tmp/bin/main.exe
	SYNC_ASSETS_COMMAND =	@go run github.com/makiuchi-d/arelo@v1.13.1 \
	--target "./public" \
	--delay "100ms" \
	--templ generate --notify-proxy
else
	MAIN_PATH = tmp/bin/main
	SYNC_ASSETS_COMMAND =	@go run github.com/cosmtrek/air@v1.51.0 \
	--build.bin "true" \
	--build.delay "100" \
	--build.exclude_dir "" \
	--build.include_dir "public" \
	--screen.clear_on_rebuild true \
	--log.main_only true
endif


# run air to detect any go file changes to re-build and re-run the server.
server:
	@air \
		--build.cmd "go build --tags dev -o ${MAIN_PATH} ./cmd" \
		--build.bin "${MAIN_PATH}" \
		--build.delay "100" \
		--build.include_ext "go" \
		--build.stop_on_error "false" \
		--misc.clean_on_exit true \
		--screen.clear_on_rebuild true \
		--log.main_only true


db-status:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) status

db-reset:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) reset

db-down:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) down

db-up:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) up

db-mig-create:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) create $(filter-out $@,$(MAKECMDGOALS)) sql

