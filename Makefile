.PHONY: create-migration, migrate-up, migrate-down

DB_URL ?= postgres://user:pgpwd4@localhost:54326/appDB?sslmode=disable

create-migration-%:
	goose -dir internal/repository/migrations create $* sql

migrate-up:
	goose -dir internal/repository/migrations postgres "$(DB_URL)" up

migrate-down:
	goose -dir internal/repository/migrations postgres "$(DB_URL)" down

swag:
	swag init -g ./cmd/app/main.go