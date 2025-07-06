# ===== Локальная разработка =====
run-bot-dev:
	@echo "Запуск бота в dev-режиме..."
	go run ./cmd/main.go

# ===== Продакшен =====
run-bot-prod:
	@echo "Запуск prod-окружения..."
	docker-compose -f services/server/prod/docker-compose.prod.yml up -d --build

down-bot-prod:
	@echo "Остановка prod-окружения..."
	docker-compose -f services/server/prod/docker-compose.prod.yml down