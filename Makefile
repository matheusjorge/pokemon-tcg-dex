compose-up:
	docker compose up -d

migrate-up:
	migrate -path=src/internal/repositories/migrations -database "postgresql://ash:pikachu@0.0.0.0:5432/pokedex?sslmode=disable" -verbose up

migrate-down:
	migrate -path=src/internal/repositories/migrations -database "postgresql://ash:pikachu@0.0.0.0:5432/pokedex?sslmode=disable" -verbose down

go-run:
	go run src/main.go

sidecar:
	uv run python image_embedding_sidecar

build:
	go build -o tcgdex ./src
