set dotenv-load := true


run:
    go run cmd/main.go

build:
    go build

test:
    go test .../.
    