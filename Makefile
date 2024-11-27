# Variables
DOCKER_COMPOSE_FILE=docker-compose.yaml

.PHONY: up
up:
	@echo "Starting all services with Docker Compose..."
	docker compose -f $(DOCKER_COMPOSE_FILE) up --build

# Stop services with Docker Compose
.PHONY: down
down:
	@echo "Stopping all services..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down

# Clean Docker containers and images
.PHONY: clean
clean:
	@echo "Cleaning Docker containers and images..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down -v --rmi all
	docker system prune -f

# Cleaning up unused Docker containers, images, and networks
.PHONY: prune
prune:
	@echo "Pruning unused Docker containers, images, and networks..."
	docker system prune -f

# Run migrations
.PHONY: migrate-up
migrate-up:
	psql -U user -d orderdb -f migrations/001_create_orders_table.sql

.PHONY: migrate-down
migrate-down:
	# implement

# Run services individually (for local testing without Docker Compose)

.PHONY: run-api
run-api:
	go run cmd/ordersystem/main.go 
	
