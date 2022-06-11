run:
	docker-compose up -d --remove-orphans --build

test:
	go test ./...

lint:
	golint ./...

migrate-up:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up

migrate-drop:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" drop

migrate-down:
	migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down

.DEFAULT_GOAL := run