# 🔗 URL Shortener Service (Go + Gin + Redis)

A production-ready URL Shortener backend built using Go, Gin, and Redis.

---

## 🏗 Architecture Diagram

![Architecture Diagram](url_shortener_architecture.png)

---

## 🔌 API Endpoints Table

| Method | Endpoint | Description | Request Body | Response |
|------|---------|-------------|-------------|---------|
| POST | /api/v1 | Create short URL | `{url, short, expiry}` | Short URL + rate limits |
| GET | /:shortID | Redirect to original URL | None | HTTP 301 Redirect |
| GET | /api/v1/:shortID | Get original URL | None | `{ data: url }` |
| PUT | /api/v1/:shortID | Edit short URL | `{url, expiry}` | Success message |
| DELETE | /api/v1/:shortID | Delete short URL | None | Success message |
| POST | /api/v1/addTag | Add tag to URL | `{shortid, tag}` | Updated data |

---

## 🐳 Docker & Docker Compose

### Dockerfile
```dockerfile
FROM golang:1.22-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o url-shortner
EXPOSE 8080
CMD ["./url-shortner"]
```

### docker-compose.yml
```yaml
version: "3.9"
services:
  redis:
    image: redis:7
    ports:
      - "6379:6379"

  app:
    build: .
    ports:
      - "8080:8080"
    env_file:
      - .env
    depends_on:
      - redis
```

---

## 🚀 Run with Docker
```bash
docker compose up --build
```

---

## 📄 License
Open-source.
