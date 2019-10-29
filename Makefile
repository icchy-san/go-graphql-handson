up:
	docker-compose up -d

down:
	docker-compose down

restart: stop start

gqlgen:
	docker-compose exec app /bin/bash -c "go generate ./..."

start:
	docker-compose exec app /bin/bash -c "go run server/server.go"
