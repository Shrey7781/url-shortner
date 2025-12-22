# Build stage
FROM golang:alpine AS builder
WORKDIR /app
RUN apk add --no-cache git
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o myapp .

# Run stage
FROM alpine
WORKDIR /app
COPY --from=builder /app/myapp .
EXPOSE 8080
ENTRYPOINT ["./myapp"]
