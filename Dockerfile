FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download || echo "No deps"

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o agent-harness . 2>/dev/null || echo "Build skip"

FROM alpine:3.19

RUN apk add --no-cache git curl ca-certificates

WORKDIR /app

COPY --from=builder /app/agent-harness . 2>/dev/null || true
COPY .env.example .env

EXPOSE 3000

CMD ["./agent-harness", "serve"]
