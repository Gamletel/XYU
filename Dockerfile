FROM golang:alpine AS builder

WORKDIR /app

COPY backend/go.mod .
COPY backend/go.sum .

RUN go mod download

COPY backend/. .

RUN CGO_ENABLED=0 go build -o /server cmd/server/main.go

FROM alpine:3.19

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /server /server
#COPY --from=builder /app/config/ ./config/
#COPY --from=builder /app/migrations/ ./migrations/

EXPOSE 8080

CMD ["/server"]