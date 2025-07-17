# Catalyst Players - Soccer Championship Management System

A comprehensive soccer championship management system built with Go, Gin, and GORM using hexagonal architecture. This system allows you to manage players, teams, leagues, seasons, matches, and track tournament progress with support for group stages, knockout rounds, and final tournaments.

## Features

### Core Management
- **Player Management**: Create, update, and manage player profiles with team assignments
- **Team Management**: Manage teams with player rosters and team statistics
- **Stadium Management**: Manage soccer stadiums and venues
- **League Management**: Create and manage soccer leagues
- **Season Management**: Manage tournament seasons with different stages
- **Match Management**: Schedule and track matches with detailed statistics

### Tournament Features
- **Group Stages**: Support for group-based tournaments
- **Knockout Rounds**: Round of 16, quarter-finals, semi-finals
- **Final Tournament**: Championship and third-place matches
- **Standings**: Automatic calculation of team standings
- **Player Statistics**: Track goals, cards, and individual performance
- **Match Results**: Record and update match scores and results

### API Endpoints
- Complete REST API for all entities
- Support for complex queries and relationships
- CORS enabled for frontend integration
- Health check endpoints
- Comprehensive error handling

## Architecture

This project follows **Hexagonal Architecture** (Clean Architecture) principles:

```
┌─────────────────────────────────────────────────────────────┐
│                    Presentation Layer                      │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐        │
│  │   Handlers  │ │   Routes    │ │  Middleware │        │
│  └─────────────┘ └─────────────┘ └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                   Application Layer                        │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐        │
│  │   Services  │ │ Use Cases   │ │ Business    │        │
│  │             │ │             │ │  Logic      │        │
│  └─────────────┘ └─────────────┘ └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                     Domain Layer                           │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐        │
│  │  Entities   │ │Repository   │ │   Domain    │        │
│  │             │ │ Interfaces  │ │  Services   │        │
│  └─────────────┘ └─────────────┘ └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
┌─────────────────────────────────────────────────────────────┐
│                  Infrastructure Layer                      │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐        │
│  │  Database   │ │Repository   │ │   External  │        │
│  │  Connection │ │Implement.   │ │  Services   │        │
│  └─────────────┘ └─────────────┘ └─────────────┘        │
└─────────────────────────────────────────────────────────────┘
```

## Technology Stack

- **Language**: Go 1.21
- **Framework**: Gin (HTTP web framework)
- **ORM**: GORM (Go Object Relational Mapper)
- **Database**: MySQL 8.0
- **Architecture**: Hexagonal Architecture
- **Containerization**: Docker & Docker Compose
- **API**: RESTful API with JSON responses

## Project Structure

```
catalyst-players/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── domain/
│   │   ├── entities/           # Domain entities
│   │   └── repositories/       # Repository interfaces
│   ├── application/
│   │   └── services/           # Business logic services
│   ├── infrastructure/
│   │   ├── database/           # Database connection
│   │   └── repositories/       # Repository implementations
│   └── presentation/
│       ├── handlers/           # HTTP handlers
│       └── routes/             # Route definitions
├── Dockerfile                  # Docker configuration
├── docker-compose.yml          # Production Docker setup
├── docker-compose.dev.yml      # Development Docker setup
├── nginx.conf                  # Nginx configuration
├── env.example                 # Environment variables template
├── go.mod                      # Go module dependencies
└── README.md                   # Project documentation
```

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (for local development)

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd catalyst-players
   ```

2. **Start the application**
   ```bash
   # Production
   docker-compose up -d
   
   # Development
   docker-compose -f docker-compose.dev.yml up -d
   ```

3. **Access the API**
   - Production: http://localhost:8080
   - Development: http://localhost:8081
   - Health Check: http://localhost:8080/api/v1/health

### Local Development

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Set up environment**
   ```bash
   cp env.example .env
   # Edit .env with your database configuration
   ```

3. **Run the application**
   ```bash
   go run cmd/main.go
   ```

## API Documentation

### Base URL
- Production: `http://localhost:8080/api/v1`
- Development: `http://localhost:8081/api/v1`

### Health Check
```
GET /api/v1/health
```

### Core Endpoints

#### Tags
```
POST   /api/v1/tags                    # Create tag
GET    /api/v1/tags                    # Get all tags
GET    /api/v1/tags/:id                # Get tag by ID
PUT    /api/v1/tags/:id                # Update tag
DELETE /api/v1/tags/:id                # Delete tag
```

#### Stadiums
```
POST   /api/v1/stadiums                # Create stadium
GET    /api/v1/stadiums                # Get all stadiums
GET    /api/v1/stadiums/:id            # Get stadium by ID
PUT    /api/v1/stadiums/:id            # Update stadium
DELETE /api/v1/stadiums/:id            # Delete stadium
```

#### Teams
```
POST   /api/v1/teams                   # Create team
GET    /api/v1/teams                   # Get all teams
GET    /api/v1/teams/:id               # Get team by ID
GET    /api/v1/teams/:id/players       # Get team with players
PUT    /api/v1/teams/:id               # Update team
DELETE /api/v1/teams/:id               # Delete team
GET    /api/v1/teams/:id/matches       # Get team matches
GET    /api/v1/teams/:id/match-stats   # Get team match statistics
GET    /api/v1/teams/:id/tags          # Get team tags
```

#### Players
```
POST   /api/v1/players                 # Create player
GET    /api/v1/players                 # Get all players
GET    /api/v1/players/:id             # Get player by ID
GET    /api/v1/players/:id/team        # Get player with team
PUT    /api/v1/players/:id             # Update player
DELETE /api/v1/players/:id             # Delete player
GET    /api/v1/players/:id/match-stats # Get player match statistics
GET    /api/v1/players/:id/stats/:season_id # Get player season stats
GET    /api/v1/players/:id/tags        # Get player tags
```

#### Leagues
```
POST   /api/v1/leagues                 # Create league
GET    /api/v1/leagues                 # Get all leagues
GET    /api/v1/leagues/:id             # Get league by ID
GET    /api/v1/leagues/:id/seasons     # Get league seasons
PUT    /api/v1/leagues/:id             # Update league
DELETE /api/v1/leagues/:id             # Delete league
```

#### Seasons
```
POST   /api/v1/seasons                 # Create season
GET    /api/v1/seasons                 # Get all seasons
GET    /api/v1/seasons/active          # Get active seasons
GET    /api/v1/seasons/:id             # Get season by ID
GET    /api/v1/seasons/:id/league      # Get season with league
GET    /api/v1/seasons/:id/teams       # Get season with teams
GET    /api/v1/seasons/:id/matches     # Get season matches
GET    /api/v1/seasons/:id/matches/completed # Get completed matches
GET    /api/v1/seasons/:id/standings   # Get season standings
GET    /api/v1/seasons/:id/top-scorers # Get top scorers
PUT    /api/v1/seasons/:id             # Update season
PUT    /api/v1/seasons/:id/activate    # Activate season
PUT    /api/v1/seasons/:id/complete    # Complete season
DELETE /api/v1/seasons/:id             # Delete season
```

#### Matches
```
POST   /api/v1/matches                 # Create match
GET    /api/v1/matches                 # Get all matches
GET    /api/v1/matches/upcoming        # Get upcoming matches
GET    /api/v1/matches/date-range      # Get matches by date range
GET    /api/v1/matches/:id             # Get match by ID
GET    /api/v1/matches/:id/details     # Get match with details
GET    /api/v1/matches/:id/players     # Get match player statistics
PUT    /api/v1/matches/:id             # Update match
PUT    /api/v1/matches/:id/score       # Update match score
DELETE /api/v1/matches/:id             # Delete match
GET    /api/v1/matches/:season_id/:stage # Get matches by stage
```

#### Match Players (Statistics)
```
POST   /api/v1/match-players           # Create match player stat
GET    /api/v1/match-players           # Get all match player stats
GET    /api/v1/match-players/:id       # Get match player stat by ID
PUT    /api/v1/match-players/:id       # Update match player stat
DELETE /api/v1/match-players/:id       # Delete match player stat
```

## Database Schema

The system uses the following main entities:

- **Tags**: Categorization for players and teams
- **Stadiums**: Soccer venues
- **Teams**: Soccer teams with players
- **Players**: Individual players with team assignments
- **Leagues**: Soccer leagues
- **Seasons**: Tournament seasons within leagues
- **Matches**: Individual games with scores and statistics
- **Match Players**: Individual player statistics per match

## Environment Variables

Copy `env.example` to `.env` and configure:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=catalyst_soccer
DB_CHARSET=utf8mb4

# Server Configuration
SERVER_PORT=8080
SERVER_MODE=debug

# Application Configuration
APP_NAME=catalyst-players
APP_VERSION=1.0.0
APP_ENV=development

# CORS Configuration
CORS_ALLOWED_ORIGINS=*
CORS_ALLOWED_METHODS=GET,POST,PUT,DELETE,OPTIONS
CORS_ALLOWED_HEADERS=*

# Logging Configuration
LOG_LEVEL=debug
LOG_FORMAT=json
```

## Development

### Running Tests
```bash
go test ./...
```

### Code Formatting
```bash
go fmt ./...
```

### Linting
```bash
golangci-lint run
```

### Building
```bash
go build -o catalyst-players cmd/main.go
```

## Deployment

### Production Deployment
```bash
docker-compose up -d
```

### Development Deployment
```bash
docker-compose -f docker-compose.dev.yml up -d
```

### Scaling
```bash
docker-compose up -d --scale catalyst-api=3
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Support

For support and questions, please open an issue in the repository.

docker-compose -f docker-compose.dev.yml up --build -d