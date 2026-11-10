# Переменные для удобства
DB_DSN="host=localhost port=5432 user=fintech_user password=fintech_password dbname=fintech_limits sslmode=disable"
MIGRATIONS_DIR=./migrations

.PHONY: deps init-db migrate-status migrate-up migrate-down

# 1. Установка утилиты goose (команда для новой машины или прода)
deps:
	go install github.com/pressly/goose/v3/cmd/goose@latest

# 2. Создание файла новой миграции (например: make migrate-create name=create_users_table)
migrate-create:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

# 3. Проверка статуса миграций
migrate-status:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) status

# 4. Накат всех миграций (Up)
migrate-up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) up

# 5. Откат последней миграции (Down)
migrate-down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) down
