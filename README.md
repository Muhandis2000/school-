# Online School Management System

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?logo=go)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16+-4169E1?logo=postgresql)
![Gin](https://img.shields.io/badge/Gin-1.9+-000000?logo=go)

Этот проект представляет собой REST API для управления онлайн-школой, написанный на языке Go с использованием фреймворка Gin-Gonic. Он включает работу с базой данных PostgreSQL, аутентификацию через JWT, логирование и конфигурацию через JSON и .env файлы.

## Основные функции

- 📝 Регистрация и аутентификация пользователей
- 🎓 Управление курсами и уроками
- 👨‍🎓 Запись студентов на курсы
- 🔐 Ролевая модель доступа (admin/teacher/student)
- 📊 Логирование операций
- 🔧 Миграции базы данных

## Технологический стек

- **Язык**: Go 1.24+
- **Фреймворк**: Gin-Gonic
- **База данных**: PostgreSQL 16+
- **Аутентификация**: JWT
- **Конфигурация**: JSON + .env
- **Логирование**: Файловые логи

## 📁 Структура проекта

online-school/
│
├── main.go                        # Точка входа в приложение (запускает сервер Gin)
├── config/                        # Конфигурации проекта
│   ├── config.yaml                # Файл с настройками (порт, БД и т.д.)
│   └── config.go                  # Код для чтения конфигурации из YAML
├── internal/                      # Внутренние модули (логика проекта)
│   ├── database/                  
│   │   └── db.go                  # Подключение и инициализация PostgreSQL с sqlx
│   ├── handlers/                  # Обработчики HTTP-запросов (контроллеры)
│   │   ├── auth.go                # Регистрация, логин, JWT, профиль
│   │   ├── admin.go               # CRUD и верификация студентов/учителей
│   │   ├── teacher.go             # Отметка посещения и выставление оценок
│   │   └── student.go             # Просмотр уроков, сдача домашки, сертификат
│   ├── logger/
│   │   └── logger.go              # Логгирование с использованием logrus
│   ├── middleware/
│   │   └── auth_middleware.go     # JWT-проверка и роль доступа
│   ├── models/
│   │   └── models.go              # SQLX-структуры и модели данных
│   └── utils/
│       ├── jwt.go                 # Создание, парсинг и валидация JWT
│       └── password.go           # Хэширование и сравнение паролей (bcrypt)
├── migrations/
│   └── 0001_init.sql             # SQL-файл для создания всех таблиц базы данных
├── docs/
│   └── swagger.yaml              # Swagger/OpenAPI 3.0 спецификация API
├── go.mod                        # Модуль Go (описание зависимостей)
├── go.sum                        # Контрольные суммы зависимостей
└── README.md                     # Документация проекта


---

## 🔐 Авторизация

- Регистрация: `POST /auth/register`
- Вход: `POST /auth/login`
- Профиль: `GET /auth/me` (нужен JWT)

JWT хранится в `Authorization: Bearer <token>`.

---

## 👑 Роли и доступ

| Роль      | Доступ к маршрутам                            |
|-----------|-----------------------------------------------|
| Admin     | Верификация, CRUD всех сущностей              |
| Teacher   | Посещаемость, оценки                          |
| Student   | Уроки, расписание, домашки, сертификаты       |

## Проверка в Свагер

[Открыть Swagger UI](http://localhost:8080/swagger/index.html)

## Ссылка на проект

[https://github.com/Muhandis2000/school.git](https://github.com/Muhandis2000/school-.git)
