.PHONY: up down logs

up:
		docker compose up -d
run: 
		go run main.go
build: 
		go build main.go && chmod +x main && ./main
down:
		docker-compose down
logs:
		docker-compose logs -f