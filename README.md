# Bitaksi Hub

A microservices-based application built with Go.

## Services

- **Gateway**: API gateway service
- **Driver Service**: Driver management service

## Prerequisites

- Go 1.x
- Docker & Docker Compose

## Getting Started

```bash
# Start all services
docker-compose up

# Build and run
docker-compose up --build
```

The gateway service runs on port **8080**.

## API Documentation

Swagger UI is available at: `http://localhost:8080/swagger`

## Project Structure

```
bitaksi-hub/
├── gateway/          # API Gateway
├── driver-service/   # Driver management
└── docker-compose.yml
```

## License

MIT
