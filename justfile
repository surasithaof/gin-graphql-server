set dotenv-load := true
set export

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

env name:
	ln -sf .env.{{name}} .env

envcreate name:
    cp .env.example .env.{{name}}

db-install-migrate:
    brew install golang-migrate

db-create-migrate name:
    migrate create -ext sql -dir migrations -seq {{name}}

db-migrate:
    migrate -database ${DB_CONN} -path migrations up

db-rollback:
    migrate -database ${DB_CONN} -path migrations down

db-migration-force version:
    migrate -database ${DB_CONN} -path migrations force {{version}}
