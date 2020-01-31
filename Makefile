build-client:
	yarn --cwd ./client build

start: build-client
	cd server && $(MAKE) start

stop:
	cd server && $(MAKE) stop

restart: build-client
	cd server && $(MAKE) restart
	