GOCMD=go
GOBUILD=$(GOCMD) build
DOCKER_COMPOSE_FILE := docker-compose.yaml

up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down
.PHONY: up down