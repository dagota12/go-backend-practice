## Task Manger API with MONGO DB

### overview

Task Management REST API using Go programming language and Gin Framework. This API will support basic CRUD operations for managing tasks.

## End Points

BaseURL

    http://localhost:8080/api

**endpoint:** `/task` `GET`

**Description**: Get all tasks

**Response**

- status code: **200** - statusSuccess

```json
{
    {
        "id": "66ad4284dad62d1f3a862cc2",
        "title": "Shopping",
        "description": "buy item",
        "due_date": "2024-07-31T20:43:39.698Z",
        "status": "Completed"
    },
    {
        "id": "06ad4284dad62d1f3a862cc2",
        "title": "Meeting with Clare",
        "description": "meeting",
        "due_date": "2024-07-31T20:43:39.698Z",
        "status": "Completed"
    }
}
```

**endpoint:** `/task/:id` `GET`

**Description**: get tasks with ID = `id`

**Response** if `success`

- status code: **200** - statusSuccess

```json
{
  "id": "06ad4284dad62d1f3a862cc2",
  "title": "Meeting with Clare",
  "description": "meeting",
  "due_date": "2024-07-31T20:43:39.698Z",
  "status": "Completed"
}
```

**Error handling** if `id` not found

    curl --location 'http://localhost:8080/api/task/66adee35d74df484e83593b0'

- status code: **404** - NotFound

```json
{
  "error": "mongo: no documents in result"
}
```

**endpoint:** `/task` `POST`

**Description**: Add Task tasks returns the newly added task
**Request**

```json
{
  "title": "New Task",
  "description": "describtion",
  "due_date": "2024-07-31T23:43:39.6984829+03:00",
  "status": "Pending"
}
```

**Response**

```json
{
  "id": "66ae29b909be0b7a467a5673",
  "title": "New Task",
  "description": "describtion",
  "due_date": "2024-07-31T23:43:39.6984829+03:00",
  "status": "Pending"
}
```

**endpoint:** `/task/:id` `PUT`

**Description**: updates Task with ID = `id`

**Request**

```json
{
  "title": "Task Update",
  "description": "tempora",
  "due_date": "2024-07-31T23:43:39.6984829+03:00",
  "status": "Pending"
}
```

**Response**

```json
{
  "id": "66ad4284dad62d1f3a862cc2",
  "title": "Task Update",
  "description": "tempora",
  "due_date": "2024-07-31T23:43:39.6984829+03:00",
  "status": "Pending"
}
```

**Error handling** if `id` not found

```json
{
  "error": "mongo: no documents in result"
}
```

**endpoint:** `/task/:id` `DELETE`

**Description**: delete Task with ID = `id`

**Request**

    curl --location --request DELETE 'http://localhost:8080/api/task/66ad4284dad62d1f3a862cc2'

**Response**

- status code: 204 - StatusNoContent

**Error handling** if `id` not found

```json
{
  "error": "mongo: no documents in result"
}
```
