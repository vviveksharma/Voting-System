# Makefile

compose-build:
	docker-compose build

compose-with-debug: compose-build
	@echo "Starting in the debug mode for container"
	@docker compose up 

compose-up: compose-build
	@docker compose up -d

compose-stop:
	@echo "stopping docker compose in background"
	@docker compose down

compose-clean: compose-stop
	docker-compose rm -f

