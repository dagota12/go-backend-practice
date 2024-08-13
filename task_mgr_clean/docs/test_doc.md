## How to Run test

this documentation shows how to run tests for the task managment api

### Requiremets

- install [Go](https://go.dev/doc/install) 1.22.5 in your machine
- [mongodb](https://www.mongodb.com/docs/manual/installation/) for storing the data install it in your locally or create [atlas](https://www.mongodb.com/docs/manual/installation/) account for and store it in the .env file as `DB_URI`

#### 1. Clone the Repository:

```sh
git clone https://github.com/dagota12/go-backend-practice
cd go-backend-practice
cd task_mgr_clean
```

#### 2. to run all tests at once

```sh
 go test ./...
```

- to run speccific test file or package

  ```sh
   go test ./<folder-name>/tests/
   go test ./<folder-name>/tests/<test-file>
  ```
