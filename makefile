docker/mysql/seed:
	mysql -h 127.0.0.1 -P 3306 -uroot -proot shinkan < seeds/seed.sql

compose/up:
	@docker-compose up -d

compose/down:
	@docker-compose down
	rm -rf db/data
