# ========================
# SETUP
# ========================

ENV_FILE ?= .env
ifneq (,$(wildcard $(ENV_FILE)))
include $(ENV_FILE)
export ENV PROJECT_NAME
endif

ifeq ($(OS),Windows_NT)
	MAKE_CMD = mingw32-make
else
	MAKE_CMD = make
endif

# ========================
# MIGRATE
# ========================

MIGRATE = providers/migrate

migrate-create:
	$(MAKE_CMD) -C $(MIGRATE) create
migrate-up:
	$(MAKE_CMD) -C $(MIGRATE) up
migrate-down:
	$(MAKE_CMD) -C $(MIGRATE) down


# ========================
# POSTGRES
# ========================

POSTGRES = providers/postgres

# Default target
postgres-help:
	$(MAKE_CMD) -C $(POSTGRES) help
# Basic Cluster Operations
postgres-up:
	$(MAKE_CMD) -C $(POSTGRES) up
postgres-up-async:
	$(MAKE_CMD) -C $(POSTGRES) up-async
postgres-up-sync-all:
	$(MAKE_CMD) -C $(POSTGRES) up-sync-all
postgres-up-sync-one:
	$(MAKE_CMD) -C $(POSTGRES) up-sync-one
postgres-down:
	$(MAKE_CMD) -C $(POSTGRES) down
postgres-restart:
	$(MAKE_CMD) -C $(POSTGRES) restart
postgres-status:
	$(MAKE_CMD) -C $(POSTGRES) status
postgres-status-sync:
	$(MAKE_CMD) -C $(POSTGRES) status-sync
postgres-logs:
	$(MAKE_CMD) -C $(POSTGRES) logs
postgres-setup:
	$(MAKE_CMD) -C $(POSTGRES) setup
postgres-check:
	$(MAKE_CMD) -C $(POSTGRES) check
postgres-clean:
	$(MAKE_CMD) -C $(POSTGRES) clean
postgres-purge:
	$(MAKE_CMD) -C $(POSTGRES) purge
# Replication Mode Management
postgres-mode-async:
	$(MAKE_CMD) -C $(POSTGRES) mode-async
postgres-mode-sync-all:
	$(MAKE_CMD) -C $(POSTGRES) mode-sync-all
postgres-mode-sync-one:
	$(MAKE_CMD) -C $(POSTGRES) mode-sync-one
postgres-mode-status:
	$(MAKE_CMD) -C $(POSTGRES) mode-status
# Testing and Monitoring
postgres-test:
	$(MAKE_CMD) -C $(POSTGRES) test
postgres-backup:
	$(MAKE_CMD) -C $(POSTGRES) backup
postgres-monitor:
	$(MAKE_CMD) -C $(POSTGRES) monitor
postgres-log-cont:
	$(MAKE_CMD) -C $(POSTGRES) log-cont
postgres-list-db:
	$(MAKE_CMD) -C $(POSTGRES) list-db
postgres-list-tables:
	$(MAKE_CMD) -C $(POSTGRES) list-tables
postgres-show-table:
	$(MAKE_CMD) -C $(POSTGRES) show-table
# Emergency Operations
postgres-promote-slave1:
	$(MAKE_CMD) -C $(POSTGRES) promote-slave1
postgres-promote-slave2:
	$(MAKE_CMD) -C $(POSTGRES) promote-slave2
postgres-check-master:
	$(MAKE_CMD) -C $(POSTGRES) check-master
postgres-check-lag:
	$(MAKE_CMD) -C $(POSTGRES) check-lag
# Auto Failover
postgres-pgauto-up:
	$(MAKE_CMD) -C $(POSTGRES) pgauto-up
postgres-pgauto-down:
	$(MAKE_CMD) -C $(POSTGRES) pgauto-down
postgres-pgauto-status:
	$(MAKE_CMD) -C $(POSTGRES) pgauto-status
postgres-pgauto-status-sync:
	$(MAKE_CMD) -C $(POSTGRES) pgauto-status-sync
postgres-pgauto-monitor:
	$(MAKE_CMD) -C $(POSTGRES) pgauto-monitor
postgres-pgauto-purge:
	$(MAKE_CMD) -C $(POSTGRES) pgauto-purge
postgres-clean-all:
	$(MAKE_CMD) -C $(POSTGRES) clean-all

# ========================
# RABBITMQ
# ========================

RABBITMQ = providers/rabbitmq

.PHONY: rabbitmq-up rabbitmq-logs rabbitmq-stop rabbitmq-restart rabbitmq-rm rabbitmq-down rabbitmq-ps rabbitmq-exec

rabbitmq-up:
	$(MAKE_CMD) -C $(RABBITMQ) up
rabbitmq-logs:
	$(MAKE_CMD) -C $(RABBITMQ) logs
rabbitmq-stop:
	$(MAKE_CMD) -C $(RABBITMQ) stop
rabbitmq-restart:
	$(MAKE_CMD) -C $(RABBITMQ) restart
rabbitmq-rm:
	$(MAKE_CMD) -C $(RABBITMQ) rm
rabbitmq-down:
	$(MAKE_CMD) -C $(RABBITMQ) down
rabbitmq-ps:
	$(MAKE_CMD) -C $(RABBITMQ) ps
rabbitmq-exec:
	$(MAKE_CMD) -C $(RABBITMQ) exec

# ========================
# REDIS
# ========================

REDIS = providers/redis

.PHONY: redis-up redis-logs redis-stop redis-restart redis-rm redis-down redis-ps redis-exec

redis-up:
	$(MAKE_CMD) -C $(REDIS) up
redis-logs:
	$(MAKE_CMD) -C $(REDIS) logs
redis-stop:
	$(MAKE_CMD) -C $(REDIS) stop
redis-restart:
	$(MAKE_CMD) -C $(REDIS) restart
redis-rm:
	$(MAKE_CMD) -C $(REDIS) rm
redis-down:
	$(MAKE_CMD) -C $(REDIS) down
redis-ps:
	$(MAKE_CMD) -C $(REDIS) ps
redis-exec:
	$(MAKE_CMD) -C $(REDIS) exec

# ========================
# PROMTAIL
# ========================

PROMTAIL = providers/promtail

promtail-up:
	$(MAKE_CMD) -C $(PROMTAIL) up
promtail-logs:
	$(MAKE_CMD) -C $(PROMTAIL) logs
promtail-stop:
	$(MAKE_CMD) -C $(PROMTAIL) stop
promtail-restart:
	$(MAKE_CMD) -C $(PROMTAIL) restart
promtail-rm:
	$(MAKE_CMD) -C $(PROMTAIL) rm
promtail-down:
	$(MAKE_CMD) -C $(PROMTAIL) down
promtail-ps:
	$(MAKE_CMD) -C $(PROMTAIL) ps
promtail-exec:
	$(MAKE_CMD) -C $(PROMTAIL) exec

# ========================
# LOKI
# ========================

LOKI = providers/loki

loki-up:
	$(MAKE_CMD) -C $(LOKI) up
loki-logs:
	$(MAKE_CMD) -C $(LOKI) logs
loki-stop:
	$(MAKE_CMD) -C $(LOKI) stop
loki-restart:
	$(MAKE_CMD) -C $(LOKI) restart
loki-rm:
	$(MAKE_CMD) -C $(LOKI) rm
loki-down:
	$(MAKE_CMD) -C $(LOKI) down
loki-ps:
	$(MAKE_CMD) -C $(LOKI) ps
loki-exec:
	$(MAKE_CMD) -C $(LOKI) exec

# ========================
# GRAFANA
# ========================

GRAFANA = providers/grafana

grafana-up:
	$(MAKE_CMD) -C $(GRAFANA) up
grafana-logs:
	$(MAKE_CMD) -C $(GRAFANA) logs
grafana-stop:
	$(MAKE_CMD) -C $(GRAFANA) stop
grafana-restart:
	$(MAKE_CMD) -C $(GRAFANA) restart
grafana-rm:
	$(MAKE_CMD) -C $(GRAFANA) rm
grafana-down:
	$(MAKE_CMD) -C $(GRAFANA) down
grafana-ps:
	$(MAKE_CMD) -C $(GRAFANA) ps
grafana-exec:
	$(MAKE_CMD) -C $(GRAFANA) exec

# ========================
# CLOUDPUB
# ========================

CLOUDPUB = providers/cloudpub

.PHONY: cloudpub-up cloudpub-logs cloudpub-stop cloudpub-restart cloudpub-rm cloudpub-down cloudpub-ps cloudpub-exec

cloudpub-up:
	$(MAKE_CMD) -C $(CLOUDPUB) up
cloudpub-logs:
	$(MAKE_CMD) -C $(CLOUDPUB) logs
cloudpub-stop:
	$(MAKE_CMD) -C $(CLOUDPUB) stop
cloudpub-restart:
	$(MAKE_CMD) -C $(CLOUDPUB) restart
cloudpub-rm:
	$(MAKE_CMD) -C $(CLOUDPUB) rm
cloudpub-down:
	$(MAKE_CMD) -C $(CLOUDPUB) down
cloudpub-ps:
	$(MAKE_CMD) -C $(CLOUDPUB) ps
cloudpub-exec:
	$(MAKE_CMD) -C $(CLOUDPUB) exec

# ========================
# CLOUDPUB-WEBHOOK
# ========================

CLOUDPUB-WEBHOOK = providers/cloudpub-webhook

.PHONY: cloudpub-webhook-up cloudpub-webhook-logs cloudpub-webhook-stop cloudpub-webhook-restart cloudpub-webhook-rm cloudpub-webhook-down cloudpub-webhook-ps cloudpub-webhook-exec

cloudpub-webhook-up:
	$(MAKE_CMD) -C $(CLOUDPUB-WEBHOOK) up
cloudpub-webhook-logs:
	$(MAKE_CMD) -C $(CLOUDPUB-WEBHOOK) logs
cloudpub-webhook-stop:
	$(MAKE_CMD) -C $(CLOUDPUB-WEBHOOK) stop
cloudpub-webhook-restart:
	$(MAKE_CMD) -C $(CLOUDPUB-WEBHOOK) restart
cloudpub-webhook-rm:
	$(MAKE_CMD) -C $(CLOUDPUB-WEBHOOK) rm
cloudpub-webhook-down:
	$(MAKE_CMD) -C $(CLOUDPUB-WEBHOOK) down
cloudpub-webhook-ps:
	$(MAKE_CMD) -C $(CLOUDPUB-WEBHOOK) ps
cloudpub-webhook-exec:
	$(MAKE_CMD) -C $(CLOUDPUB-WEBHOOK) exec

# ========================
# TRAEFIK
# ========================

TRAEFIK = providers/traefik

.PHONY: traefik-up traefik-logs traefik-stop traefik-restart traefik-rm traefik-down traefik-ps traefik-exec

traefik-up:
	$(MAKE_CMD) -C $(TRAEFIK) up
traefik-logs:
	$(MAKE_CMD) -C $(TRAEFIK) logs
traefik-stop:
	$(MAKE_CMD) -C $(TRAEFIK) stop
traefik-restart:
	$(MAKE_CMD) -C $(TRAEFIK) restart
traefik-rm:
	$(MAKE_CMD) -C $(TRAEFIK) rm
traefik-down:
	$(MAKE_CMD) -C $(TRAEFIK) down
traefik-ps:
	$(MAKE_CMD) -C $(TRAEFIK) ps
traefik-exec:
	$(MAKE_CMD) -C $(TRAEFIK) exec

# ========================
# PROMETHEUS
# ========================

PROMETHEUS = providers/prometheus

.PHONY: prometheus-up prometheus-logs prometheus-stop prometheus-restart prometheus-rm prometheus-down prometheus-ps prometheus-exec

prometheus-up:
	$(MAKE_CMD) -C $(PROMETHEUS) up
prometheus-logs:
	$(MAKE_CMD) -C $(PROMETHEUS) logs
prometheus-stop:
	$(MAKE_CMD) -C $(PROMETHEUS) stop
prometheus-restart:
	$(MAKE_CMD) -C $(PROMETHEUS) restart
prometheus-rm:
	$(MAKE_CMD) -C $(PROMETHEUS) rm
prometheus-down:
	$(MAKE_CMD) -C $(PROMETHEUS) down
prometheus-ps:
	$(MAKE_CMD) -C $(PROMETHEUS) ps
prometheus-exec:
	$(MAKE_CMD) -C $(PROMETHEUS) exec