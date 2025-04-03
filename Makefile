build:
	@go build -o bin/mi-segunda-api-de-golang cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/mi-segunda-api-de-golang

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down