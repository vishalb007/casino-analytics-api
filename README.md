# Casino Analytics API

## Features
- MongoDB Aggregation Pipelines
- Redis Caching
- Gin REST API
- 2M+ dataset support

## Setup

```bash
cp .env.example .env
docker-compose up -d
go run scripts/seed.go
go run cmd/server/main.go