install:
	go get ./...

prepare: install migrate-up

start:
	go run cmd/main.go

test:
	go test -v -coverpkg=./... -coverprofile=coverage.out ./app/tests/...

cover:
	go tool cover -func=coverage.out

lint:
	golangci-lint run ./...
