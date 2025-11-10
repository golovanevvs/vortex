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

NAME = ${NGINX_CONTAINER}
PROJECT_NAME = ${NGINX_CONTAINER}
SERVICE_NAME = ${NGINX_CONTAINER}

up:
	@echo "Starting ${NAME}..."
	docker compose -p $(PROJECT_NAME) up -d
	@echo "${NAME} is up and running."

reload:
	@echo "Reloading ${NAME}..."
	docker compose -p $(PROJECT_NAME) exec $(SERVICE_NAME) nginx -s reload
	@echo "${NAME} reloaded."

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
# NGINX
# ========================

# NGINX = nginx

# .PHONY: nginx-up nginx-reload nginx-logs nginx-stop nginx-restart nginx-rm nginx-down nginx-ps nginx-exec

# nginx-up:
# 	$(MAKE_CMD) -C $(NGINX) up
# nginx-reload:
# 	$(MAKE_CMD) -C $(NGINX) reload
# nginx-logs:
# 	$(MAKE_CMD) -C $(NGINX) logs
# nginx-stop:
# 	$(MAKE_CMD) -C $(NGINX) stop
# nginx-restart:
# 	$(MAKE_CMD) -C $(NGINX) restart
# nginx-rm:
# 	$(MAKE_CMD) -C $(NGINX) rm
# nginx-down:
# 	$(MAKE_CMD) -C $(NGINX) down
# nginx-ps:
# 	$(MAKE_CMD) -C $(NGINX) ps
# nginx-exec:
# 	$(MAKE_CMD) -C $(NGINX) exec