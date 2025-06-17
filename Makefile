# ===== Конфигурация =====
BACKEND_DIR := backend
FRONTEND_ENTRY_DIR := frontend/entry
PROXY_DIR := services/proxy

# ===== Валидация окружения =====
check-env:
ifndef ENV
	$(error ENV не определен. Используйте ENV=dev или ENV=prod)
endif

# ===== Инициализация проекта =====
init-env:
	@if [ -f .env ]; then \
		cp .env $(BACKEND_DIR)/.env && \
		echo "Файл .env скопирован в $(BACKEND_DIR)/"; \
	else \
		echo "Предупреждение: .env файл не найден в корне проекта"; \
	fi

# ===== Управление базой данных =====
db-up:
	@$(MAKE) -C $(BACKEND_DIR) db-up

db-down:
	@$(MAKE) -C $(BACKEND_DIR) db-down

db-migrate:
	@$(MAKE) -C $(BACKEND_DIR) db-migrate

# ===== Бэкенд сервер =====
run-server: check-env
ifeq ($(ENV),dev)
	@$(MAKE) -C $(BACKEND_DIR) run-server-dev
else ifeq ($(ENV),prod)
	@$(MAKE) -C $(BACKEND_DIR) run-server-prod
else
	$(error Недопустимое значение ENV: $(ENV). Используйте dev или prod)
endif

down-server: check-env
ifeq ($(ENV),prod)
	@$(MAKE) -C $(BACKEND_DIR) down-server-prod
else
	@echo "Предупреждение: down-server поддерживается только для prod окружения"
endif

# ===== Фронтенд (entry) =====
run-entry: check-env
ifeq ($(ENV),dev)
	@$(MAKE) -C $(FRONTEND_ENTRY_DIR) run-entry-dev
else ifeq ($(ENV),prod)
	@$(MAKE) -C $(FRONTEND_ENTRY_DIR) run-entry-prod
endif

down-entry: check-env
ifeq ($(ENV),dev)
	@$(MAKE) -C $(FRONTEND_ENTRY_DIR) down-entry-dev
else ifeq ($(ENV),prod)
	@$(MAKE) -C $(FRONTEND_ENTRY_DIR) down-entry-prod
endif

# ===== Прокси =====
run-proxy:
	docker compose -f $(PROXY_DIR)/docker-compose.yml up -d --build

down-proxy:
	docker compose -f $(PROXY_DIR)/docker-compose.yml down

# ===== Утилиты =====
clean:
	@docker system prune -f
	@echo "Система очищена"