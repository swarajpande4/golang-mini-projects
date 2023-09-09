## JWT Authentication Example in Golang

This repository contains the source code for a JWT authenticaion example in golang.

### Project dependencies:
- [gin-gonic/gin](https://github.com/gin-gonic/gin) framework for building web applications.
- [gorm.io/gorm](https://gorm.io/index.html) ORM library for golang.
- [gorm.io/driver/mysql](https://gorm.io/docs/connecting_to_the_database.html) Gorm based driver for MySQL.
- [dgrijalva/jwt-go](https://github.com/golang-jwt/jwt) JWT related functionality for golang.
- [joho/godotenv](https://github.com/joho/godotenv) library for managing environment variables.

### Local setup:

- Install go1.21.0 on your workspace

- Run the following commands to download modules and run the application
    ```
    cd app
    go mod download
    go run main.go
    ```

### APIs
- For registering the user
    ```
    curl -X POST "http://localhost:8080/api/user/register" -H "Content-Type: application/json" -d '{"name": "<NAME>", "username": "<USERNAME>", "email": "<EMAIL>", "password": "<PASSWORD>"}'
    ```

- For creating the token 
    ```
    curl -X POST "http://localhost:8080/api/token" -H "Content-Type: application/json" -d '{"email": "<EMAIL>", "password": "<PASSWORD>"}'

    ```

- For accessing the dummy api (authentication required)
    ```
    curl -X GET "http://localhost:8080/api/secured/ping" -H "Content-Type: application/json" -H "Authorization: <Bearer Token>"

    ```


### References:
- https://codewithmukesh.com/blog/jwt-authentication-in-golang/
