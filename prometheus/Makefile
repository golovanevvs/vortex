MAKEFLAGS += --no-print-directory

PROJECT_DIR := $(abspath ../..)

ENV_FILES := .env
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

NAME = ${PROMETHEUS_CONTAINER}
PROJECT_NAME = ${PROMETHEUS_CONTAINER}
SERVICE_NAME = ${PROMETHEUS_CONTAINER}

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
# PROMETHEUS
# ========================

# PROMETHEUS = providers/prometheus

# .PHONY: prometheus-up prometheus-logs prometheus-stop prometheus-restart prometheus-rm prometheus-down prometheus-ps prometheus-exec

# prometheus-up:
# 	$(MAKE_CMD) -C $(PROMETHEUS) up
# prometheus-logs:
# 	$(MAKE_CMD) -C $(PROMETHEUS) logs
# prometheus-stop:
# 	$(MAKE_CMD) -C $(PROMETHEUS) stop
# prometheus-restart:
# 	$(MAKE_CMD) -C $(PROMETHEUS) restart
# prometheus-rm:
# 	$(MAKE_CMD) -C $(PROMETHEUS) rm
# prometheus-down:
# 	$(MAKE_CMD) -C $(PROMETHEUS) down
# prometheus-ps:
# 	$(MAKE_CMD) -C $(PROMETHEUS) ps
# prometheus-exec:
# 	$(MAKE_CMD) -C $(PROMETHEUS) exec