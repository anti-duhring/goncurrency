DB_HOST = localhost
DB_USER = postgres
DB_NAME = postgres
DB_PASSWORD = postgres

dev:
	nodemon --exec DB_HOST=$(DB_HOST) DB_USER=$(DB_USER) DB_NAME=$(DB_NAME) DB_PASSWORD=$(DB_PASSWORD) go run main.go --signal SIGTERM

up:
	docker compose down
	docker compose up 

build:
	go build -o api .
