.PHONY:build up


build:
	docker compose -f "docker-compose.yaml" up -d --build
	
up:
	docker compose up -d