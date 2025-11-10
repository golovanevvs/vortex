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

NAME = ${REDIS_INSIGHT_CONTAINER}
PROJECT_NAME = ${REDIS_INSIGHT_CONTAINER}
SERVICE_NAME = ${REDIS_INSIGHT_CONTAINER}

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


# ========================
# REDIS_INSIGHT
# ========================

# REDIS_INSIGHT = redisinsight

# .PHONY: redis-insight-up redis-insight-logs redis-insight-stop redis-insight-restart redis-insight-rm redis-insight-down redis-insight-ps redis-insight-exec

# redis-insight-up:
# 	$(MAKE_CMD) -C $(REDIS_INSIGHT) up
# redis-insight-logs:
# 	$(MAKE_CMD) -C $(REDIS_INSIGHT) logs
# redis-insight-stop:
# 	$(MAKE_CMD) -C $(REDIS_INSIGHT) stop
# redis-insight-restart:
# 	$(MAKE_CMD) -C $(REDIS_INSIGHT) restart
# redis-insight-rm:
# 	$(MAKE_CMD) -C $(REDIS_INSIGHT) rm
# redis-insight-down:
# 	$(MAKE_CMD) -C $(REDIS_INSIGHT) down
# redis-insight-ps:
# 	$(MAKE_CMD) -C $(REDIS_INSIGHT) ps
# redis-insight-exec:
# 	$(MAKE_CMD) -C $(REDIS_INSIGHT) exec