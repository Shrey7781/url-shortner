🚀 URL Shortener (Golang + Gin + Redis)
A fast, scalable URL shortening service built with Go, Gin, Redis, and Docker Compose.

Supports custom short IDs, expiration, tagging, rate-limiting, and analytics.
This README includes:
✔ Full project overview

✔ All API endpoints with JSON request/response

✔ Docker & docker-compose instructions

✔ Redis setup

✔ Frontend integration guide

✔ Folder & architecture overview
📘 1. Overview
This project provides a backend microservice that allows users to:

Generate short URLs
Use custom short codes
Add tags to categorize links
Edit long URLs
Delete shortened URLs
Resolve (/short → original) with redirection
Track rate limits using Redis
Set optional expiration
The service is optimized for:

⚡ High-speed lookups (Redis)
🐳 Full containerization
🌐 Easy frontend integration (JSON APIs)
📂 2. Project Structure

url-shortner/
│
├── api/
│   ├── routes/
│   │   ├── shorten.go
│   │   ├── resolve.go
│   │   ├── getUrl.go
│   │   ├── editUrl.go
│   │   ├── deleteUrl.go
│   │   └── addTag.go
│   ├── database/
│   ├── models/
│   │   └── models.go
│   ├── utils/
│
├── db/
│
├── main.go
├── Dockerfile
├── docker-compose.yaml
├── go.mod
└── .env
📦 3. Environment Variables
Your .env file should contain:


DOMAIN=http://localhost:8080REDIS_ADDR=redis:6379REDIS_PASSWORD=RATE_LIMIT=10
🐳 4. Run With Docker Compose
Start backend + Redis:

docker compose up --build
Services launched:
ServicePortDescriptionGo API8080Main serverRedis6379Storage DB
Backend is now available at:


http://localhost:8080
🔌 5. API Endpoints (Complete Documentation)
✅ 5.1 Create Short URL
POST /api/shorten
Request Body

{
  "url": "https://example.com/long/path",
  "short": "custom123", 
  "expiry": 30}
FieldTypeDescriptionurlstringOriginal long URLshortstring (optional)Custom short codeexpirynumber (minutes)Expiration time
Response

{
  "url": "https://example.com/long/path",
  "short": "http://localhost:8080/custom123",
  "expiry": 30,
  "rate_limit": 9,
  "rate_limit_reset": 30}
The backend applies Redis-based rate limiting per IP.
🔁 5.2 Resolve Short URL (Redirect)
GET /:shortID
Example:


GET /custom123
Server returns:


302 FoundLocation: https://example.com/long/path
🔍 5.3 Get URL Metadata
GET /api/geturl/:shortID
Response:


{
  "shortid": "custom123",
  "url": "https://example.com/long/path",
  "expiry": 30,
  "created_at": "2025-01-01T08:00:00Z"}
✏️ 5.4 Edit URL
PUT /api/editurl
Request


{
  "shortid": "custom123",
  "url": "https://newsite.example/updated"}
Response


{
  "message": "URL updated successfully"}
🗑 5.5 Delete URL
DELETE /api/deleteurl/:shortid
Response


{
  "message": "URL deleted successfully"}
🏷 5.6 Add Tag to URL
PUT /api/addtag
Request


{
  "shortid": "custom123",
  "tag": "marketing"}
Response


{
  "message": "Tag added successfully",
  "tag": "marketing"}
🧰 6. Models (from your code)
Request Model (shorten)

type Request struct {
    URL         string `json:"url"`
    CustomShort string `json:"short"`
    Expiry      int    `json:"expiry"`
}
Tag Request

type TagRequest struct {
    ShortID string `json:"shortid"`
    Tag     string `json:"tag"`
}
Response

type Response struct {
    URL             string `json:"url"`
    CustomShort     string `json:"short"`
    Expiry          int    `json:"expiry"`
    XRateRemaining  int    `json:"rate_limit"`
    XRateLimitReset int    `json:"rate_limit_reset"`
}
🧑‍💻 7. Frontend Developer Guide
Base URL:

http://localhost:8080
Example: Creating short URL (JavaScript)

const res = await fetch("http://localhost:8080/api/shorten", {
  method: "POST",
  headers: { "Content-Type": "application/json" },
  body: JSON.stringify({
    url: "https://google.com",
    short: "ggl",
    expiry: 60
  })
});console.log(await res.json());
Example: Redirect user

window.location.href = "http://localhost:8080/ggl";
🏛 8. Redis Usage (from your code)
Redis is used for:

URL storage
TTL expiration
Rate limiting per IP
Quick resolving
Each entry:


key: shortIDvalue: longURLTTL: expiry minutes
Rate limit:


key: <Client IP>value: remaining_requestsTTL: limit reset time
🧪 9. Testing with cURL
Create URL

curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://github.com","expiry":10}'
Resolve URL

curl -I http://localhost:8080/abc123
Add Tag

curl -X PUT http://localhost:8080/api/addtag \
  -H "Content-Type: application/json" \
  -d '{"shortid":"abc123","tag":"social"}'
🤝 10. Contributing
PRs and issues are welcome.
