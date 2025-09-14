# Configuration Guide

This application supports configuration through environment variables and `.env` files.

## Configuration Options

### Database Configuration

Set these environment variables:

- `DB_HOST` - Database host (default: `localhost`)
- `DB_PORT` - Database port (default: `5432`)
- `DB_USER` - Database username (default: `postgres`)
- `DB_PASSWORD` - Database password (default: `password`)
- `DB_NAME` - Database name (default: `todo_db`)
- `DB_SSLMODE` - SSL mode (default: `disable`)

### Application Configuration

- `LOG_LEVEL` - Logging level (`debug`, `info`, `warn`, `error`) (default: `info`)

## Using .env Files

1. Copy the example file:
   ```bash
   cp .env.example .env
   ```

2. Edit the `.env` file with your configuration:
   ```env
   # Database Configuration
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_password
   DB_NAME=todo_db
   DB_SSLMODE=disable

   # Application Configuration
   LOG_LEVEL=info
   ```

3. The application will automatically load the `.env` file when it starts.

## Environment Variables Priority

Configuration is loaded in the following order (later values override earlier ones):

1. Default values
2. `.env` file (if it exists)
3. System environment variables

## Example Configurations

### Local Development
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=todo_dev
DB_SSLMODE=disable
LOG_LEVEL=debug
```

### Production
```env
DB_HOST=prod-db.example.com
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=todo_prod
DB_SSLMODE=require
LOG_LEVEL=info
```


## Connection Pool Settings

The database connection pool is configured with:
- Maximum connections: 10
- Minimum connections: 1

These settings can be modified in the `internal/db/db.go` file if needed.

