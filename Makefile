PHONY: db-migrate

DB_URL=postgres://postgres:postgres@localhost:5432/chirpy

db-migrate-up:
	cd sql/schema && goose postgres $(DB_URL) up

db-migrate-down:
	cd sql/schema && goose postgres $(DB_URL) down
