# Techcasts Backend Readme

## Overview

Welcome to the backend repository of Techcasts, an open-source music and podcast streaming service built with Go. This README will guide you through setting up the backend environment, understanding the project structure, and contributing to its development.

## Table of Contents

1. [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Configuration](#configuration)
2. [Project Structure](#project-structure)
3. [Running the Backend](#running-the-backend)
4. [API Documentation](#api-documentation)
5. [Contributing](#contributing)
6. [License](#license)

## Getting Started

### Prerequisites

Make sure you have the following installed on your machine:

- Go (version 1.16 or higher)
- PostgreSQL (or any other supported database)
- Redis (for caching purposes)

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/techcasts-backend.git
    cd techcasts-backend
    ```

2. Install dependencies:

    ```bash
    go mod tidy
    ```

### Configuration

1. Copy the example configuration file:

    ```bash
    cp .env.example .env
    ```

2. Update the `.env` file with your database and Redis credentials.

## Project Structure

The project follows a standard Go project structure:

- `cmd`: Contains the application entry point.
- `internal`: Holds the internal packages, including the application logic and database models.
- `migrations`: Stores database migration files.
- `pkg`: Contains reusable packages.
- `scripts`: Includes utility scripts.
- `tests`: Contains tests for the application.

## Running the Backend

To run the backend locally, use the following commands:

```bash
go run cmd/techcasts/main.go
```

By default, the server will start at `http://localhost:8080`.

## API Documentation

The API documentation is generated using Swagger. After running the backend, you can access the Swagger UI at `http://localhost:8080/swagger/index.html`.

## Contributing

If you want to contribute to Techcasts, follow these steps:

1. Fork the repository.
2. Create a new branch: `git checkout -b feature/your-feature`.
3. Make your changes and commit: `git commit -m "Add your feature"`.
4. Push to the branch: `git push origin feature/your-feature`.
5. Open a pull request.

## License

Techcasts is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute the code as per the terms of the license.