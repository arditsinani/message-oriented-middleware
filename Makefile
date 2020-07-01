.PHONY: all init clone build rebuild up stop restart status

DC := docker-compose
DR := docker

all: up

status:
	@echo "*** Containers statuses ***"
	$(DC) ps

build: stop
	@echo "*** Building containers... ***"
	$(DC) build

rebuild: stop
	@echo "*** Rebuilding containers... ***"
	$(DC) build --no-cache

up:
	@echo "*** Spinning up containers mom implementation... ***"
	docker-compose up -d
	@$(MAKE) --no-print-directory status

stop:
	@echo "*** Halting containers... ***"
	$(DC) stop
	@$(MAKE) --no-print-directory status

down:
	@echo "*** Removing containers... ***"
	$(DC) down
	@$(MAKE) --no-print-directory status

# Restart
restart:
	@echo "*** Restarting containers... ***"
	@$(MAKE) --no-print-directory stop
	@$(MAKE) --no-print-directory up

restart-ms-extractor:
	@echo "*** Restarting ms-extractor... ***"
	$(DC) restart ms-extractor

# Console
console-ms-extractor:
	$(DC) exec ms-extractor sh

# Mongo shell
console-mongo:
	$(DR) exec -it mongo_one mongo

# Logs
logs-ms-extractor:
	$(DC) logs -f -t --tail 30 ms-extractor

logs-ms-consumer:
	$(DC) logs -f -t --tail 30 ms-consumer

logs-mongo:
	$(DC) logs -f -t --tail 30 mongo_one mongo_two mongo_three

logs-kafka:
	$(DC) logs -f -t --tail 30 kafka1

clean:
	@echo "*** Removing containers. All data will be lost!!!... ***"
	$(DC) down --rmi all
	@rm -rf mongo/db/*
	@rm -rf mongo/dump/*
	@rm -rf zk-multiple-kafka-single/*
