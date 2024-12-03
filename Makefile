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

front:
	uv run streamlit run tcgdex_front/Homepage.py

build:
	go build -o tcgdex ./src

api: build
	./tcgdex api

docker-build:
	docker compose build

docker-up: docker-build
	docker compose up -d

docker-down:
	docker compose down
