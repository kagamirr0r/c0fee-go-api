# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Running the Application
- Start server: `go run cmd/server/main.go` (runs on port 8080)
- Start services: `docker compose up` (PostgreSQL on 5432, MinIO on 9000/9001)

### Database Operations
- Run migrations: `go run cmd/db/migrate/main.go up`
- Check migration status: `go run cmd/db/migrate/main.go status`
- Create new migration: `goose create -dir infrastructure/db/migrations {{migration_name}} go`
- Seed database: `go run cmd/db/seed/main.go` (interactive - choose table or "all")

### Build and Dependencies
- Install dependencies: `go mod tidy`
- Build: `go build cmd/server/main.go`

## Architecture Overview

This is a Go REST API using Echo framework with Clean Architecture principles:

**Layer Structure:**
- `cmd/` - Application entry points (server, database tools)
- `controller/` - HTTP handlers (presentation layer)
- `usecase/` - Business logic layer
- `repository/` - Data access layer
- `model/` - Domain entities
- `infrastructure/` - External concerns (database, S3, migrations)
- `router/` - Route definitions and middleware

**Key Components:**
- **Database**: PostgreSQL with GORM ORM, Goose migrations
- **Storage**: MinIO (S3-compatible) for file storage
- **Authentication**: JWT-based with custom middleware
- **Validation**: go-playground/validator with custom validators

**Dependency Flow**: Router → Controller → UseCase → Repository → Database
- Each layer depends only on interfaces from the layer below
- Repository interfaces are defined in usecase layer
- Main.go wires all dependencies together

**Domain Entities**: Users, Beans, Countries, Areas, Farms, Farmers, Roasters, Varieties, ProcessMethods, BeanRatings

**Database**: Uses UUIDs as primary keys, foreign key relationships between entities, separate migration files for each table.

**File Upload**: Handles user avatars and bean images via MinIO S3 service, organized by entity type and ID.