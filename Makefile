# ===== Инициализация окружения (копирование файла) =====
init-env:
	@cp .env ./backend/.env || true
	@echo "Файл скопирован: .env -> ./backend/.env"

# ===== Локальная БД =====
db-up:
	@$(MAKE) -C backend db-up

db-down:
	@$(MAKE) -C backend db-down

# ===== БД =====
db-migrate:
	@$(MAKE) -C backend db-migrate

# ===== Локальная разработка =====
run-server-dev:
	@$(MAKE) -C backend run-server-dev

# ===== Продакшен =====
run-server-prod:
	@$(MAKE) -C backend run-server-prod

down-server-prod:
	@$(MAKE) -C backend down-server-prod
