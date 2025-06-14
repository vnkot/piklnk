# PIKLNK - URL Shortener Service

<p align="center">
  <img src="./.assets/banner.png" width="600" alt="PIKLNK Banner"/>
</p>

## Разработка
### Сервер
Чтобы все полетело, в корне проекта нужно создать `.env` файл со следующим содержимым:
```bash
SECRET=<your_secret>
DSN=postgres://piklnk:piklnk@localhost:5432/postgres?sslmode=disable
```
Где:
- *SECRET* - секрет для jwt токенов с алгоритмом HS256
- *DSN* - dsn для подключения к БД 

Инициализируйте окружение:
```bash
make init-env
```

Запустите базу данных:
```bash
make db-up
```
Примените миграции:
```bash
make db-migrate
```

Запустите сервер:
```bash
make run-dev
```

Сервер будет доступен на http://localhost:8000