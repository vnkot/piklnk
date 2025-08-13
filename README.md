# PIKLNK - URL Shortener Service

<p align="center">
  <img src="./.assets/banner.png" width="600" alt="PIKLNK Banner"/>
</p>

## Структура проекта
- backend: go rest-api сервер, реализующий основную логику сервиса
- tgbot: телеграм бот для работы с сервисом
- frontend/entry: стартовая страница (на текущий момент заглушка), отвечающая за быстрое создание короткой ссылки
- bruno: готовые запросы для тестирования API
- services/proxy: разводит запросы вида `/` на frontend/entry и `/*` на backend

## Разработка
### Сервер
Чтобы всё заработало, в корне проекта создайте файл .env со следующим содержимым:
```bash
SECRET=<your_secret>
DSN=postgres://piklnk:piklnk@localhost:5432/postgres?sslmode=disable
```
Где:
- SECRET — секрет для JWT токенов с алгоритмом HS256
- DSN — строка подключения к базе данных. Если не используете внешнюю БД, оставьте как в примерее

Инициализируйте окружение:
```bash
make init-env
```

Запустите локальную базу данных (если не используете внешнюю):
```bash
make db-up
```
Примените миграции:
```bash
make db-migrate
```

Запустите сервер в режиме разработки:
```bash
make  ENV=dev run-server
```

Сервер будет доступен на http://localhost:8000

### Телеграм бот
Для работы в .env файл нужно добавить следующее параметры:
```bash
TGBOTTOKEN=<your_bot_token>
APIURL=<your_address_for_server>
```

После чего можно запустить бот:
```bash
make  ENV=dev run-bot
```

### Frontend (entry)
Здесь всё просто:
```bash
make  ENV=dev run-entry
```

### Proxy
А это прям совсем легко
```bash
make run-proxy
```
