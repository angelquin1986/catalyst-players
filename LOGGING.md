# Logging System

This project implements a comprehensive logging system that supports both development and production environments.

## Features

- **Development Mode**: Logs are output to console (stdout/stderr)
- **Production Mode**: Logs are saved to files in the `logs/` directory
- **Multiple Log Levels**: Info, Error, Debug, and Warning
- **Timestamped Log Files**: Daily log files with date stamps
- **Structured Logging**: Consistent log format across all levels

## Configuration

The logging behavior is controlled by the `APP_ENV` environment variable:

- `APP_ENV=development` (default): Console logging
- `APP_ENV=production`: File logging

## Environment Variables

Add these to your `.env` file:

```env
APP_ENV=development  # or production
LOG_LEVEL=debug
LOG_FORMAT=json
```

## Usage

### In Development Mode

When `APP_ENV=development`, logs will be output to the console:

```
[INFO] 2024/01/15 10:30:45 Retrieving all seasons
[INFO] 2024/01/15 10:30:45 Successfully retrieved 5 seasons
[ERROR] 2024/01/15 10:30:46 Failed to retrieve season with ID 1: record not found
```

### In Production Mode

When `APP_ENV=production`, logs will be saved to files in the `logs/` directory:

```
logs/
├── info-2024-01-15.log
├── error-2024-01-15.log
├── debug-2024-01-15.log
└── warn-2024-01-15.log
```

## Implementation Example

The logging system is implemented in the `SeasonRepositoryImpl` as an example:

```go
// GetAll retrieves all seasons
func (r *SeasonRepositoryImpl) GetAll() ([]entities.Season, error) {
    r.logger.Info("Retrieving all seasons")
    var seasons []entities.Season
    err := r.db.Preload("Teams").Find(&seasons).Error
    if err != nil {
        r.logger.Error("Failed to retrieve all seasons: %v", err)
        return nil, err
    }
    r.logger.Info("Successfully retrieved %d seasons", len(seasons))
    return seasons, err
}
```

## Log Levels

- **Info**: General information about operations
- **Error**: Error conditions and failures
- **Debug**: Detailed debugging information
- **Warn**: Warning messages for potential issues

## Testing

Run the logger tests to verify functionality:

```bash
go test ./internal/infrastructure/logger/
```

## File Structure

```
internal/infrastructure/logger/
├── logger.go          # Main logging implementation
└── logger_test.go     # Tests for logging functionality
```

## Integration

To add logging to other repositories or services:

1. Import the logger package
2. Add a logger field to your struct
3. Initialize the logger in the constructor
4. Add logging statements to your methods

Example:

```go
import "catalyst-players/internal/infrastructure/logger"

type MyRepository struct {
    db     *gorm.DB
    logger logger.Logger
}

func NewMyRepository(db *gorm.DB) MyRepository {
    return &MyRepository{
        db:     db,
        logger: logger.NewLogger(),
    }
}
``` 