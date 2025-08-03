.PHONY: build up down clean

build:
	@echo -e "Building image...\n"
	docker build -t elasticsearch-training-api:latest .

up: build
	@echo -e "Running images...\n"
	docker compose up -d

down:
	@echo -e "Stopping services...\n"
	docker compose down

clean:
	@echo -e "Stopping, removing containers.\Removing networks.\nRemoving volumes.\n"
	docker compose down -v
