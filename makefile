db : 
	docker run --name postgresdb -p 4321:5432 -e POSTGRES_PASSWORD=secret -d postgres:latest

rundb :
	docker start postgresdb

stopdb :
	docker stop postgresdb

run : 
	go run main.go

.PHONY: db, rundb, stopdb, run 