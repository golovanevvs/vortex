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

.PHONY: up logs stop restart rm down ps exec create-db list-db list-tables show-table

NAME = ${POSTGRES_CONTAINER}
PROJECT_NAME = ${POSTGRES_CONTAINER}
SERVICE_NAME = ${POSTGRES_CONTAINER}

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

create-db:
	@echo "Creating database: $(DB)"
	docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d postgres -c "CREATE DATABASE $(DB);"
	@echo "Database $(DB) created."

list-db:
	@echo "Listing databases..."
	@docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -l

list-tables:
	@echo "Listing tables in database \"$(DB)\"..."
	@docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(DB) -c "\dt"

show-table:
	@echo "Showing 10 rows of table \"$(TABLE)\"..."
	@docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d $(DB) \
	-c "SELECT * FROM $(TABLE) LIMIT 10"

drop-db:
	@if [ -z "$(DB)" ]; then \
		echo "ERROR: DB is required."; \
		exit 1; \
	fi
	@echo "Dropping database: $(DB)"
	docker exec -it $(POSTGRES_CONTAINER) psql -U $(POSTGRES_USER) -d postgres -c "DROP DATABASE IF EXISTS $(DB);"
	@echo "Database $(DB) dropped (if it existed)."


# ========================
# POSTGRES
# ========================

# POSTGRES = postgres

# .PHONY: postgres-up postgres-logs postgres-stop postgres-restart postgres-rm postgres-down postgres-ps postgres-exec postgres-create-db postgres-list-db postgres-list-tables postgres-show-table postgres-drop-db

# postgres-up:
# 	$(MAKE_CMD) -C $(POSTGRES) up
# postgres-logs:
# 	$(MAKE_CMD) -C $(POSTGRES) logs
# postgres-stop:
# 	$(MAKE_CMD) -C $(POSTGRES) stop
# postgres-restart:
# 	$(MAKE_CMD) -C $(POSTGRES) restart
# postgres-rm:
# 	$(MAKE_CMD) -C $(POSTGRES) rm
# postgres-down:
# 	$(MAKE_CMD) -C $(POSTGRES) down
# postgres-ps:
# 	$(MAKE_CMD) -C $(POSTGRES) ps
# postgres-exec:
# 	$(MAKE_CMD) -C $(POSTGRES) exec
# postgres-create-db:
# 	$(MAKE_CMD) -C $(POSTGRES) create-db DB="$(DB)"
# postgres-list-db:
# 	$(MAKE_CMD) -C $(POSTGRES) list-db"
# postgres-list-tables:
# 	$(MAKE_CMD) -C $(POSTGRES) list-tables DB="$(DB)"
# postgres-show-table:
# 	$(MAKE_CMD) -C $(POSTGRES) show-table DB="$(DB)" TABLE="$(TABLE)"
# postgres-drop-db:
# 	$(MAKE_CMD) -C $(POSTGRES) drop-db DB="$(DB)"
