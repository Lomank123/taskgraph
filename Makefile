run:
	-@source .env && go run cmd/app/main.go || true
build:
	go build -o bin/taskgraph cmd/app/main.go
test:
	go test ./...
run-db:
	source .env &&docker compose up -d taskgraph-db

# Migrations

goose-up:
	source .env && goose -dir migrations postgres "host=localhost port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB}" up
goose-down:
	source .env && goose -dir migrations postgres "host=localhost port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB}" down
goose-status:
	source .env && goose -dir migrations postgres "host=localhost port=${POSTGRES_PORT} user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=${POSTGRES_DB}" status