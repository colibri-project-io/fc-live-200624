STACK_NAME=dev-fc-city-service

start:
	docker-compose -p ${STACK_NAME} up -d --remove-orphans

stop:
	docker-compose -p ${STACK_NAME} stop

restart: stop start

clean:
	docker-compose -p ${STACK_NAME} down -v

logs:
	docker-compose -p ${STACK_NAME} logs -f

run:
	go run main.go
