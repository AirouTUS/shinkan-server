DB_CONTAINER := shinkan_db
DB_PASSWORD := root
DB_ENTRY := /docker-entrypoint-initdb.d

mysql/init:
	docker-compose -f db/database.yaml up --detach
	docker rm shinkan_db_waiter

mysql/seed:
	mysql -h 127.0.0.1 -P 3306 -uroot -proot shinkan < seeds/seed.sql

mysql/down:
	docker-compose -f db/database.yaml down
	@rm -rf db/data

compose/up:
	@go mod vendor
	docker-compose up -d

compose/down:
	docker-compose down

docker/build/api:
	@go mod vendor
	docker build -t shinkan-server -f './cmd/api/Dockerfile' .

docker/build/admin:
	@go mod vendor
	docker build -t shinkan-admin-server -f './cmd/admin/Dockerfile' .

docker/run: mysql/init mysql/seed compose/up

docker/down: compose/down mysql/down

test/init: mysql/init
	-@printf '\033[32m%s\033[m\n' '----- run test ---------'
	@docker exec -it $(DB_CONTAINER) bin/bash -c \
	'mysql -u root --password="$(DB_PASSWORD)" < $(DB_ENTRY)/test_ddl.sql;\
	mysql -u root --password="$(MYSQL_PASSWORD)" < $(DB_ENTRY)/test_dml.sql'\
	> /dev/null

test/down:
	@docker exec -it $(DB_CONTAINER) bin/bash -c \
	'mysql -u root --password="$(DB_PASSWORD)" -e "DROP DATABASE IF EXISTS shinkan_test"'\
	> /dev/null

go/test:
	-@go test -cover ./...

test: test/init go/test test/down
