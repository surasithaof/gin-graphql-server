set dotenv-load := true
set export

run:
    go run cmd/main.go
build:
    go build
test:
    go test .../.
show-cov:
    go tool cover -html=coverage.out
generate:
    go generate ./...


gql-install:
    go get -u github.com/99designs/gqlgen
gql-init:
    go run github.com/99designs/gqlgen init
gql-gen:
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


# release:
#     # autotag --scheme=conventional
release:
    git sv next-version

env name:
	ln -sf .env.{{name}} .env

# https://github.com/pantheon-systems/autotag
#  go install  github.com/git-chglog/git-chglog/cmd/git-chglog