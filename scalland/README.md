# 

The website of Scalland Conultancy Services

## Quick Start

```bash
# Install dependencies
go mod tidy

# Run migrations
make migrate

# Start server
make run
```

## Commands

| Command | Description |
|---------|-------------|
| `serve` | Start the HTTP server |
| `migrate` | Run database migrations |
| `seed` | Seed the database |
| `upgrade` | Self-upgrade to latest version |
| `version` | Print version information |
| `install` | Install as systemd service |

## Configuration

Configuration files are in `configs/`:
- `app.yml` — Production config
- `app.dev.yml` — Development config
- `app.local.yml` — Local overrides (gitignored)

Environment variables override config values with prefix `_`.

## Build

```bash
make build              # Build binary
make dist               # Build release binaries (linux/amd64 + linux/arm64)
make test               # Run tests
```

## Project Structure

```
├── configs/            # YAML configuration files
├── deployments/        # Dockerfile, docker-compose, build scripts
├── internal/
│   ├── cmd/            # Cobra commands
│   ├── handlers/       # HTTP handlers
│   ├── middleware/      # Auth, CORS, compression middleware
│   ├── model/          # Data models
│   ├── routes/         # Router setup
│   ├── updater/        # Self-upgrade system
│   └── worker/         # Background email worker
├── migrations/         # SQL migration files
├── pkg/
│   ├── log/            # slog-based logger
│   ├── schemas/        # Data schemas
│   └── utils/          # Config, DB, JWT, B2, crypto, filesystem
├── web/template/       # HTML templates and static assets
├── main.go             # Entry point
├── Makefile            # Build targets
└── VERSION             # Semantic version
```

## License

[]()
