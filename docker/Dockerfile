# Phase 1: Build
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Abhängigkeiten kopieren und herunterladen
COPY go.mod go.sum ./
RUN go mod download

# Quellcode kopieren
COPY . .

# Binary bauen
RUN CGO_ENABLED=0 GOOS=linux go build -o forgejo_event_listener main.go

# Phase 2: Finales Image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Binary von Phase 1 kopieren
COPY --from=builder /app/forgejo_event_listener .

# Port (falls benötigt, basierend auf internal/api/server.go)
EXPOSE 8080

CMD ["./forgejo_event_listener", "serve"]
