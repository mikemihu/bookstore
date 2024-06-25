help: ## show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

wire: ## Generate wire
	@go generate ./internal/provider/.

run: ## Run the app
	@go run ./cmd/.

gen-mock: ## Generate mocks
	@go generate ./mocks/.

test: ## Run unit test with coverage info
	@go test ./... -v --cover

# local development db dsn
GOOSE_DBSTRING="host=127.0.0.1 port=5432 user=gotu password=password dbname=gotu sslmode=disable"

migrate: ## Run db migration up
	@GOOSE_DRIVER=postgres \
 	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
 	goose -dir migrations up

migrate-down: ## Run db migration down
	@GOOSE_DRIVER=postgres \
 	GOOSE_DBSTRING=$(GOOSE_DBSTRING) \
 	goose -dir migrations down

.PHONY: help wire run gen-mock test migrate migrate-down
