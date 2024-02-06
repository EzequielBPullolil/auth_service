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
|  **GET**   | /v1/users/:id     | To read user info by id                        |
|  **PUT**   | /v1/users/        | To update an existing user (need auth token)   |
|  **POST**  | /v1/auth/signup   | To sign up a new user account                  |
|  **POST**  | /v1/auth/login    | To login an existing user and create JWT token |
|  **GET**   | /v1/auth/validate | To validate user auth token (need auth token)  |
