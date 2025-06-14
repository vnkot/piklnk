# ===== Переменные =====
BACKEND_DIR     = backend
DB_COMPOSE      = services/db/postgresql/docker-compose.yml
MIGRATIONS_PATH = migrations/auto.go
APP_ENTRYPOINT  = ./cmd/main.go

# ===== Цели по умолчанию =====
.DEFAULT_GOAL := help

# ===== Help =====
help:
	@echo "Доступные команды:"
	@echo "  init-env       - Создать симлинк .env в backend"
	@echo "  db-up          - Запустить PostgreSQL в Docker"
	@echo "  db-down        - Остановить PostgreSQL"
	@echo "  db-migrate     - Применить миграции БД"
	@echo "  run-dev        - Запустить сервер в режиме разработки"

# ===== Инициализация окружения =====
init-env:
	@ln -sf ../.env $(BACKEND_DIR)/.env || true
	@echo "Симлинк создан: $(BACKEND_DIR)/.env -> ../.env"

# ===== Управление Базой Данных =====
db-up:
	docker-compose -f $(DB_COMPOSE) up -d

db-down:
	docker-compose -f $(DB_COMPOSE) down

db-migrate:
	cd $(BACKEND_DIR) && go run $(MIGRATIONS_PATH)

# ===== Запуск сервера =====
run-dev:
	cd $(BACKEND_DIR) && go run $(APP_ENTRYPOINT)