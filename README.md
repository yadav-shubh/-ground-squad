# Go Microservice with GraphQL and gRPC

A high-performance microservice template built with Go, featuring GraphQL and gRPC APIs.

## Features

- **GraphQL API** for flexible data querying
- **gRPC** for high-performance service communication
- **Health Check** endpoints for monitoring
- **Structured Logging** with configurable levels
- **Configuration** via environment variables
- **Docker** support for containerization

## Prerequisites

- Go 1.20 or higher
- Protocol Buffers (protoc)
- Go plugins for protoc:
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
  - `go install github.com/99designs/gqlgen@latest`

## Getting Started

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd base-backend
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Generate code from proto files:
   ```bash
   make proto
   ```

4. Generate GraphQL code:
   ```bash
   make graphql
   ```

5. Build and run the service:
   ```bash
   make run
   ```

## Configuration

Configure the service using environment variables in `.env` file:

```env
PORT=8080
GRPC_PORT=50051
ENVIRONMENT=development
LOG_LEVEL=debug
```

## API Documentation

### GraphQL
- **Playground**: `http://localhost:8080/graphql`
- **Endpoint**: `POST http://localhost:8080/query`

### gRPC
- **Port**: 50051
- **Health Check**: `grpc.health.v1.Health/Check`

## Development

### Building

```bash
make build
```

### Running Tests

```bash
make test
```

### Linting

```bash
make lint
```

### Code Generation

- Generate gRPC code: `make proto`
- Generate GraphQL code: `make graphql`

## Docker

Build the Docker image:

```bash
docker build -t base-backend .
```

Run the container:

```bash
docker run -p 8080:8080 -p 50051:50051 base-backend
```

## Project Structure

```
.
├── cmd/              # Application entry points
├── generated/        # Generated code (gRPC, GraphQL)
├── modules/          # Business logic modules
│   └── health/       # Health check module
│       ├── server/   # gRPC server implementation
│       └── proto/    # Protocol buffer definitions
├── pkg/              # Reusable packages
└── utils/            # Utility functions
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
