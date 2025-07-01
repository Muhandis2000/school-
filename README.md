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

- `main.go` – запуск HTTP-сервера
- `config/` – YAML + чтение конфигурации
- `internal/` – вся логика проекта:
  - `database/` – подключение БД
  - `handlers/` – HTTP-обработчики
  - `models/` – SQLx-модели таблиц
  - `utils/` – JWT, пароли
  - `middleware/` – защита маршрутов
  - `logger/` – логирование
- `migrations/0001_init.sql` – SQL скрипт создания таблиц
- `docs/swagger.yaml` – Swagger спецификация
- `.env` – конфиденциальные данные
- `go.mod / go.sum` – зависимости Go

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
