docker_build:
	docker build -t seagal-bot-go .
docker_run:
	docker run --name seagal-bot-go --env-file .env -d seagal-bot-go
docker_start:
	docker container start seagal-bot-go
docker_stop:
	docker container stop seagal-bot-go
