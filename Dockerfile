FROM golang:alpine AS builder

WORKDIR /app

COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download

COPY backend/. .

RUN CGO_ENABLED=0 go build -o /server ./cmd/server/main.go

FROM alpine:3.19

WORKDIR /app

RUN apk add --no-cache curl && \
    curl -o /migrate https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64  && \
    chmod +x /migrate

COPY --from=builder /server /server
#COPY --from=builder /app/.env /app/.env
COPY backend/db/migrations /migrations

EXPOSE 8080

CMD ["/server"]