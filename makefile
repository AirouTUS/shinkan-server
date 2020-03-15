docker/mysql/seed:
	mysql -h 127.0.0.1 -P 3306 -uroot -proot shinkan < seeds/seed.sql

docker/run:
	@docker-compose -f db/database.yaml up --detach
	-@docker rm shinkan_db_waiter
	@go mod vendor
	@docker-compose up -d

docker/down:
	@docker-compose down
	@docker-compose -f db/database.yaml down
	@rm -rf db/data
