MAKEFLAGS += --no-print-directory

NAME = Webhook

up:
	@echo "Starting ${NAME}..."
	start "Webhook Server" cmd /k run-webhook.bat
	@echo "${NAME} is up and running."

# ========================
# WEBHOOK
# ========================

# WEBHOOK = webhook

# .PHONY: webhook-up

# webhook-up:
# 	$(MAKE_CMD) -C $(WEBHOOK) up