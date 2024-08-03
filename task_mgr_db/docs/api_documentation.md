## Task Manger API

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
  "1": {
    "id": "1",
    "title": "Task 1",
    "description": "First task Update",
    "due_date": "2024-07-31T23:43:39.6984829+03:00",
    "status": "Pending"
  },
  "2": {
    "id": "2",
    "title": "Task 2",
    "description": "Second task",
    "due_date": "2024-08-01T23:43:39.6984829+03:00",
    "status": "In Progress"
  }
}
```

**endpoint:** `/task/:id` `GET`

**Description**: get tasks with ID = `id`

**Response** if `success`

- status code: **200** - statusSuccess

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "First task Update",
  "due_date": "2024-07-31T23:43:39.6984829+03:00",
  "status": "Pending"
}
```

**Error handling** if `id` not found

    curl --location --globoff 'http://localhost:8080/api/task/43'

- status code: **400** - badRequest

```json
{
  "error": "task with id 43 not found"
}
```

**endpoint:** `/task` `POST`

**Description**: Add Task tasks returns the newly added task

**Response**

```json
{
  "id": "5",
  "title": "New Task",
  "description": "describtion",
  "due_date": "2024-07-31T23:43:39.6984829+03:00",
  "status": "Pending"
}
```

**endpoint:** `/task/:id` `PUT`

**Description**: updates Task with ID = `id`

**Response**

```json
{
  "id": "1",
  "title": "Task 1",
  "description": "First task Update",
  "due_date": "2024-07-31T23:43:39.6984829+03:00",
  "status": "Pending"
}
```

**Error handling** if `id` not found

```json
{
  "error": "task with id 23 not found"
}
```

**Error handling** if ID uri `id` is not equal to the body `id`

```curl
curl --location --request PUT 'http://localhost:8080/api/task/3' \
--header 'Content-Type: application/json' \
--data '{
        "id": "1",
        "title": "Task 1",
        "description": "First task Update",
        "due_date": "2024-07-31T23:43:39.6984829+03:00",
        "status": "Pending"
    }'
```

- status code: **400** - BadRequest

```json
{
  "error": "task ID in request body (1) does not match task ID in URL (3)"
}
```

**endpoint:** `/task/:id` `DELETE`

**Description**: delete Task with ID = `id`

**Response**

- status code: 204 - StatusNoContent

**Error handling** if `id` not found

```json
{
  "error": "task with id 23 not found"
}
```
