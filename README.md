# 🔗 URL Shortener Service (Go + Gin + Redis)

A production-ready URL Shortener backend built using Go, Gin, and Redis, featuring custom short URLs, expiry support, rate limiting, URL validation, tagging, and redirection handling.

---

## 📌 Features
- Create short URLs
- Custom short IDs
- URL expiry (TTL)
- Rate limiting per IP
- Redirect resolution
- Add tags to URLs
- Edit & delete URLs
- Redis-backed storage

---

## 🧠 Tech Stack
- Go (Golang)
- Gin Web Framework
- Redis
- Docker (optional)
- godotenv
- govalidator

---

## 📁 Project Structure
```
url-shortner/
├── api/
│   ├── database/
│   ├── models/
│   └── routes/
├── main.go
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Environment Variables
```
APP_PORT=8080
DB_ADDR=localhost:6379
DB_PASS=
DOMAIN=http://localhost:8080
API_QUOTA=10
```

---

## 🚀 Running the Project

### Prerequisites
- Go 1.18+
- Redis

### Start Redis
```
redis-server
```

### Run App
```
go mod download
go run main.go
```

---

## 🔌 API Endpoints

### Create Short URL
POST /api/v1

Request:
```json
{
  "url": "https://example.com",
  "short": "myid",
  "expiry": 24
}
```

Response:
```json
{
  "url": "https://example.com",
  "short": "http://localhost:8080/myid",
  "expiry": 24,
  "rate_limit": 9,
  "rate_limit_reset": 29
}
```

---

### Resolve URL
GET /:shortID  
Redirects to original URL.

---

### Edit URL
PUT /api/v1/:shortID

---

### Delete URL
DELETE /api/v1/:shortID

---

### Add Tag
POST /api/v1/addTag

```json
{
  "shortid": "myid",
  "tag": "marketing"
}
```

---

## 👨‍💻 Frontend Notes
All APIs are JSON-based and frontend-ready. Works well with React, Vue, or Next.js.

---

## 📄 License
Open-source.
