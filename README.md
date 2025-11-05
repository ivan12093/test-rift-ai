# Word of Wisdom TCP Server

TCP server "Word of Wisdom" with DDoS protection through Proof of Work (HashCash) algorithm.

## Description

The server implements a challenge-response protocol for DDoS protection:
1. Server sends a random challenge and difficulty level to the client
2. Client solves the Proof of Work task (finding a nonce where the hash starts with a certain number of zeros)
3. Client sends the solution back to the server
4. Server verifies the solution and sends a random quote from the collection

## Architecture

The project is implemented using **Clean Architecture** and standard Go project structure:

```
.
├── cmd/                    # Application entry points
│   ├── server/            # TCP server
│   └── client/            # TCP client
├── internal/              # Internal packages (not exported)
│   ├── domain/            # Domain layer (business logic)
│   │   ├── entity/        # Entities (Quote, Challenge, POWResult)
│   │   ├── repository/    # Repository interfaces
│   │   └── service/       # Service interfaces
│   ├── application/       # Application layer (use cases)
│   │   └── usecase/       # Business scenarios
│   ├── infrastructure/    # Infrastructure layer (implementations)
│   │   ├── repository/    # Repository implementations
│   │   └── service/       # Service implementations
│   └── presentation/      # Presentation layer (interfaces)
│       ├── server/        # TCP server and handlers
│       ├── client/        # TCP client
│       └── protocol/      # Message protocol
├── config/                # Application configuration
├── quotes.txt            # File with wisdom quotes
├── Dockerfile.server      # Dockerfile for server
├── Dockerfile.client      # Dockerfile for client
└── docker-compose.yml     # Docker Compose configuration
```

### Architecture Layers:

1. **Domain Layer** - application core, doesn't depend on external libraries
   - Entities: Quote, Challenge, POWResult
   - Interfaces: QuoteRepository

2. **Application Layer** - business scenarios (use cases)
   - GetQuoteUseCase
   - GenerateChallengeUseCase
   - VerifyPOWUseCase
   - SolvePOWUseCase
   - Interfaces: POWService (defined where it's used)

3. **Infrastructure Layer** - external dependency implementations
   - FileQuoteRepository - working with quotes file
   - HashCashPOW - Proof of Work implementation

4. **Presentation Layer** - interaction interfaces
   - TCP Server/Client handlers
   - Protocol (parsing and formatting messages)

## Requirements

- Docker and Docker Compose
- Go 1.21+ (for local development only)

## Running with Docker Compose

```bash
# Start server and client
docker-compose up --build

# Start server only
docker-compose up server

# Run client in separate terminal
docker-compose run client
```

## Local Development

### Requirements
- Go 1.21+
- Docker and Docker Compose (for containerization)

### Server

```bash
# Install dependencies
go mod download

# Run server
go run cmd/server/main.go

# Or with environment variables
export PORT=8080
export DIFFICULTY=20
export QUOTES_FILE=quotes.txt
export TIMEOUT_SECONDS=30
go run cmd/server/main.go
```

The server will listen on port `8080`.

### Client

## Protocol

### Message Format

**Challenge (server → client):**
```
CHALLENGE:<challenge>:<difficulty>\n
```

**Solution (client → server):**
```
SOLUTION:<challenge>:<nonce>\n
```

**Quote (server → client):**
```
QUOTE:<quote>\n
```

**Error (server → client):**
```
ERROR:<error description>\n
```

## Difficulty Configuration

POW difficulty is configured via environment variable `DIFFICULTY`:
```go
const DIFFICULTY = 20 // Number of zero bits at the start of hash
```

- Lower value = easier solution = faster response
- Higher value = harder solution = better DDoS protection

## Adding Quotes

Edit the `quotes.txt` file, adding one quote per line.

## Example Workflow

```
Server: CHALLENGE:abc123def456:20
Client: [solving POW...]
Client: SOLUTION:abc123def456:12345
Server: QUOTE:Patience is the art of hoping.
```

## Environment Variables

### Server
- `PORT` - server port (default: 8080)
- `DIFFICULTY` - POW difficulty in bits (default: 20)
- `QUOTES_FILE` - path to quotes file (default: quotes.txt)
- `TIMEOUT_SECONDS` - timeout for POW solution in seconds (default: 30)

### Client
- `SERVER_ADDR` - server address (default: server:8080)

## Testing

The project includes unit tests for critical components:

### Run all tests

```bash
go test ./...
```

### Run tests with coverage

```bash
go test -cover ./...
```

### Run tests for specific package

```bash
# POW service tests
go test ./internal/infrastructure/service/...

# Use cases tests
go test ./internal/application/usecase/...

# Protocol tests
go test ./internal/presentation/protocol/...

# Repository tests
go test ./internal/infrastructure/repository/...
```

### Test Coverage

- **POW Service** - testing challenge generation, solution and POW verification
- **Use Cases** - testing business logic with mocks
- **Protocol** - testing message parsing and formatting
- **Repository** - testing quotes file operations
- **Entities** - testing domain entities
- **Config** - testing configuration loading

## Security

- Each challenge is unique and used only once
- Solution timeout - 30 seconds
- Proof of Work requires computational resources, protecting from simple DDoS attacks
- Difficulty can be adjusted depending on load