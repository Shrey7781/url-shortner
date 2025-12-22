🚀 URL Shortener ServiceA high-performance, scalable URL shortening microservice built with Go, Gin, Redis, and Docker.📘 1. Project OverviewThis service provides a robust backend for generating and managing short URLs. It leverages Redis for lightning-fast lookups and features built-in rate-limiting and expiration to ensure system stability.Key Features⚡ High-Speed Lookups: Powered by Redis in-memory storage.🛡️ Rate Limiting: IP-based protection against abuse.🏷️ Metadata Support: Custom short IDs and tagging for analytics.⏲️ Automatic Expiry: Set TTL (Time-To-Live) for temporary links.🐳 Containerized: Ready for production with Docker Compose.📂 2. Project ArchitectureThe project follows a clean, modular structure for easy maintenance.Plaintexturl-shortner/
├── api/
│   ├── routes/        # API route handlers (Shorten, Resolve, etc.)
│   ├── database/      # Redis connection logic
│   ├── models/        # Struct definitions
│   └── utils/         # Helper functions
├── main.go            # Application entry point
├── Dockerfile         # Container build instructions
└── docker-compose.yaml
⚙️ 3. Setup & Installation🔑 Environment VariablesCreate a .env file in the root directory and configure the following:VariableDescriptionDefault ValueDOMAINThe base URL for shortened linkshttp://localhost:8080REDIS_ADDRRedis connection addressredis:6379REDIS_PASSWORDPassword for Redis (if any)""RATE_LIMITMax requests per IP per window10🐳 Run with DockerStart the entire stack (API + Redis) with one command:Bashdocker compose up --build
🔌 4. API Documentation✅ 4.1 Create Short URLPOST /api/shortenRequest Body:JSON{
  "url": "https://example.com/long/path",
  "short": "custom123", 
  "expiry": 30
}
Response (200 OK):| Field | Type | Description || :--- | :--- | :--- || url | string | The original long URL || short | string | The generated short link || rate_limit | int | Remaining requests for your IP || rate_limit_reset | int | Minutes until quota resets |🔁 4.2 Resolve URL (Redirect)GET /:shortIDPerforms a 302 Found redirect to the original long URL.🛠️ 4.3 Management EndpointsActionMethodEndpointDescriptionGet InfoGET/api/geturl/:shortIDRetrieve metadata & creation dateEditPUT/api/editurlChange the destination of a short IDDeleteDELETE/api/deleteurl/:shortidRemove a link from the databaseTagPUT/api/addtagAdd a category tag to a link🧑‍💻 5. Frontend Integration GuideTo integrate this service into your frontend, use the standard fetch API:JavaScript// Example: Creating a shortened link
const createShortLink = async (longUrl) => {
    const response = await fetch("http://localhost:8080/api/shorten", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
            url: longUrl,
            expiry: 60 // expires in 1 hour
        })
    });
    return await response.json();
};
🏛️ 6. Data Storage StrategyThe service utilizes two distinct Redis patterns to ensure performance:URL Storage:Key: shortIDValue: longURLTTL: Configured via the expiry field.Rate Limiting:Key: Client_IPLogic: Decrements a counter on every request until it reaches zero.🧪 7. Quick Testing (cURL)Generate a link:Bashcurl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://github.com","expiry":10}'
Check a redirect:Bashcurl -I http://localhost:8080/<YOUR_SHORT_ID>
🤝 ContributingContributions make the open-source community an amazing place to learn, inspire, and create.Fork the ProjectCreate your Feature Branch (git checkout -b feature/AmazingFeature)Commit your Changes (git commit -m 'Add some AmazingFeature')Push to the Branch (git push origin feature/AmazingFeature)Open a Pull Request
