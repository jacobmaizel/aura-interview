FROM golang:latest

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
# RUN apt-get update && apt-get install -y postgresql-client

COPY go.mod go.sum ./
RUN go mod download

# COPY ./scripts/ ./scripts/
# COPY ./migrations/ ./migrations/

# RUN chmod +x ./scripts/run_migrations.sh
# RUN ./scripts/run_migrations.sh


CMD ["air", "-c", ".air.toml"]
