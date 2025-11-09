MAKEFLAGS += --no-print-directory

PROJECT_DIR := $(abspath ../..)

ENV_FILES := .env ../../../../../.env
$(foreach file,$(ENV_FILES), \
  $(if $(wildcard $(file)), $(eval include $(file))) \
)
export

ifeq ($(OS),Windows_NT)
	MAKE_CMD = mingw32-make
else
	MAKE_CMD = make
endif

.PHONY: up logs stop restart rm down ps exec

NAME = ${GRAFANA_CONTAINER}
PROJECT_NAME = ${GRAFANA_CONTAINER}
SERVICE_NAME = ${GRAFANA_CONTAINER}

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
# GRAFANA
# ========================

# GRAFANA = providers/grafana

# grafana-up:
# 	$(MAKE_CMD) -C $(GRAFANA) up
# grafana-logs:
# 	$(MAKE_CMD) -C $(GRAFANA) logs
# grafana-stop:
# 	$(MAKE_CMD) -C $(GRAFANA) stop
# grafana-restart:
# 	$(MAKE_CMD) -C $(GRAFANA) restart
# grafana-rm:
# 	$(MAKE_CMD) -C $(GRAFANA) rm
# grafana-down:
# 	$(MAKE_CMD) -C $(GRAFANA) down
# grafana-ps:
# 	$(MAKE_CMD) -C $(GRAFANA) ps
# grafana-exec:
# 	$(MAKE_CMD) -C $(GRAFANA) exec