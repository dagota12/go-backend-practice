## Task Manger API with MONGO DB and JWT - Auth

### overview

Task Management REST API using Go programming language and Gin Framework. This API will support basic CRUD operations for managing tasks.

also includes basic role based authentication system where users can register and login to acces the resource based on their role

### outline

- [tasks](#tasks)
- [users](#users)

## End Points

## tasks

BaseURL

    http://localhost:8080/

**endpoint:** `api/task` `GET`

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

**endpoint:** `api/task/:id` `GET`

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

**endpoint:** `api/task` `POST`

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

**endpoint:** `api/task/:id` `PUT`

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

**endpoint:** `api/task/:id` `DELETE`

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

## users

### register

**endpoint:** `api/user/register` `POST`

**Description**: register user, only admin can register a user if there is no user registerd the first to register himself becomes admin

- Input form data
  - username: String, reqired
  - pasword: String, required

**Request**

```sh
curl --location 'http://localhost:8080/api/user/register' \
--header 'Content-Type: application/json' \
--data '{
    "username":"alex",
    "password":"fakepwd"
}p'
```

**response**

- returns the newly inserted user

```json
{
  "id": "String",
  "username": "String"
}
```

**Error Handling**

```
{
    "error": string
}
```

### login

**endpoint:** `api/user/login` `POST`

**Description**: user login to the system

- Input form data
  - username: String, reqired
  - pasword: String, required

**Request**

```sh
  curl --location 'http://localhost:8080/api/user/login' \
  --form 'username="bob"' \
  --form 'password="fakepwd"'
```

**response**

- returns an access token for the current user

```json
{
  "token": "String"
}
```

**Error Handling**

```json
{
  "error": "String"
}
```
