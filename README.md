# Book Store

A prototype of online book store.

## Requirement

Please make sure you have the following tools installed before you start.

- Go v1.22.4 or newer
- Docker
- [Goose](https://github.com/pressly/goose) — database migration tool.

## How to Run

1. Copy config file
   ```shell
   cp config/app.sample.json config/app.development.json
   ``` 
2. Spin up database & redis with docker
   ```shell
   (cd development && docker-compose up -d)
   ```
3. Run migration  
   ```shell
   make migrate
   ```
4. Run app
   ```shell
   make run
   ```

## Directory Structure

```
cmd/                entry point for the app
config/             config as code file in .json format
development/        related files for development environment
docs/               documentation files
internal/           
 ├─ app/            app startup & shutdown, and web route & cron registration
 ├─ config/         config schema and loader
 ├─ contexts/       context helper
 ├─ constant/       constants widely used in the app
 ├─ delivery/       request, response casting and middleware
 ├─ entity/         entities such models, request & response, filters, etc. 
 ├─ provider/       dependency provider and injection
 ├─ repository/     data source such postgres and redis
 └─ usecase/        business logic
migrations/         database migration in .sql format
pkg/                utility package
```

## Postman

API request & response sample in [Postman Documenter](https://documenter.getpostman.com/view/893849/2sA3drKFJ2)

## Commands

Find all existing commands by run `make help` command.
```
gen-mock                       Generate mocks
help                           Show this help
lint                           Run golangci-lint
migrate                        Run db migration up
migrate-down                   Run db migration down
run                            Run the app
test                           Run unit test with coverage info
wire                           Generate wire
```

Create new migration file

```shell
cd migrations
goose -s create <file_name> sql
```

