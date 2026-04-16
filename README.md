# HNG Stage 1 - API

A simple Go API for HNG DevOps Stage 1.

## Endpoints

- `GET /` Returns `{"message": "API is running"}`
- `GET /health` Returns `{"message": "healthy"}`
- `GET /me` Returns your my information

All endpoints return `Content-Type: application/json` and HTTP status **200**.

## How to run locally

```bash
# Clone the repo
git clone https://github.com/kurocifer/hng-stage1-api.git
cd hng-stage1-api

# Run the API
go run main.go
```

The API will be available at http://localhost:8080

## Live Deployment
Live URL: http://kurocifer-hng-one.duckdns.org