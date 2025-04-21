GO := go
MINIMOCK := minimock
MOCK_DIR := ./mocks
SOURCE_DIR := ./internal
GOOSE := goose
MIGRATIONS_DIR := ./scripts/migrations
DB_STRING := "postgres://postgres:@localhost:5432/screen-time-limiter?sslmode=disable" # Замените на ваши параметры подключения к БД
DB_STRING_TEST := "postgres://postgres:@localhost:5432/screen-time-limiter-test?sslmode=disable" # Замените на ваши параметры подключения к БД

# Цель по умолчанию
all: generate-mocks

# Генерация моков
generate-mocks:
	@echo "Generating mocks..."
	@$(GO) get github.com/gojuno/minimock/v3
	@$(GO) install github.com/gojuno/minimock/v3/cmd/minimock@latest
	@$(MINIMOCK) -i ./internal/app/handlers.* -o internal/app/handlers
	@echo "Mocks generated successfully!"

# Очистка сгенерированных моков
clean-mocks:
	@echo "Cleaning mocks..."
	@rm -rf $(MOCK_DIR)
	@echo "Mocks cleaned successfully!"

# Запуск тестов
test: generate-mocks
	@echo "Running tests..."
	@$(GO) test ./...
	@echo "Tests completed!"

# Очистка
clean: clean-mocks
	@echo "Cleaning up..."
	@$(GO) clean
	@echo "Cleanup completed!"

# Установка goose
install-goose:
	@echo "Installing goose..."
	@$(GO) install github.com/pressly/goose/v3/cmd/goose@latest
	@echo "Goose installed successfully!"

# Применение миграций
migrate-up: install-goose
	@echo "Applying migrations..."
	@$(GOOSE) -dir $(MIGRATIONS_DIR) postgres $(DB_STRING) up
	@echo "Migrations applied successfully!"

migrate-up-test: install-goose
	@echo "Applying migrations..."
	@$(GOOSE) -dir $(MIGRATIONS_DIR) postgres $(DB_STRING_TEST) up
	@echo "Migrations applied successfully!"

# Откат миграций
migrate-down: install-goose
	@echo "Rolling back migrations..."
	@$(GOOSE) -dir $(MIGRATIONS_DIR) postgres $(DB_STRING) down
	@echo "Migrations rolled back successfully!"

migrate-down-test: install-goose
	@echo "Rolling back migrations..."
	@$(GOOSE) -dir $(MIGRATIONS_DIR) postgres $(DB_STRING_TEST) down
	@echo "Migrations rolled back successfully!"

# Создание новой миграции
create-migration: install-goose
	@echo "Creating new migration..."
	@$(GOOSE) -dir $(MIGRATIONS_DIR) create $(name) sql
	@echo "Migration $(name) created successfully!"

.PHONY: all generate-mocks clean-mocks test clean install-goose migrate-up migrate-down create-migration

run:
	CONFIG_FILE=/Users/mihailisaev/Documents/_projects/screen-time-limiter/configs/values_local.yaml go run cmd/screen-time-limiter/main.go