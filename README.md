# Product REST API
Rest API для управления складом продуктов. API предоставляет фукнционал  для добавления, обновления, удаления и
получения информации о продуктах, хранящихся на складе.

## Стек технологий :

- Go
- Gin
- PostgreSQL
- pgx
- goose (migrations)
- docker-compose

##  Структура проекта

```
sarkor-test/
|
├── cmd/                    # Точка входа (main.go)
├── internal/
│   ├── entity/             # Сущности
│   ├── repository/         # Работа с базой данных
│   ├── rest/               # Запуск сервера , роутинг и реализация handler
│   └── usecase/            # Бизнес логика
| 
├── config/                 # Конфигурация приложения
├── db/                     # Подключение к бд и миграции
├── .env                    # Переменные окружения
├── Makefile                # Скрипты для миграций
├── docker-compose.yml      # Docker конфигурация
├── test.http               # Пример использования           
├── go.mod       
├── go.sum         
└── README.md
```
## Доступные endpoints

| Метод  | URL                  | Описание                      | Тело запроса (пример)                          |
|--------|----------------------|-------------------------------|-----------------------------------------------|
| POST   | `/products`          | Создать новый продукт          | `{ "name": "Product1", "description": "Desc", "price": 100, "quantity": 10 }` |
| GET    | `/products`          | Получить список всех продуктов | —                                             |
| GET    | `/products/:id`      | Получить продукт по ID         | —                                             |
| PUT    | `/products/:id`      | Обновить продукт по ID         | `{ "name": "Updated", "description": "New desc", "price": 120, "quantity": 5 }` |
| DELETE | `/products/:id`      | Удалить продукт по ID          | —                                             |

## Запуск приложения

1. Клонируйте репозиторий:

```bash
git clone https://github.com/y4ns3/sarkor-test
```
```bash
cd sarkor-test
```

2. Создайте файл окружения на основе примера
#### при необходимости измените параметры под себя

```bash
cp .env.example .env
```

3. Запустите Docker Compose:
```bash
docker-compose up
```
4. Выполните миграцию

```bash
make migrate
```
5. Запустите приложение
```bash
go run cmd/app/main.go
```
