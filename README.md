# Mini Go Fiber Application

A robust Go web application built with Fiber framework, featuring structured logging, PostgreSQL, Redis, and standardized API responses.

## Features

- **Fiber Framework**: High-performance web framework
- **PostgreSQL**: Using sqlx for database operations
- **Redis**: For caching and session management
- **Structured Logging**: Using Uber's Zap logger with file rotation
- **Environment Configuration**: Support for development and production environments
- **Standardized API Responses**: Consistent response format
- **Docker Support**: PostgreSQL and Redis containers
- **Error Handling**: Centralized error handling middleware

## Prerequisites

- Go 1.19 or higher
- Docker and Docker Compose
- Windows OS

## Project Structure

```plaintext
.
├── cmd/
│   └── app/
│       └── main.go
├── internal/
│   ├── config/
│   │   ├── config.go
│   │   └── settings.go
│   ├── middleware/
│   ├── pkg/
│   │   ├── database/
│   │   ├── logger/
│   │   └── util/
│   └── router/
├── logs/
├── .env
├── .gitignore
├── docker-compose.yml
└── README.md