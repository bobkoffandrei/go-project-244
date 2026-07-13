build:
	go build -o bin/gendiff ./cmd/gendiff

.PHONY: test
test:
	go test -v ./...

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix