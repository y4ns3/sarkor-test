include .env

DB_URL = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)

migrate:
	goose -dir ./db/migrations postgres $(DB_URL) up

migrate-down:
	goose -dir ./db/migrations postgres $(DB_URL) down
