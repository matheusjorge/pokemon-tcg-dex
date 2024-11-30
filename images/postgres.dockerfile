FROM postgres:latest

RUN apt update && apt install ca-certificates git build-essential postgresql-common libpq-dev -y
RUN apt install postgresql-17-pgvector -y
