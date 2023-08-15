## User Management APIs

This repository contains the source code for building a user management application (APIs)
with Golang using the Gin-gonic framework and MongoDB.

### Project dependencies:
- [gin-gonic/gin](https://github.com/gin-gonic/gin) framework for building web applications.
- [mongo-driver/mongo](https://go.mongodb.org/mongo-driver/mongo) driver for connecting to MongoDB.
- [joho/godotenv](https://github.com/joho/godotenv) library for managing environment variables.
- [go-playground/validator/v10](https://github.com/go-playground/validator) library for validating structs and fields.

### Local setup:

- Install go1.21.0 on your workspace

- Run the following commands to download modules and run the application
    ```
    cd app
    go mod download
    go run main.go
    ```

- To run go container with volume while running,
    ```
    podman build --target dev . -t go
    podman run -it -v ${PWD}:/work:z go sh
    go version
    go run app.go
    go build
    ./app
    ```

- To build and run the container,
    ```
    podman build . -t app
    dock run app
    ```

### References:
- https://github.com/Mr-Malomz/gin-mongo-api
- https://youtu.be/jpKysZwllVw
- https://pkg.go.dev/cmd/go#hdr-Workspace_maintenance
