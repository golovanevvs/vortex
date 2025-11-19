MAKEFLAGS += --no-print-directory

ENV_FILE ?= .env
ifneq (,$(wildcard $(ENV_FILE)))
include $(ENV_FILE)
export
endif

ifeq ($(OS),Windows_NT)
	MAKE_CMD = mingw32-make
else
	MAKE_CMD = make
endif

.PHONY: up logs stop restart rm down ps exec

NAME = ${KAFKA_CONTAINER}
PROJECT_NAME = ${KAFKA_CONTAINER}
SERVICE_NAME = ${KAFKA_CONTAINER}

up:
	@echo "Starting ${NAME}..."
	docker compose -p $(PROJECT_NAME) up -d
	@echo "${NAME} is up and running."

logs:
	docker compose -p $(PROJECT_NAME) logs -f

stop:
	@echo "Stopping ${NAME}..."
	docker compose -p $(PROJECT_NAME) stop
	@echo "${NAME} stopped."

restart: stop up

rm:
	@echo "Removing ${NAME} (keeping data)..."
	docker compose -p $(PROJECT_NAME) rm -f
	@echo "${NAME} removed."

down:
	@echo "Full cleanup ${NAME}: containers, volumes, networks, data..."
	docker compose -p $(PROJECT_NAME) down -v --remove-orphans
	@echo "${NAME} cleaned."

ps:
	docker compose -p $(PROJECT_NAME) ps

exec:
	docker compose -p $(PROJECT_NAME) exec $(SERVICE_NAME) sh

create-topic:
	@echo "Creating Kafka topic: $(TOPIC)"
	docker exec $(KAFKA_CONTAINER) kafka-topics --create \
        --topic $(TOPIC) \
        --bootstrap-server localhost:9092 \
        --partitions 1 \
        --replication-factor 1 \
        --if-not-exists
	@echo "Topic $(TOPIC) created."

# ========================
# KAFKA
# ========================

# KAFKA = kafka

# .PHONY: kafka-up kafka-logs kafka-stop kafka-restart kafka-rm kafka-down kafka-ps kafka-exec kafka-create-topic

# kafka-up:
# 	$(MAKE_CMD) -C $(KAFKA) up
# kafka-logs:
# 	$(MAKE_CMD) -C $(KAFKA) logs
# kafka-stop:
# 	$(MAKE_CMD) -C $(KAFKA) stop
# kafka-restart:
# 	$(MAKE_CMD) -C $(KAFKA) restart
# kafka-rm:
# 	$(MAKE_CMD) -C $(KAFKA) rm
# kafka-down:
# 	$(MAKE_CMD) -C $(KAFKA) down
# kafka-ps:
# 	$(MAKE_CMD) -C $(KAFKA) ps
# kafka-exec:
# 	$(MAKE_CMD) -C $(KAFKA) exec
# kafka-create-topic:
# 	$(MAKE_CMD) -C $(KAFKA) create-topic TOPIC="$(TOPIC)"
