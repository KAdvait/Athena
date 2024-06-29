build:
	@go build -o bin/athena cmd/main.go

run: build
	@./bin/athena

test:
	@go test -v ./...