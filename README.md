# LRU Cache Backend

A simple HTTP-based LRU (Least Recently Used) cache server with TTL support implemented in Go.

## Features

- HTTP API for cache operations (GET/SET)
- LRU eviction policy
- Time-to-live (TTL) support for cached items
- CORS support for cross-origin requests
- Configurable cache capacity

## Project Structure

```
lru_cache_backend/
├── cmd/
│   └── server/         # Application entrypoint
│       └── main.go
├── internal/
│   ├── api/            # HTTP handlers and middleware
│   │   ├── handlers.go
│   │   └── middleware.go
│   └── cache/          # Cache implementation
│       ├── lru.go
│       └── linked_list.go
├── go.mod
├── go.sum
└── README.md
```

## Installation

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/lru_cache_backend.git
   cd lru_cache_backend
   ```

2. Build the application:
   ```
   go build -o cache-server ./cmd/server
   ```

## Configuration

The application can be configured using environment variables:

- `PORT`: Server port (default: 8080)
- `CACHE_CAPACITY`: Maximum number of items in the cache (default: 1024)

## Running the Server

```
./cache-server
```

Or with custom configuration:

```
PORT=3000 CACHE_CAPACITY=2048 ./cache-server
```

## API Endpoints

### Get a value from the cache

```
GET /get?key=example_key
```

Response:
```json
{
  "key": "example_key",
  "value": "example_value"
}
```

### Set a value in the cache

```
POST /set?key=example_key&value=example_value&ttl=60
```

Parameters:
- `key`: Cache key
- `value`: Value to store
- `ttl`: Time-to-live in seconds

Response:
```json
{
  "key": "example_key",
  "value": "example_value"
}
```

## Development

To run the server in development mode:

```
go run ./cmd/server/main.go
```

## License

[MIT](LICENSE)
