include .env
export $(shell sed 's/=.*//' .env)

# Build da imagem
docker_build:
	docker build -t $(IMAGE_NAME) -f $(DOCKERFILE_PATH) .

# Rodar o contêiner
docker_run:
	docker run -d -p $(PORT):$(PORT) --name $(IMAGE_NAME)-container $(IMAGE_NAME)

# Parar o contêiner
docker_stop:
	docker stop $(IMAGE_NAME)-container || true
	docker rm $(IMAGE_NAME)-container || true

# Rebuild e rodar
docker_up: stop build run

# Mostrar logs do contêiner
docker_logs:
	docker logs -f $(IMAGE_NAME)-container

# Limpar imagem e contêiner
docker_clean: stop
	docker rmi $(IMAGE_NAME) || true

# Roda migração
migrate:
	go run ./cmd/migrate/main.go

# Roda a api
run:
	go run ./cmd/api/main.go

