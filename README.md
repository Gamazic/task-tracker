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
* **Minimal amount of database requests**. Aiming for optimal performance by
minimizing the number of database requests. Some Domain-Driven Design (DDD) or Clean Architecture (CA)
applications suffer from IO performance issues due to excessive querying or
complex domain object reconstruction. This project seeks to strike a balance
that is easily comprehensible for humans while being less IO bound.
* **Minimal dependencies**. Leveraging Go's rich standard library and idiomatic problem-solving
approaches to minimize external dependencies.
The project maximizes the utilization of Go's inherent capabilities.


Still in progress.
