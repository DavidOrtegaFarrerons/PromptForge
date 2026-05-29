FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o api ./cmd/api

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/api ./api
COPY --from=builder /app/internal/infrastructure/postgres/migrations ./migrations

EXPOSE 8082

CMD ["./api"]