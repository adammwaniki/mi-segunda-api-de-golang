build:
	@go build -o bin/mi-segunda-api-de-golang cmd/migrate/main.go

test:
	@go test -v ./...

run: build
	@./bin/mi-segunda-api-de-golang