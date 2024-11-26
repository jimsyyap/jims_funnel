# Full Stack Web Application

## Prerequisites
- Go 1.21+
- Node.js 18+
- Docker (optional)
- PostgreSQL

## Local Development

### Backend Setup
1. Navigate to backend directory
2. Copy `.env.example` to `.env`
3. Update database credentials

### Frontend Setup
1. Navigate to frontend directory
2. Install dependencies: `npm install`

## Running the Application

### Without Docker
- Run backend: `make backend-run`
- Run frontend: `make frontend-run`

### With Docker
```bash
# Build and start services
docker-compose up --build

# Stop services
docker-compose down
```

## Environment Variables
- `DB_HOST`: Database host
- `DB_USER`: Database username
- `DB_PASSWORD`: Database password
- `DB_NAME`: Database name
- `PORT`: Server port
