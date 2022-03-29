build:
	@echo "\033[0;32mBuilding binary...\033[m"
	go build -o wb-L0 -v ./cmd/wb-L0/main.go

run: build
	./wb-L0

containers:
	docker-compose build
	docker-compose up -d --force-recreate

logs:
	docker-compose logs

clean:
	-docker-compose down
	-docker volume rm $$(docker volume ls -q)
	-rm -rf spam/spam

exec:
	docker exec -it nats-streaming sh

status:
	docker ps -a

test:
	@$(MAKE) test -s -C spam

pub:
	go run nats-streaming-publish/publish.go

.PHONY: all lib clean fclean re

.DEFAULT_GOAL := containers