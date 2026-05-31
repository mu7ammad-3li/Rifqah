# Rifqah: Privacy-First Voice Support Platform

This platform enables secure, privacy-first peer support groups with centralized SFU and edge-based audio processing.

## Project Structure

- `backend/`: Go backend services (Auth, Signaling, State Management).
- `mobile/`: Flutter mobile application (Audio Pipeline, UI/UX).
- `infrastructure/`: Infrastructure configuration (PostgreSQL, Redis, LiveKit).
- `docs/`: Project documentation and plans.

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.21+
- Flutter (for mobile development)

### Running the Infrastructure

```bash
docker-compose up -d
```

### Running the Backend

```bash
cd backend
go run cmd/api/main.go
```
