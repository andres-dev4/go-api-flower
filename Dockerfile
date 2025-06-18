# Etapa de construcción
FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG VERSION=dev
RUN go build -ldflags "-X main.Version=$VERSION" -o flower-api .

# Etapa final
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/flower-api .
COPY --from=builder /app/docs ./docs
EXPOSE 8080
CMD ["./flower-api"]
