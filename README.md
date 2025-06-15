# PIKLNK - URL Shortener Service

<p align="center">
  <img src="./.assets/banner.png" width="600" alt="PIKLNK Banner"/>
</p>

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
make run-server-dev
```

Сервер будет доступен на http://localhost:8000