all: run

down:
	cd docker && docker-compose down

run:
	cd docker && docker-compose up
