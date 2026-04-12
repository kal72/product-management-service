# Product management service

## Description

This is a project built with Golang Clean architecture.

## Tech Stack

- Golang v1.26
- MySQL v8.4
- Redis v8.6.2

## Framework & Library

- GoFiber (HTTP Framework) : https://github.com/gofiber/fiber
- GORM (ORM) : https://github.com/go-gorm/gorm
- Viper (Configuration) : https://github.com/spf13/viper
- Golang Migrate (Database Migration) : https://github.com/golang-migrate/migrate
- Go Playground Validator (Validation) : https://github.com/go-playground/validator
- Logrus (Logger) : https://github.com/sirupsen/logrus

## Configuration

All configuration is in `config.json` file.

## API Spec

The **Postman Collection** for testing all endpoints is located in the `api` directory. You can import it directly into your Postman workspace.

JSON Response status:
| Status        | Description                   | HTTP Status code|
| ------------- |:-----------------------------:| --------------- |
| 99            | Internal server error         | 500             |
| 00            | Operation success             | 200             |
| 01            | Data not found in database    | 200             |
| 04            | request validation error      | 400             |


## Database Migration

All database migration scripts are located in the `db/migrations` folder. We use [golang-migrate](https://github.com/golang-migrate/migrate) to handle schema changes.

To create a new migration sequence:
```shell
migrate create -ext sql -dir db/migrations new_migration_name
```

To execute all pending migrations (Up):
```shell
migrate -database "mysql://root:password@tcp(localhost:3306)/db_name" -path db/migrations up
```

To rollback the last applied migration (Down):
```shell
migrate -database "mysql://root:password@tcp(localhost:3306)/db_name" -path db/migrations down
```

## Run Application

### 1. Infrastructure Setup (MySQL & Redis)

A `docker-compose.yml` file is already provided to easily run the required MySQL and Redis background services.
To start them, run the following command from the project root:
```shell
docker-compose up -d
```

### 2. Starting the Server

Once the databases are up and running, execute the command below to start the application:
```shell
go run cmd/web/main.go
```
Check application running in browser
```shell
http://localhost:8080/ping
```