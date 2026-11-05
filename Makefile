# Everyday commands. `make` alone lists them.
APP := vx-that-for-me

.DEFAULT_GOAL := help
.PHONY: help test run deploy status logs restart security

help: ## List the commands
	@grep -E '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  \033[1m%-10s\033[0m %s\n", $$1, $$2}'

test: ## Run the tests
	mise exec -- go test ./...

run: ## Run locally with the variables in .env
	mise exec -- go run .

deploy: ## Build and install or update the service (asks for sudo)
	scripts/deploy.sh

status: ## Show the service status
	systemctl status $(APP) --no-pager

logs: ## Follow the service logs
	journalctl -u $(APP) -f

restart: ## Restart the service, e.g. after editing the env file
	sudo systemctl restart $(APP)

security: ## Score the unit's sandbox
	systemd-analyze security $(APP) --no-pager
