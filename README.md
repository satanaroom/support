# Сервис для записи данных об объекте "опора"

## Особенности сервиса

- Для запуска требуется docker-контейнер PostgreSQL:
```bash
docker run --name=support -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
```
## API:
### POST api/items
Создание записи
#### Пример ввода:
```json
{
  "number": "3T434qwv",
  "name": "УОА10-1",
  "date": "2022-01-01T13:49:51.141Z"
}
```
#
### GET api/items
Получение всех записей из базы
#
### GET api/items/:id
Получение записи по id
#### Пример ввода:
```json
{
  "id": "2"
}
```
#
### PUT api/items/:id
Обновление записи
#### Пример ввода:
```json
{
    "name": "3T434qwv",
    "number": "УА10-1",
    "date": "2022-01-01T13:49:51.141Z"
}
```
## Запуск:

```bash
make
```

## Применение миграций:

```bash
make migrate
```