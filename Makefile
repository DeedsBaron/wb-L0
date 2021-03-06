build:
	@echo "\033[0;32mBuilding binary...\033[m"
	@echo $(CURDIR)
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
	 go test -v -count=1 ./... -cover

pub:
	go run nats-streaming-publish/publish.go $(FILE)

.PHONY: all lib clean fclean re

.DEFAULT_GOAL := containers