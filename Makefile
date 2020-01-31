## build-client: Build the client.
build-client:
	yarn --cwd ./client build

## start: Build the client and start development mode.
start: build-client
	cd server && $(MAKE) start

## stop: Stop development mode.
stop:
	cd server && $(MAKE) stop

## restart: Build the client and restart development mode.
restart: build-client
	cd server && $(MAKE) restart

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECT_NAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
