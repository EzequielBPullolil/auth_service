# auth_service

`auth_services` is a Golang-developed REST-API service for user authentication using JWT.

## Description

The auth service aims to be minimalist and highly efficient, providing basic functionalities such as user registration, modification, and deletion, as well as login, logout, and session validation.

It also ensures high data integrity and security.

## Api Endpoints

The service provides the following endpoints for its components, following the REST-API architecture.

| HTTP Verbs | Endpoints         | Action                                         |
| :--------: | ----------------- | ---------------------------------------------- |
|  **GET**   | /v1/users/        | To read user info (need auth token)            |
|  **GET**   | /v1/users/:id     | To read user info by id (need valid token)     |
|  **PUT**   | /v1/users/        | To update an existing user (need auth token)   |
| **DELETE** | /v1/users/        | To delete an existing user (need auth token)   |
|  **POST**  | /v1/auth/signup   | To sign up a new user account                  |
|  **POST**  | /v1/auth/login    | To login an existing user and create JWT token |
|  **GET**   | /v1/auth/validate | To validate user auth token (need auth token)  |

## users endpoints

### Request

`GET /v1/users/` This endpoint requires a valid JWT token

```bash
curl -X GET
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -b "auth_token=<SU_TOKEN>" \
  https://<API_URL>/v1/users
```

### response

```json
    {
        "id": [user_id],
        "name": [user_name],
        "email": [user_email]
    }
```

---

### Request

`GET /v1/users/:id` This endpoints requires a valid application token

```bash
curl -X GET \
 -H "Accept: application/json" \
 -H "Content-Type: application/json" \
 -b "app_token=<SU_TOKEN>" \
 https://<API_URL>/v1/users
```

### response

```json
    {
        "id": [user_id],
        "name": [user_name],
        "email": [user_email]
    }
```

---

### Request

`PUT /v1/users/:id` This endpoints requires a valid JWT token

```bash
curl -X PUT \
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "<name>",
    "password": "<password>",
    "email": "<email>"
  }' \
  https://<API_URL>/v1/users
```

### response

```json
    {
        "id": [user_id],
        "name": [user_name],
        "email": [user_email]
    }
```

---

### Request

`DELETE /v1/users/` This endpoint requires a valid JWT token

```bash
curl -X DELETE
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -b "auth_token=<SU_TOKEN>" \
  https://<API_URL>/v1/users
```

### response

```json
{
  "status": "user deleted"
}
```

---
