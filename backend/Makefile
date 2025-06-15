# ===== Локальная БД =====
db-up:
	@echo "Запуск PostgreSQL..."
	docker-compose -f services/db/postgresql/docker-compose.yml up -d

db-down:
	@echo "Остановка PostgreSQL..."
	docker-compose -f services/db/postgresql/docker-compose.yml down

# ===== БД =====
db-migrate:
	@echo "Применение миграций..."
	go run ./migrations/auto.go

# ===== Локальная разработка =====
run-server-dev:
	@echo "Запуск сервера в dev-режиме..."
	go run ./cmd/main.go

# ===== Продакшен =====
run-server-prod:
	@echo "Запуск prod-окружения..."
	docker-compose -f services/server/prod/docker-compose.prod.yml up -d --build

down-server-prod:
	@echo "Остановка prod-окружения..."
	docker-compose -f services/server/prod/docker-compose.prod.yml down

