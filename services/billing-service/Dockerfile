FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY services/billing-service/go.mod services/billing-service/go.sum ./services/billing-service/
COPY proto/billing/go.mod ./proto/billing/

RUN cd services/billing-service && go mod download

COPY services/billing-service/ ./services/billing-service/
COPY proto/billing/ ./proto/billing/

RUN cd services/billing-service && go build -o api ./cmd/api

FROM alpine:3.22
WORKDIR /app
COPY --from=builder /app/services/billing-service/api ./api
COPY --from=builder /app/services/billing-service/internal/infrastructure/postgres/migrations ./migrations

EXPOSE 8082
CMD ["./api"]