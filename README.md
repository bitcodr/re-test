# Re-Test

## API Documentation to Test Online

#### Update Packs
#### API Endpoint
- **API:** [https://stormy-stream-88610-e6c05773897c.herokuapp.com/packets/update](https://stormy-stream-88610-e6c05773897c.herokuapp.com/packets/update)

#### Request
```json
{
    "packets": [250, 500, 1000, 2000, 5000]
}
```

#### Response
```json
[
    250,
    500,
    1000,
    2000,
    5000
]
```

#### Calculate Order Packs

#### API Endpoint
- **API:** [https://stormy-stream-88610-e6c05773897c.herokuapp.com/orders/calculate](https://stormy-stream-88610-e6c05773897c.herokuapp.com/orders/calculate)

#### Request
```json
{
    "items": 251
}
```

#### Response
```json
{
    "item": 251,
    "packets": {
        "500": 1
    }
}
```

---
## To test on a Local machine, follow the steps as below
## Prerequisites

Before running any of the Makefile targets, make sure you have the following prerequisites:

- Go (Golang) installed on your system.
- Docker and Docker Compose (if you intend to use Docker-related targets).
- The project's directory structure and necessary configuration files are set up as expected.

## Available Makefile Targets

### help

This target displays a list of available Makefile targets and their descriptions.

```bash
make help
```

### release

This target builds a production binary file of application. It uses the `create-env` target to ensure the necessary environment files are present.

```bash
make release
```

### build

This target builds a development binary file of application. Like `release`, it also relies on the `create-env` target.

```bash
make build
```

### run

This target runs application in development mode, automatically reloading when code changes are detected. It depends on `create-env` and uses `CompileDaemon` for automatic reloading.

```bash
make run
```

### create-env

This target creates environment configuration files, such as `.env` and `config.yml`. It ensures that these files exist.

```bash
make create-env
```

### mod

This target retrieves dependency packages using Go Modules (`go mod tidy`).

```bash
make mod
```

### test

This target runs unit tests for application and provides test coverage information. It relies on `tparse` for parsing test results.

```bash
make test
```

### race

This target runs the data race detector on your application to identify data race conditions. Like `test`, it uses `tparse` for parsing test results.

```bash
make race
```

### coverage

This target checks the code coverage of your tests and generates a coverage report.

```bash
make coverage
```

### lint

This target performs code linting on your project to identify potential issues. It uses `golangci-lint` for linting.

```bash
make lint
```

### docs

This target generates and updates documentation using the Swagger tool (`swag`). It parses dependencies for documentation.

```bash
make docs
```

### docker-build

This target builds Docker containers for your application using Docker Compose. Ensure you have Docker and Docker Compose installed.

```bash
make docker-build
```

### docker-up

This target starts your application using Docker Compose in detached mode. It automatically reloads your application within the Docker container.

```bash
make docker-up
```

### docker-down

This target stops the Docker containers created with Docker Compose.

```bash
make docker-down
```

### docker-log

This target prints the logs of Docker containers. It can be useful for debugging when your application is running inside Docker containers.

```bash
make docker-log
```