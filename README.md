# crud-car

Тестовое CRUD API для сущности машина.

Методы:

`POST /create` создание новой машины
```json
{
	"company": "bmw",
	"model": "x5",
	"capacity": 40
}
```

`POST /list` получение списка машин
```json
{
	"offset": 1,
	"limit": 10
}
```

`POST /get` получение машины по идентификатору
```json
{
    "id": 4
}
```

`POST /update` обновление машины по идентификатору
```json
{
    "id": 4,
	"company": "audi",
	"model": "a30",
	"capacity": 40
}
```

`POST /delete` удаление машины по идентификатору
```json
{
    "id": 4
}
```
