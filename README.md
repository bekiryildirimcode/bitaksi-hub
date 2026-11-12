# bitaksi-hub

Small **Go microservices** demo for a ride-hailing style backend:

- **gateway** — Echo API gateway with JWT auth, request logging, and per-IP rate limiting; reverse-proxies authenticated traffic to the driver service; Swagger UI
- **driver-service** — Driver CRUD + nearby-driver lookup over MongoDB (haversine distance)

Everything runs with Docker Compose.

## Architecture

```
Client → :8080 gateway → driver-service:8000 → MongoDB
              │
         /login (JWT)
         /swagger
         /api/v1/* (JWT + proxy)
```

## Stack

- Go, Echo
- MongoDB 7
- Docker / Docker Compose
- OpenAPI / Swagger UI

## Quick start

```bash
docker compose up --build
```

- Gateway: http://localhost:8080  
- Swagger: http://localhost:8080/swagger  

## Auth

`POST /login` with JSON body (`firstName`, `lastName`) returns a JWT.  
Authenticated driver routes go through `/api/v1/*` on the gateway.

## Driver service routes

Mounted behind the gateway under `/api/v1/` (proxied to driver-service `/api/v1/driver`):

| Method | Path | Purpose |
| --- | --- | --- |
| `GET` | `/api/v1/driver` | List drivers (pagination) |
| `POST` | `/api/v1/driver` | Create driver |
| `GET` | `/api/v1/driver/nearby` | Nearby drivers |
| `PUT` | `/api/v1/driver/:id` | Update driver |

## Project structure

```
bitaksi-hub/
├── gateway/           # JWT gateway + reverse proxy + Swagger
├── driver-service/    # Driver domain (handler → service → repository)
└── docker-compose.yml
```

## Local development (without Docker)

Run MongoDB, then each service from its directory (`go run .`) with the env vars from `docker-compose.yml` (`MONGO_URI`, `DRIVER_SERVICE_URL`).

## License

MIT
