## todo-list

Todo list made with golang and chi

### How to run

Requirements:

- make
- docker

Start the development container

```bash
make dev
```

### Endpoints

#### Public endpoints

##### POST /register - create a new user

```bash
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Vinicius",
    "email": "vinicius@gmail.com",
    "password": "123"
}'
```

##### POST /login - log in an existent user

```bash
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "vinicius@gmail.com",
    "password": "123"
}'
```

#### Private endpoints

All private endpoints require a bearer token in Authorization header

##### GET /user/{id} - get an existent user

```bash
curl --location 'http://localhost:8080/user/:id' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>'
```

##### POST /task - create a new task

```bash
curl --location 'http://localhost:8080/task' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer <token>' \
--data '{
    "name": "task one",
    "priority": "HIGH",
    "description": "my task description",
    "release_date": "2024-04-15T12:00:00Z"
}'
```
