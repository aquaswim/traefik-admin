# Traefik Admin

Admin UI for traefik, it mimics the functionality of nginx reverse proxy manager but using traefik.

# How To Run

## Using Docker

You can run Traefik Admin using the pre-built Docker image from GitHub Container Registry:

### Docker Compose

```yaml
version: '3'
services:
  traefik:
    image: "traefik:v3.4"
    restart: unless-stopped
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - traefik-config:/etc/traefik
    command:
      - "--providers.http.endpoint=http://traefik-admin:3000/api/traefik-config/yaml"
      - "--providers.docker=true"
      - "--providers.docker.exposedbydefault=false"
      - "--entrypoints.web.address=:80"
      - "--serverstransport.insecureskipverify=true"
    ports:
      - "80:80"
      - "443:443"
  traefik-admin:
    image: ghcr.io/aquaswim/traefik-admin:latest
    container_name: traefik-admin
    ports:
      - "3000:3000"
    volumes:
      - traefik-admin-data:/config
    restart: unless-stopped

volumes:
  traefik-config:
  traefik-admin-data:
```

### Configuration

The Docker image exposes the following:

- Port: 3000
- Volume: `/config` - stores the database and configuration
- Environment variables:
  - `APP_ADDRESS`: Address to listen on (default: `:3000`)
  - `DB_PATH`: Path to store the database (default: `/config`)
  - `TZ`: Timezone (default: `Etc/UTC`)

## From Source

### Prerequisites

- Go 1.24 or higher
- Node.js LTS (v22) or higher

### Running the Application

```bash
# Download Go dependencies
go mod tidy

# Build the frontend
cd web
npm install
npm run build
cd ..

# Run the application
go run cmd/traefik-admin/main.go
```

The application will start on port 3000. You can access it at http://localhost:3000.

## API Endpoints

### Frontend
- `GET /`: Serves the React frontend application
- `GET /assets/*`: Serves the frontend static assets

### API
- Services:
  - `GET /api/services/`: Get all services
  - `GET /api/services/:id`: Get service by ID
  - `POST /api/services/`: Create a service
  - `PUT /api/services/:id`: Update a service
  - `DELETE /api/services/:id`: Delete a service
- Routes:
  - `GET /api/routes/`: Get all routes
  - `GET /api/routes/:id`: Get route by ID
  - `POST /api/routes/`: Create a route
  - `PUT /api/routes/:id`: Update a route
  - `DELETE /api/routes/:id`: Delete a route
- Traefik Configuration:
  - `GET /api/traefik-config/json`: Get Traefik configuration in JSON format
  - `GET /api/traefik-config/yaml`: Get Traefik configuration in YAML format
