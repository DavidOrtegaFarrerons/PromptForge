module github.com/DavidOrtegaFarrerons/promptforge/services/content-service

go 1.26.2

require (
	github.com/golang-migrate/migrate/v4 v4.19.1
	github.com/google/uuid v1.6.0
	github.com/jackc/pgx/v5 v5.9.2
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/net v0.51.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.34.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260226221140-a57be14db171 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

require (
	github.com/DavidOrtegaFarrerons/promptforge/proto/billing v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.81.1
)

replace github.com/DavidOrtegaFarrerons/promptforge/proto/billing => ../../proto/billing
