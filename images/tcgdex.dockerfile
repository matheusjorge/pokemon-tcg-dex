FROM golang:1.22 AS builder

WORKDIR /app

COPY src/ src
COPY go.mod go.sum Makefile .

RUN make build

FROM python:3.11.10-slim-bookworm

RUN apt update && apt install build-essential -y

RUN pip install uv

WORKDIR app

COPY pyproject.toml uv.lock Makefile .

COPY --from=builder /app/tcgdex .

COPY image_embedding_sidecar/ image_embedding_sidecar
COPY tcgdex_front/ tcgdex_front
RUN touch README.md

ENTRYPOINT /bin/bash
