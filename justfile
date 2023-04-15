set dotenv-load := true


run:
    go run cmd/main.go

build:
    go build

test:
    go test .../.

gql-install:
    go get github.com/99designs/gqlgen

gql-init:
    go run github.com/99designs/gqlgen init

gql-generate:
    go run github.com/99designs/gqlgen generate