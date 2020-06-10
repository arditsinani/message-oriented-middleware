.PHONY: all init clone build rebuild up stop restart status

all: up

status:
	@echo "*** Containers statuses ***"
	@docker-compose ps

build: stop
	@echo "*** Building containers... ***"
	docker-compose build

rebuild: stop
	@echo "*** Rebuilding containers... ***"
	docker-compose build --no-cache

up:
	@echo "*** Spinning up containers mom implementation... ***"
	docker-compose up -d
	@$(MAKE) --no-print-directory status

stop:
	@echo "*** Halting containers... ***"
	docker-compose stop
	@$(MAKE) --no-print-directory status

# Restart
restart:
	@echo "*** Restarting containers... ***"
	@$(MAKE) --no-print-directory stop
	@$(MAKE) --no-print-directory up

restart-ms-extractor:
	@echo "*** Restarting ms-extractor... ***"
	docker-compose restart ms-extractor

# Console
console-ms-extractor:
	@docker-compose exec ms-extractor sh

# Logs
logs-ms-extractor:
	@docker-compose logs -f -t --tail 30 ms-extractor

logs-ms-consumer:
	@docker-compose logs -f -t --tail 30 ms-consumer

logs-mongo:
	@docker-compose logs -f -t --tail 30 mongo_one mongo_two mongo_three

logs-kafka:
	@docker-compose logs -f -t --tail 30 kafka1

clean:
	@echo "*** Removing containers. All data will be lost!!!... ***"
	@docker-compose down --rmi all
	@rm -rf mongo/db/*
	@rm -rf mongo/dump/*
	@rm -rf zk-multiple-kafka-single/*
