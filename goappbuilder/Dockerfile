# ---------- Build stage ----------
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install git (needed for some Go modules)
RUN apk add --no-cache git

# Copy dependency files first (better caching)
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o app


# ---------- Runtime stage ----------
FROM alpine:3.19

WORKDIR /app

# Add CA certificates (important for DB / HTTPS)
RUN apk add --no-cache ca-certificates

# Copy only the binary from builder
COPY --from=builder /app/app .

# Expose port if your app listens on one (optional)
EXPOSE 8080

CMD ["./app"]

