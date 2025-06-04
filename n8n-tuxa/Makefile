.PHONY: network volume start-infra stop-infra clear-infra

network:
	docker network create n8n-network

volume:
	docker volume create n8n_data
	docker volume create postgres_volume

start-infra:
	docker compose --env-file .env -f docker-compose.yml up -d 

stop-infra:
	docker compose --env-file .env -f docker-compose.yml down

clear-infra:
	docker network rm n8n-network
	docker volume rm n8n_data
	docker volume rm postgres_volume
