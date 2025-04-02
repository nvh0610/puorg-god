GO_CMD = go
include .env
export $(shell sed 's/=.*//' .env)


run:
	$(GO_CMD) run main.go

create-migration:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(name)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE__TYPE)://$(DATABASE__USER):$(DATABASE__PASSWORD)@tcp($(DATABASE__HOST):$(DATABASE__PORT))/$(DATABASE__NAME)" up

migrate-down:
	migrate -path $(MIGRATION_DIR) -database "$(DATABASE__TYPE)://$(DATABASE__USER):$(DATABASE__PASSWORD)@tcp($(DATABASE__HOST):$(DATABASE__PORT))/$(DATABASE__NAME)" down

docker-run:
	docker run --name $(MYSQL_CONTAINER_NAME) -e MYSQL_ROOT_PASSWORD=$(DATABASE__PASSWORD) -e MYSQL_DATABASE=$(DATABASE__NAME) -p $(DATABASE__PORT):3306 -d mysql:latest
	docker run --name $(REDIS_CONTAINER_NAME) -p $(REDIS_PORT):6379 -d redis:latest

docker-stop:
	docker stop $(MYSQL_CONTAINER_NAME)
	docker rm $(MYSQL_CONTAINER_NAME)
	docker stop $(REDIS_CONTAINER_NAME)
	docker rm $(REDIS_CONTAINER_NAME)


setup:
	make docker-run
	@echo "Waiting for MySQL to be ready..."
	@until docker exec $(MYSQL_CONTAINER_NAME) mysql -u$(DATABASE__USER) -p$(DATABASE__PASSWORD) -e "SELECT 1" > /dev/null 2>&1; do \
		echo "Waiting for MySQL..."; \
		sleep 5; \
	done
	make migrate-up

teardown:
	make migrate-down
	make docker-stop
