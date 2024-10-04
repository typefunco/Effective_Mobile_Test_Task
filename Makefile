# i'LL WRITE IT

run:
	- go run cmd/app/main.go

start_service:
	@echo "Starting service"
	- docker compose up -d --build