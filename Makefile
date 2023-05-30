dev:
	docker-compose up -d --build

dev-down:
	docker-compose down

.PHONY: dev dev-down