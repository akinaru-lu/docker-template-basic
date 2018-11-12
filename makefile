all: start

start:
	docker-compose -f gateway/docker-compose.yml up -d
	docker-compose -f app/docker-compose.yml up -d

stop:
	docker-compose -f app/docker-compose.yml down
	docker-compose -f gateway/docker-compose.yml down

restart:
	$(MAKE) start
	$(MAKE) stop
