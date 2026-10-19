# Weather API with Redis

A REST API built with Go that retrieves weather timeline data from the [Visual Crossing Weather API](https://www.visualcrossing.com/weather-api/) and uses Redis caching to reduce redundant external API requests and improve response times.

## Features

- **Weather API Integration** — Fetch weather timeline data from Visual Crossing.
- **Redis Caching** — Cache weather responses to avoid repeated upstream API calls.
- **IP-Based Rate Limiting** — Limit incoming requests to help protect API resources from excessive usage.
- **HTTP Client Handling** — Handle requests to third-party services and network errors.
- **Environment Configuration** — Manage API credentials and Redis settings through environment variables.
- **Docker Compose** — Run Redis locally with persistent storage and a health check.
- **Health Check Endpoint** — Check whether the HTTP server is responding.

## Tech Stack

- **Language:** Go
- **HTTP Server:** Go `net/http`
- **Cache:** Redis 7
- **Redis Client:** `github.com/redis/go-redis/v9`
- **Weather Provider:** Visual Crossing Weather API
- **Configuration:** Environment variables and `godotenv`
- **Containerization:** Docker Compose
- **Rate Limiting:** `golang.org/x/time/rate`

## Prerequisites

- [Go 1.26 or later](https://go.dev/dl/)
- [Docker](https://www.docker.com/)
- A [Visual Crossing API key](https://www.visualcrossing.com/weather-api/)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/Permanaa/weather-api-with-redis.git
cd weather-api-with-redis
```

### 2. Configure environment variables

Create a local environment file:

```bash
cp .env.example .env
```

Configure the variables in `.env`:

```env
VISUAL_CROSSING_WEATHER_API_KEY=your_api_key
BASE_URL=your_weather_api_base_url

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

| Variable | Description |
|---|---|
| `VISUAL_CROSSING_WEATHER_API_KEY` | API key for Visual Crossing |
| `BASE_URL` | Base URL used for weather API requests |
| `REDIS_ADDR` | Redis server address |
| `REDIS_PASSWORD` | Redis password, if authentication is enabled |
| `REDIS_DB` | Redis database index |

Keep your actual API key in `.env`. Never commit credentials to version control.

### 3. Start Redis

Start the Redis service using Docker Compose:

```bash
docker compose up -d
```

Verify that Redis is ready:

```bash
docker exec -it weather-redis redis-cli ping
```

Expected output:

```text
PONG
```

### 4. Install Go dependencies

```bash
go mod download
```

### 5. Run the API

```bash
go run ./cmd/server
```

The server listens on port `1804`.

## API Endpoints

### Health Check

Check whether the server is responding.

```http
GET /healthz
```

Example:

```bash
curl -i http://localhost:1804/healthz
```

### Get Weather Timeline

Retrieve weather timeline data for a location.

```http
GET /timeline/{location}
```

Example:

```bash
curl -i "http://localhost:1804/timeline/London%2CUK"
```

URL-encode location values when they contain reserved characters. The actual accepted location format and response schema depend on the weather provider integration.

## How Redis Caching Works

The API uses Redis to cache weather data and reduce redundant requests to Visual Crossing.

1. A client requests weather data for a location.
2. The application checks Redis for a cached result.
3. If cached data is available and valid, the application can return it without another upstream request.
4. If the cache entry is missing, the application fetches weather data from Visual Crossing.
5. The response is stored in Redis for subsequent requests, according to the application's cache expiration policy.

This cache-aside approach can reduce latency, lower external API consumption, and improve efficiency for repeated requests.

## Rate Limiting

The API applies IP-based rate limiting through middleware to control incoming request frequency.

Rate limiting helps reduce excessive requests and protect application resources. Requests that exceed the configured limit are rejected according to the middleware's response behavior.

The exact rate, burst allowance, and HTTP response details are defined in the rate limiter implementation.

## Inspect Redis Cache

Connect to Redis:

```bash
docker exec -it weather-redis redis-cli
```

List keys in the current database:

```redis
SCAN 0
```

Inspect a cached value:

```redis
GET "London,UK"
```

The actual key format and stored value type depend on the cache implementation.

Exit the Redis CLI:

```redis
exit
```

## Stop the Services

Stop Redis:

```bash
docker compose down
```

To remove the Redis volume and its persisted data as well:

```bash
docker compose down -v
```

**Warning:** Removing the volume deletes the data stored in it.

## Project Structure

```text
weather-api-with-redis/
├── cmd/
│   └── server/       # Application entry point
├── internal/
│   ├── config/       # Application configuration
│   ├── db/           # Redis connection
│   ├── handler/      # HTTP handlers
│   ├── helper/       # Cache helpers
│   └── middleware/   # IP rate limiting
├── .env.example      # Environment template
├── compose.yml       # Redis service configuration
├── go.mod
├── go.sum
└── README.md
```

## Learning Objectives

This project demonstrates practical backend development concepts:

- Integrating third-party REST APIs with Go.
- Managing HTTP requests and network timeouts.
- Implementing Redis caching.
- Applying middleware-based IP rate limiting.
- Handling configuration and secrets through environment variables.
- Running supporting infrastructure with Docker Compose.
- Debugging API and cache behavior.
