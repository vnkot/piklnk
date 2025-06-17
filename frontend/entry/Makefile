DEV_COMPOSE := services/dev/docker-compose.yml
PROD_COMPOSE := services/prod/docker-compose.yml

# ===== Локальная разработка =====
run-entry-dev:
	@echo "Запуск development сервера..."
	docker-compose -f $(DEV_COMPOSE) up -d

down-entry-dev:
	@echo "Остановка development сервера..."
	docker-compose -f $(DEV_COMPOSE) down

# ===== Продакшен =====
run-entry-prod:
	@echo "Запуск production сервера..."
	docker-compose -f $(PROD_COMPOSE) up -d

down-entry-prod:
	@echo "Остановка production сервера..."
	docker-compose -f $(PROD_COMPOSE) down
