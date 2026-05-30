FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY services/content-service/go.mod services/content-service/go.sum ./services/content-service/
COPY proto/billing/go.mod ./proto/billing/

RUN cd services/content-service && go mod download

COPY services/content-service/ ./services/content-service/
COPY proto/billing/ ./proto/billing/

RUN cd services/content-service && go build -o api ./cmd/api

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/services/content-service/api ./api
COPY --from=builder /app/services/content-service/internal/infrastructure/postgres/migrations ./migrations

EXPOSE 8083

CMD ["./api"]