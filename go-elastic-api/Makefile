start:
	docker compose --env-file .env -f docker-compose.yml up -d 

stop:
	docker compose --env-file .env -f docker-compose.yml down

run : 
	go run main.go

test:
	go test -v -cover -short ./...
