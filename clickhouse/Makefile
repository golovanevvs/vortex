.PHONY: config
config:
		rm -rf clickhouse01 clickhouse02 clickhouse03 clickhouse04
		mkdir -p clickhouse01/data clickhouse01/logs clickhouse01/config.d
		mkdir -p clickhouse02/data clickhouse02/logs clickhouse02/config.d
		mkdir -p clickhouse03/data clickhouse03/logs clickhouse03/config.d
		mkdir -p clickhouse04/data clickhouse04/logs clickhouse04/config.d
		REPLICA=01 SHARD=01 envsubst < config.xml > clickhouse01/config.d/config.xml
		REPLICA=02 SHARD=01 envsubst < config.xml > clickhouse02/config.d/config.xml
		REPLICA=03 SHARD=02 envsubst < config.xml > clickhouse03/config.d/config.xml
		REPLICA=04 SHARD=02 envsubst < config.xml > clickhouse04/config.d/config.xml
		cp users.xml clickhouse01/config.d/users.xml
		cp users.xml clickhouse02/config.d/users.xml
		cp users.xml clickhouse03/config.d/users.xml
		cp users.xml clickhouse04/config.d/users.xml

.PHONY: up
up:
	docker-compose up -d

.PHONY: start
start:
	docker-compose start

.PHONY: stop
stop:
	docker-compose stop

.PHONY: create-db
create-db:
	docker exec clickhouse01 clickhouse-client --query="CREATE DATABASE IF NOT EXISTS clickhouse_course"

.PHONY: list-db
list-db:
	docker exec clickhouse01 clickhouse-client --query="SHOW DATABASES"

.PHONY: down
down:
	docker-compose down