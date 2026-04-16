# 🎰 Casino Analytics API

A high-performance REST API built in Go that processes millions of casino transaction records using advanced MongoDB aggregation pipelines.

This project demonstrates modern backend engineering practices including scalable data modeling, efficient querying, caching strategies, and clean architecture.

---

## 🚀 Features

* ⚡ **MongoDB Aggregation Pipelines** for high-performance analytics
* 📊 Handles **2M+ transactions across 500+ users**
* 🧠 **Redis caching** for expensive queries
* 🔐 **Authorization middleware** for admin routes
* ✅ **Input validation** using go-playground/validator
* 🧱 Clean architecture (Handler → Service → Repository)
* 🐳 Dockerized local development setup

---

## 📊 API Endpoints

### 1. Gross Gaming Revenue (GGR)

```http
GET /gross_gaming_rev?from=YYYY-MM-DD&to=YYYY-MM-DD
```

Returns revenue grouped by currency.

---

### 2. Daily Wager Volume

```http
GET /daily_wager_volume?from=YYYY-MM-DD&to=YYYY-MM-DD
```

Returns daily wager totals per currency.

---

### 3. User Wager Percentile

```http
GET /user/{user_id}/wager_percentile?from=YYYY-MM-DD&to=YYYY-MM-DD
```

Returns user ranking percentile based on total wagered USD.

---

## 🏗️ Architecture

```
Handler → Service → Repository → MongoDB
                ↓
              Redis
```

* **Handlers**: HTTP layer
* **Service**: Business logic + caching
* **Repository**: MongoDB aggregation queries

---

## 🗄️ Data Model

Each game round includes:

* At least one **Wager**
* At least one **Payout**
* Same `roundId` and `currency`
* Chronologically valid timestamps

---

## ⚙️ Tech Stack

* **Go (Golang)**
* **MongoDB**
* **Redis**
* **Gin Web Framework**
* **MongoDB Go Driver**
* **Validator (go-playground)**

---

## 🐳 Local Setup

### 1. Clone Repo

```bash
git clone https://github.com/<your-username>/casino-analytics-api.git
cd casino-analytics-api
```

### 2. Configure Environment

```bash
cp .env.example .env
```

### 3. Start Services

```bash
docker-compose up -d
```

### 4. Seed Database

```bash
go run scripts/seed.go
```

### 5. Run API

```bash
go run cmd/server/main.go
```

---

## 🔐 Authentication

All endpoints require:

```http
Authorization: Bearer SECRET_ADMIN_KEY
```

---

## ⚡ Performance Considerations

* Aggregations executed **inside MongoDB**
* Indexed fields: `createdAt`, `userId`, `roundId`
* Redis caching reduces repeated heavy queries
* Batch inserts for efficient seeding

---

## 📈 Scaling Strategy

* Horizontal scaling with stateless API
* MongoDB indexing + potential sharding
* Pre-aggregation for heavy analytics workloads
* Redis TTL-based cache invalidation

---

## 🧪 Future Improvements

* MongoDB `$setWindowFields` for optimized percentile calculation
* Swagger / OpenAPI documentation
* Rate limiting & structured logging
* Background jobs for precomputed metrics
* Load testing with k6

---

## 💡 Key Highlights

This project focuses on:

* Writing **efficient MongoDB aggregation queries**
* Designing **scalable backend systems**
* Handling **large datasets (millions of records)**
* Applying **clean code and architecture principles**

---

## 📜 License

MIT
