FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main

FROM alpine:latest
WORKDIR /app
COPY config ./config
COPY --from=builder /app/main .
ENV CONFIG_PATH="/app/config/config_docker.json"
EXPOSE 8080
CMD ["./main"]