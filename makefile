rundb : 
	docker run --name postgresdb -p 4321:5432 -e POSTGRES_PASSWORD=secret -d postgres:latest

run : 
	go run main.go

.PHONY: makedb