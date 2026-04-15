# HNG Stage 1 - API

A simple Go API for HNG DevOps Stage 1.

## Endpoints

- `GET /` → Returns `{"message": "API is running"}`
- `GET /health` → Returns `{"message": "healthy"}`
- `GET /me` → Returns your my information

## How to run locally

```bash
go run main.go