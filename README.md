# auth_service

`auth_services` is a Golang-developed REST-API service for user authentication using JWT.

## Description

The auth service aims to be minimalist and highly efficient, providing basic functionalities such as user registration, modification, and deletion, as well as login, logout, and session validation.

It also ensures high data integrity and security.

## Api Endpoints

The service provides the following endpoints for its components, following the REST-API architecture.

| HTTP Verbs | Endpoints      | Action                                         |
| :--------: | -------------- | ---------------------------------------------- |
|  **GET**   | /users/        | To read user info (need auth token)            |
|  **GET**   | /users/:id     | To read user info by id (need valid token)     |
|  **PUT**   | /users/        | To update an existing user (need auth token)   |
| **DELETE** | /users/        | To delete an existing user (need auth token)   |
|  **POST**  | /auth/signup   | To sign up a new user account                  |
|  **POST**  | /auth/login    | To login an existing user and create JWT token |
|  **GET**   | /auth/validate | To validate user auth token (need auth token)  |

## users endpoints entry

### Request

`GET /users/` This endpoint requires a valid JWT token

```bash
curl -X GET
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -b "auth_token=<app_token>" \
  https://<API_URL>/users
```

### response

```json
{
  "id": "<user_id>",
  "name": "<user_name>",
  "email": "<user_email>"
}
```

---

### Request

`GET /users/:id` This endpoints requires a valid application token

```bash
curl -X GET \
 -H "Accept: application/json" \
 -H "Content-Type: application/json" \
 -b "app_token=<app_token>" \
 https://<API_URL>/users/<some_id>
```

### response

```json
{
  "id": "<user_id>",
  "name": "<user_name>",
  "email": "<user_email>"
}
```

---

### Request

`PUT /users/` This endpoints requires a valid app token

```bash
curl -X PUT \
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -b "app_token=<app_token>; auth_token=<auth_token>;" \
  -d '{
    "name": "<name>",
    "password": "<password>",
    "email": "<email>"
  }' \
  https://<API_URL>/users
```

### response

```json
{
  "id": "<user_id>",
  "name": "<user_name>",
  "email": "<user_email>"
}
```

---

### Request

`DELETE /users/` This endpoint requires a valid JWT token

```bash
curl -X DELETE
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -b "auth_token=<auth_token>; app_token=<app_token>" \
  https://<API_URL>/users
```

### response

```json
{
  "status": "user deleted"
}
```

## auth endpoints entry

#| **POST** | /auth/signup | To sign up a new user account | /
#| **POST** | /auth/login | To login an existing user and create JWT token | /
#| **GET** | /auth/validate | To validate user auth token (need auth token) |/

### Request

`POST /auth/singup` This endpoints requires a valid app_token

```bash
curl -X POST \
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -b "app_token=<app_token>" \
  -d '{
    "name": "<name>",
    "password": "<password>",
    "email": "<email>"
  }' \
  https://<API_URL>/auth/singup
```

### response

```json
{
  "status": "Successful user registration",
  "message": "Waiting for email validation",
  "email_validation_link": "<email_validation_link>",
  "data": {
    "user": {
      "id": "<user_id>",
      "name": "<user_name>",
      "email": "<user_email>"
    }
  }
}
```

---

### Request

`POST /auth/login` This endpoints requires a valid app_token

```bash
curl -X POST \
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -b "app_token=<app_token>;" \
  -d '{
    "email": "<email>",
    "password": "<password>"
  }' \
  https://<API_URL>/auth/login
```

### response

```json
{
  "status": "Successful user login",
  "data": {
    "token": "<auth_token>",
    "user": {
      "id": "<user_id>",
      "name": "<user_name>",
      "email": "<user_email>"
    }
  }
}
```

---

### Request

`POST /auth/validate` This endpoints requires a valid app_token and auth_token

```bash
curl -X POST \
  -H "Accept: application/json" \
  -H "Content-Type: application/json" \
  -b "app_token=<app_token>; auth_token=<auth_token>"
  https://<API_URL>/auth/validate
```

### response

#### Valid auth_token

```json
{
  "status": "Valid auth token",
  "data": {
    "user": {
      "id": "<user_id>",
      "name": "<user_name>",
      "email": "<user_email>"
    }
  }
}
```

### Invalid auth_token

```json
{
  "status": "Invalid auth token",
  "message": "It may be that the auth_token does not exist or has expired, try to re-authenticate your session"
}
```

---

## Field validation rules

Both the `PUT /users` and `POST /auth/signup` endpoints have field validation, which means that the `name`, `email` and `password` fields are validated using the following rules:

| Field        | validation criteria                                  | Description                                                 |
| ------------ | ---------------------------------------------------- | ----------------------------------------------------------- |
| **Name**     | Length > 5                                           | Ensures names are reasonably long enough.                   |
| **Email**    | is available                                         | Prevents duplicate email addresses.                         |
|              | Matches a valid email address format                 | Guarantees correct email formatting for communication.      |
| **Password** | Length > 7                                           | Enforces a minimum password strength against basic attacks. |
|              | Includes at least one number                         | Adds complexity and makes passwords harder to guess.        |
|              | Includes at least one number                         | Adds complexity and makes passwords harder to guess.        |
|              | Includes at least one symbol (special character)     | Further increases password complexity for added security.   |
|              | Includes at least one uppercase and lowercase letter | Makes passwords more resistant to brute-force attacks.      |
