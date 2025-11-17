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

PROJECT_NAME = ${CLOUDPUB_LOCAL_CONTAINER}
NAME = ${CLOUDPUB_LOCAL_CONTAINER}
SERVICE_NAME = ${CLOUDPUB_LOCAL_CONTAINER}

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
# CLOUDPUB-LOCAL
# ========================

# CLOUDPUB_LOCAL = cloudpub-local

# .PHONY: cloudpub-local-up cloudpub-local-logs cloudpub-local-stop cloudpub-local-restart cloudpub-local-rm cloudpub-local-down cloudpub-local-ps cloudpub-local-exec

# cloudpub-local-up:
# 	$(MAKE_CMD) -C $(CLOUDPUB_LOCAL) up
# cloudpub-local-logs:
# 	$(MAKE_CMD) -C $(CLOUDPUB_LOCAL) logs
# cloudpub-local-stop:
# 	$(MAKE_CMD) -C $(CLOUDPUB_LOCAL) stop
# cloudpub-local-restart:
# 	$(MAKE_CMD) -C $(CLOUDPUB_LOCAL) restart
# cloudpub-local-rm:
# 	$(MAKE_CMD) -C $(CLOUDPUB_LOCAL) rm
# cloudpub-local-down:
# 	$(MAKE_CMD) -C $(CLOUDPUB_LOCAL) down
# cloudpub-local-ps:
# 	$(MAKE_CMD) -C $(CLOUDPUB_LOCAL) ps
# cloudpub-local-exec:
# 	$(MAKE_CMD) -C $(CLOUDPUB_LOCAL) exec