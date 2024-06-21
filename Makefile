-include .env.docker

docker-compose:
	docker-compose --project-name backend-cassandra-go --env-file .env.docker ${COMMAND}
docker-up:
	make docker-compose COMMAND="up -d"
docker-down:
	make docker-compose COMMAND="down --remove-orphans"
docker-exec:
	make docker-compose COMMAND="exec -it go-cli ${COMMAND}"
docker-shell:
	make docker-exec COMMAND="sh"
docker-build:
	make docker-compose COMMAND="build --pull"
go:
	make docker-exec COMMAND="go run ./app/*.go"