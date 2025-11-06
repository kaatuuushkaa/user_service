DB_DSN := postgres://postgres:yourpassword@localhost:5432/postgres?sslmode=disable
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new-users:
	migrate create -ext sql -dir ./migrations users

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

lint:
	golangci-lint run --color=always