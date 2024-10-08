# Effective Mobile Music API

## Описание проекта

Этот проект представляет собой RESTful API для управления музыкальной библиотекой, разработанный с использованием Go и PostgreSQL. API предоставляет функционал для управления песнями и пользователями, включая аутентификацию и авторизацию.

## Основные возможности

- Регистрация и аутентификация пользователей
- CRUD операции для песен (Создание, Чтение, Обновление, Удаление)
- Разграничение доступа (пользователь/администратор)
- Аутентификация на основе JWT (JSON Web Tokens)
- Пагинация результатов при получении списка песен
- Возможность получения определенного количества куплетов песни
- Swagger

## Технологический стек

- Язык программирования: Go
- База данных: PostgreSQL
- Веб-фреймворк: Gin
- ORM: Нативный SQL (без использования ORM)
- Контейнеризация: Docker и Docker Compose
- Управление зависимостями: Go Modules
- Swagger

## Предварительные требования

- Docker
- Docker Compose
- Make (опционально, для использования команд Makefile)
- Go 1.23 или выше (для локальной разработки)
- Swagger

## Начало работы

### Клонирование репозитория

```bash
git clone https://github.com/typefunco/Effective_Mobile_Test_Task/tree/main.git
cd effective-mobile
```

### Запуск приложения

Используя Docker Compose:

```bash
docker-compose up --build
```

Или с помощью Make:

```bash
make start_service
```

API будет доступно по адресу `http://localhost:8080`.

Для запуска приложения без Docker:

```bash
make run
```

## Эндпоинты API

### Публичные эндпоинты

- `POST /sign-up`: Регистрация нового пользователя

### Защищенные эндпоинты (требуют JWT аутентификации)

- `GET /music/songs`: Получение всех песен
- `GET /music/songs/:song_id`: Получение определенного количества песен
- `GET /music/song/:song_id/:verse`: Получение определенного количества куплетов песни
- `PATCH /update`: Повышение статуса пользователя до администратора

### Эндпоинты только для пользователей с правами администратора

- `DELETE /music/song/:id`: Удаление песни
- `PATCH /music/songs/:song_id`: Обновление информации о песне
- `POST /music/song/new`: Создание новой песни

## Аутентификация

Для доступа к защищенным эндпоинтам необходимо включить JWT токен в заголовок Authorization:

```
Authorization: Bearer <ваш_jwt_токен>
```

## База данных

Проект использует PostgreSQL. База данных автоматически инициализируется примерными данными при запуске с помощью Docker Compose.

### Структура базы данных

1. Таблица `users`:
   - `user_id`: SERIAL PRIMARY KEY
   - `username`: VARCHAR(16) UNIQUE
   - `user_password`: VARCHAR
   - `is_admin`: BOOLEAN

2. Таблица `songs`:
   - `song_id`: SERIAL PRIMARY KEY
   - `song_name`: VARCHAR
   - `song_text`: VARCHAR
   - `release_date`: VARCHAR
   - `song_link`: VARCHAR
   - `song_author`: VARCHAR

## Переменные окружения

- `DBURL`: Строка подключения к PostgreSQL
- `JWT_SECRET`: Секретный ключ для генерации JWT токенов

Эти переменные устанавливаются в файле `docker-compose.yml` для развертывания в контейнерах.

## Разработка

### Структура проекта

```
.
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── entity/
│   │   ├── song.go
│   │   └── user.go
│   ├── handlers/
│   │   ├── music.go
│   │   ├── router.go
│   │   └── users.go
│   ├── middleware/
│   │   ├── auth.go
│   │   └── isAdmin.go
│   ├── repo/
│   │   └── db.go
│   └── utils/
│       └── jwt.go
├── migrations/
│   └── init.sql
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

### Команды Makefile

- `make run`: Запуск приложения локально
- `make start_service`: Запуск сервиса в контейнере

## Особенности реализации

- Использование паттерна Repository для абстрагирования работы с базой данных
- Middleware для аутентификации и проверки прав администратора
- Использование prepared statements для защиты от SQL-инъекций
- Хеширование паролей пользователей перед сохранением в базу данных
- Пагинация результатов для оптимизации производительности при больших объемах данных


## Контакты

Telegram - https://t.me/typefunco
