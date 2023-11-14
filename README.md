# nats-streaming-web

## Обзор
Данный проект представляет собой сервис, который использует NATS-Streaming для подписки на канал, PostgreSQL для хранения данных и Redis для кэширования. Сервис восстанавливает кэш из базы данных в случае сбоя и поддерживает HTTP-сервер для доступа к данным по ID.

## Требования
- Docker и Docker Compose
- Golang

## Установка и Запуск
1. **Запуск PostgreSQL, Redis и NATS-Streaming с помощью Docker Compose:**
   - Убедитесь, что Docker и Docker Compose установлены.
   - Запустите `docker-compose up` из корня проекта для инициализации и запуска всех необходимых сервисов.

2. **Настройка Базы Данных:**
   - Создайте свою базу данных в PostgreSQL.
   - Настройте пользователя и создайте таблицы согласно предоставленной модели данных.

3. **Запуск сервиса:**
   - Из директории `cmd`, запустите `go run main.go` для запуска сервера.

## Использование
- **HTTP-сервер:** После запуска сервиса, HTTP-сервер будет доступен для получения данных по ID.
- **Интерфейс:** Простейший интерфейс доступен для отображения данных. Доступен через веб-браузер.

## Архитектура
- **cmd/main.go:** Точка входа в приложение.
- **dbclient/dbclient.go:** Клиент для взаимодействия с базой данных.
- **internal/config/config.go:** Конфигурация сервиса.
- **internal/logger/logger.go:** Логгер для отслеживания действий в сервисе.
- **service:** Содержит основную логику сервиса.
- **pkg/cache/redis_client.go:** Клиент Redis для кэширования.
- **pkg/client/nats_client.go:** Клиент NATS-Streaming для подписки на канал.
- **pkg/model/model.go:** Модель данных.
- **transport/rest:** HTTP-транспорт и обработчики.
- **templates/index.html:** Шаблон для интерфейса пользователя.
- **testutils/test_data.go:** Утилита для инициализации тестовых данных.

## Тестирование
- Запустите `go test ./...` в корне проекта для запуска всех тестов.

---
