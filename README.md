# Gin GraphQL server

Basic GraphQL Server on [Gin](https://github.com/gin-gonic/gin) framework and [gqlgen](https://github.com/99designs/gqlgen) written by [Go](https://go.dev/).

- [x] Use [gqlgen](https://github.com/99designs/gqlgen) to create graphql server
- [x] Connect to Database (PostgresQL) using [Bun](https://bun.uptrace.dev/) or [GORM](https://gorm.io/) for ORM
- [x] Database migration
- [ ] Dataloader to solve N+1 database queries
- [ ] Subsciptions
- [ ] Unit test
- [ ] Security middleware (JWT, Gocloak)
- [x] Docker support

---

## Setup

[Justfile](https://github.com/casey/just) for running command

```bash
brew install just
```

[golang-migrate](https://github.com/golang-migrate/migrate) to migrate database.

```bash
brew install golang-migrate
```

## Set up .env file

Create `.env` file and set up environment variables (you can copy from `.env.example`)

## Create and migrate database

You can start postgresql database by run this command `docker-compose up`,
And do the migration by run `just db-migrate`

## Start service 🚀

```bash
just run
```
