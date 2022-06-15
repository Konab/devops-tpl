# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

run-agent:
	go mod tidy & go mod download &&\
	GIN_MODE=debug go run ./cmd/agent/main.go
.PHONY: run-agent

run-server:
	go mod tidy & go mod download &&\
	GIN_MODE=debug go run ./cmd/server/main.go
.PHONY: run-server
