# Dynamic user segmentation service

Микросервис для работы с сегментированием пользователей.

Используемые технологии:

- PostgreSQL (хранилище данных)
- Docker (запуск сервиса)
- Swagger (документация API)
- Gin (веб-фреймворк)
- golang-migrate/migrate (миграции БД)
- sqlx (работа с БД)
- logrus (логгирования)
- pq (драйвер для работы с postgres)
- viper (работа с файлами конфигураций)

Сервис расширяемый, использовался механизм Graceful Shutdown для корректного завершения работы сервиса.

# Quickstart:

Запуск сервиса `make run`

Service available on `localhost:8000`

Опционально настроить файлы `config/config.yaml` и `.env`

Swagger-документация доступна по адресу: `http://127.0.0.1:8000/swagger/index.html#`

# Examples

Примеры запросов

- [Добавление пользователя](#add-user)
- [Удаление пользователя](#delete-user)
- [Получение списка всех пользователей](#get-all-users)
- [Создание сегмента](#add-segment)
- [Изменение сегмента](#patch-segment)
- [Удаление сегмента](#delete-segment)
- [Получение списка всех сегментов](#get-all-segments)
- [Добавление пользователя в сегменты](#add-segments-byuserid)
- [Удаление пользователя из сегментов](#delete-segments-byuserid)
- [Просмотр активных сегментов пользователя](#get-user-segments)

> Данные запросы доступны для импорта в Postman `./docs/postman`

### Добавление пользователя <a name="add-user"></a>

Добавление пользователя в базу по его имени и паролю, пароль кэшируется:
```
curl -X 'POST' \
  'http://127.0.0.1:8000/user/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "username": "Semen",
    "password": "vbhnhelvfq14"
}'
```

Пример ответа:
```
{
  "id": 5
}
```

### Удаление пользователя <a name="delete-user"></a>

Удаление пользователя из базы по его id:
```
curl -X 'DELETE' \
  'http://127.0.0.1:8000/user/5' \
  -H 'accept: application/json'
```

Пример ответа:
```
{
  "status": "ok"
}
```

### Получение списка всех пользователей <a name="get-all-users"></a>

Получение всех пользователей в базе:
```
curl -X 'GET' \
  'http://127.0.0.1:8000/user/' \
  -H 'accept: application/json'
```

Пример ответа:
```
{
  "data": [
    {
      "id": 1,
      "username": "nikita"
    },
    {
      "id": 2,
      "username": "denis"
    },
    {
      "id": 3,
      "username": "ivan"
    },
    {
      "id": 4,
      "username": "sergey"
    }
  ]
}
```

### Создание сегмента <a name="add-segment"></a>

Добавление сешмента в базу по его названию:
```
curl -X 'POST' \
  'http://127.0.0.1:8000/segment/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "segmentname": "test"
}'
```

Пример ответа:
```
{
  "id": 5
}
```

### Изменение сегмента <a name="patch-segment"></a>

Изменение сегмента в базе по его id:
```
curl -X 'PATCH' \
  'http://127.0.0.1:8000/segment/5' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "segmentname": "AVITO_PREMIUM"
}'
```

Пример ответа:
```
{
  "status": "ok"
}
```

### Удаление сегмента <a name="delete-segment"></a>

Удаление сегмента из базы по его названию:
```
curl -X 'DELETE' \
  'http://127.0.0.1:8000/segment/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "segmentname": "AVITO_PREMIUM"
}'
```

Пример ответа:
```
{
  "status": "ok"
}
```

### Получение списка всех сегментов <a name="get-all-segments"></a>

Получение всех сегментов в базе:
```
curl -X 'GET' \
  'http://127.0.0.1:8000/segment/' \
  -H 'accept: application/json'
```

Пример ответа:
```
{
  "data": [
    {
      "id": 1,
      "segmentname": "AVITO_VOICE_MESSAGES"
    },
    {
      "id": 2,
      "segmentname": "AVITO_PERFORMANCE_VAS"
    },
    {
      "id": 3,
      "segmentname": "AVITO_DISCOUNT_30"
    },
    {
      "id": 4,
      "segmentname": "AVITO_DISCOUNT_50"
    }
  ]
}
```

### Добавление пользователя в сегменты <a name="add-segments-byuserid"></a>

Добавление пользователя в сегменты по id пользователя и списку названий(я) сегментов:
```
curl -X 'POST' \
  'http://127.0.0.1:8000/user-segment/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": 1,
    "segmentlist": [
        "AVITO_PERFORMANCE_VAS",
        "AVITO_DISCOUNT_50"
        ]
}'
```

Пример ответа:
```
{
  "status": "ok"
}
```

### Удаление пользователя из сегментов <a name="delete-segments-byuserid"></a>

Добавление пользователя из сегментов по id пользователя и списку названий(я) сегментов:
```
curl -X 'DELETE' \
  'http://127.0.0.1:8000/user-segment/' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
    "id": 1,
    "segmentlist": [
        "AVITO_PERFORMANCE_VAS",
        "AVITO_DISCOUNT_50"
        ]
}'
```

Пример ответа:
```
{
  "status": "ok"
}
```

### Просмотр активных сегментов пользователя <a name="get-user-segments"></a>

Получение сегментов, в которых состоит пользователь по его id:
```
curl -X 'GET' \
  'http://127.0.0.1:8000/user-segment/1' \
  -H 'accept: application/json'
```

Пример ответа:
```
{
  "data": [
    {
      "id": 1,
      "segmentname": "AVITO_VOICE_MESSAGES"
    },
    {
      "id": 3,
      "segmentname": "AVITO_DISCOUNT_30"
    }
  ]
}
```

# Decisions

