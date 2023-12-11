# Task Tracker

Task management system designed to improve task organisation and monitoring.
It allows users to create, manage and track tasks.

## For developers

This project is an example of a supporting backend application in the `Go` language.
It aims to combine high-performance backend with Domain-Driven Design (DDD)
and Clean Architecture techniques.

## Features

* **Clean Architecture**. The architecture prioritizes encapsulating business logic and domain concerns.
Specific implementations of database/web/events/logging/other are not emphasized.
* **High Performance**. Designed with a focus on delivering efficient performance and low latency.
* **Minimal dependencies**. Using Go's rich standard library and idiomatic problem-solving
  approaches to minimize external dependencies. The project maximizes Go's capabilities.
* **Minimal amount of database requests**. Aiming for optimal performance by
minimizing the number of database requests. Some Domain-Driven Design (DDD) or Clean Architecture (CA)
applications suffer from IO performance issues due to excessive querying or
complex domain object reconstruction. This project seeks to strike a balance
that is easily comprehensible for humans while being less IO bound.

## Documentation

### Run

```go
go run ./cmd/rest_api
```

### HTTP Api

Swagger [http://localhost:8080/docs/](http://localhost:8080/docs/)

#### Create user

```shell
curl -X POST 'localhost:8080/api/register' -d '{"username": "MyUsername", "password": "example"}'
```
`{"username":"MyUsername"}`

#### Create task

```shell
curl -X POST 'localhost:8080/api/tasks' -d '{"description": "code"}' -H 'Authorization: Basic ...'
```
`{"task_id":1,"description":"code","stage":"todo"}`


#### Change task stage

```shell
curl -X PATCH 'localhost:8080/api/tasks/1' -d '{"stage": "done"}' -H 'Authorization: Basic ...'
```
`{"task_id":1}`

#### Get tasks of user

```shell
curl -X GET 'localhost:8080/api/tasks' -H 'Authorization: Basic ...'
```
`[{"task_id":1,"description":"code","stage":"todo"}]`


### Inspired by

* https://github.com/Tishka17/deseos17
* https://github.com/SamWarden/user_service

___

Still in progress.
