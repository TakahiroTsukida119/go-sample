up:
	docker-compose up

build:
	docker-compose build

ps:
	docker ps

stop:
	docker-compose stop

down:
	docker-compose down

shell:
	docker exec -it go-sample-api /bin/sh

nginx:
	docker exec -it go-sample-nginx /bin/sh
