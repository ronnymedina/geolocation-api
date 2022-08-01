build:
	docker-compose up -d --build geo-api


update-doc:
	docker-compose run geo-api swag init -g ./src/main.go -o ./src/docs

