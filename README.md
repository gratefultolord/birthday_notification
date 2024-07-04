# Birthday Notification

## Описание
Сервис для удобного поздравления сотрудников с днем рождения. Поддерживает подписку и отписку на уведомления о днях рождения конкретных сотрудников.

## Установка и запуск

### Шаг 1: Клонирование репозитория
```sh
git clone git@github.com:gratefultolord/birthday_notification.git
cd birthday_notification
```
### Шаг 2: : Инициализация модуля и установка зависимостей
```sh
go mod tidy
```
### Шаг 3: Настройка базы данных
База данных SQLite будет создана автоматически при запуске приложения. Вы можете использовать встроенные тестовые данные.

### Шаг 4: Запуск приложения
```sh
go run cmd/main.go
```
## Тестирование API
Получение списка дней рождения
```sh
curl -H "Authorization: Bearer valid_token" http://localhost:8080/birthdays
```
Подписка на уведомления
```sh
curl -X POST -H "Authorization: Bearer valid_token" -H "Content-Type: application/json" -d '{"user_id": 1, "email": "subscriber@example.com"}' http://localhost:8080/subscribe
```
Отписка от уведомлений
```sh
curl -X POST -H "Authorization: Bearer valid_token" -H "Content-Type: application/json" -d '{"user_id": 1, "email": "subscriber@example.com"}' http://localhost:8080/unsubscribe
```
## Запуск тестов
```sh
go test ./tests
```
## Структура каталогов
```plaintext
cmd/: Главный файл для запуска приложения.
  main.go: Точка входа для запуска сервера.
config/: Конфигурация приложения.
  config.go: Настройка и инициализация базы данных.
internal/: Внутренние пакеты приложения.
  auth/: Обработка авторизации.
    auth.go: Middleware для авторизации.
  handlers/: HTTP-обработчики.
    birthday.go: Обработчик для получения дней рождения.
    subscription.go: Обработчики для подписки и отписки.
  models/: Модели данных.
    user.go: Модель пользователя.
  repository/: Логика доступа к данным.
    user.go: Репозиторий для пользователей.
    subscription.go: Репозиторий для подписок.
  utils/: Вспомогательные функции.
    utils.go: Функции для работы с датами и JSON-ответами.
pkg/: Общие пакеты, используемые в разных частях приложения.
  server/: Настройка сервера и роутера.
    server.go: Конфигурация маршрутов и создание сервера.
  middleware/: Промежуточное ПО.
    auth.go: Middleware для авторизации.
tests/: Юнит-тесты.
  auth_test.go: Тесты для авторизации.
  birthday_test.go: Тесты для обработчика дней рождения.
  subscription_test.go: Тесты для обработчиков подписки и отписки.
