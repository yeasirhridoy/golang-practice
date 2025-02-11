migrateup:
	migrate -path database/migrations -database "postgresql://postgres@localhost:5432/practice?sslmode=disable" -verbose up

migratedown:
	migrate -path database/migrations -database "postgresql://postgres@localhost:5432/practice?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrateup migratedown sqlc