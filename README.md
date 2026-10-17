# Weather API with Redis

A weather API built with Go that retrieves weather data from the [Visual Crossing Weather API](https://www.visualcrossing.com/weather-api/) and uses Redis to cache responses for faster subsequent requests.

This project demonstrates backend development concepts including third-party API integration, Redis caching, environment-based configuration, HTTP client handling, and Docker-based service setup.

## Features

- **Weather Data Integration** — Fetch weather information from the Visual Crossing Weather API.
- **Redis Caching** — Cache weather responses to reduce unnecessary requests to the external API.
- **Environment Configuration** — Manage API credentials and Redis connection settings through environment variables.
- **Docker Support** — Run Redis locally using Docker Compose.
- **Go Backend** — Built with Go and the `net/http` package for HTTP request handling.
- **Error Handling** — Handle external API and Redis errors.

## Tech Stack

- **Language:** Go
- **Caching:** Redis
- **Redis Client:** `github.com/redis/go-redis/v9`
- **Weather Provider:** Visual Crossing Weather API
- **Configuration:** Environment variables and `godotenv`
- **Containerization:** Docker and Docker Compose

## Prerequisites

Make sure you have the following installed:

- [Go 1.26 or later](https://go.dev/dl/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop/) or Docker Engine
- A [Visual Crossing Weather API key](https://www.visualcrossing.com/weather-api/)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Permanaa/weather-api-with-redis.git
cd weather-api-with-redis
```

### 2. Configure environment variables

Create a `.env` file from the example:

```bash
cp .env.example .env
```

Update the values in `.env`:

```env
VISUAL_CROSSING_WEATHER_API_KEY=your_api_key
BASE_URL=your_weather_api_base_url

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

| Variable | Description |
|---|---|
| `VISUAL_CROSSING_WEATHER_API_KEY` | API key used to authenticate requests to Visual Crossing |
| `BASE_URL` | Base URL used for weather API requests |
| `REDIS_ADDR` | Redis server address |
| `REDIS_PASSWORD` | Redis password, if authentication is enabled |
| `REDIS_DB` | Redis database index |

Use the appropriate base URL expected by the application. Keep your actual API key in `.env` and never commit secrets to version control.

### 3. Start Redis

Start the Redis service using Docker Compose:

```bash
docker compose up -d
```

The provided Compose configuration uses Redis 7 Alpine, exposes Redis on `127.0.0.1:6379`, enables append-only persistence, and defines a health check.

Verify that Redis is running:

```bash
docker exec -it weather-redis redis-cli ping
```

Expected output:

```text
PONG
```

### 4. Install dependencies

```bash
go mod download
```

### 5. Run the API

Start the application using its server entry point:

```bash
go run ./cmd/server
```

Check the application logs for startup errors. The HTTP port and available routes depend on the server configuration.

## API Usage

The API accepts a location and uses Redis to cache weather data for subsequent requests.

For example, a location such as `London,UK` can be used as a weather lookup key.

Once the server is running, send a request to the weather endpoint using the route and query parameter names defined by the application.

Example request pattern:

```http
GET /weather?location=London%2CUK
```

> **Note:** Confirm the actual route, query parameter name, and server port in the source code before using this example. It illustrates the intended request pattern rather than a verified endpoint contract.

### Example Workflow

1. A client requests weather data for a location.
2. The API checks Redis for a cached response.
3. If a valid cache entry exists, the API can return the cached data.
4. If the entry is missing, the API requests fresh weather data from Visual Crossing.
5. The response is cached in Redis for subsequent requests.

This caching strategy can reduce external API calls and improve response times for repeated location lookups.

## Inspect Redis Cache

Connect to Redis using the command-line client:

```bash
docker exec -it weather-redis redis-cli
```

List keys in the current database:

```redis
SCAN 0
```

Inspect a specific key:

```redis
GET "London,UK"
```

The actual key format depends on the application's cache implementation. If the response is stored under a different key or serialization format, adjust the command accordingly.

Exit the Redis CLI:

```redis
exit
```

## Stop Redis

Stop the Docker Compose services:

```bash
docker compose down
```

The named Redis volume is retained, so cached data may persist across container restarts. To remove the persisted volume as well, run:

```bash
docker compose down -v
```

**Warning:** Removing the volume deletes the data stored in that Docker volume.

## Project Structure

```text
weather-api-with-redis/
├── cmd/
│   └── server/       # Application entry point
├── internal/         # Application implementation
├── .env.example      # Environment variable template
├── compose.yml       # Redis service configuration
├── go.mod            # Go module and dependencies
├── go.sum            # Dependency checksums
└── README.md
```

## Learning Objectives

This project explores several backend engineering concepts:

- Integrating a Go application with a third-party REST API.
- Managing HTTP requests and handling network errors.
- Connecting Go applications to Redis.
- Implementing a cache-aside pattern.
- Configuring applications using environment variables.
- Running supporting infrastructure with Docker Compose.
- Inspecting and debugging cached data.
