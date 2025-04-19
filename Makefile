local.run:
	go run ./cmd/api/main.go

docker.run:
	docker compose --env-file ./docker.env up -d
docker.run.db:
	docker compose --env-file ./docker.env up -d postgres
docker.run.migrate:
	docker compose --env-file ./docker.env up -d migrate
migrate.up:
	migrate -path ./migrations -database "postgres://myuser:mypassword@localhost:5432/postgres?sslmode=disable" up

