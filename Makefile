postgres:
	docker run --name postgres16 --network dummy-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.2-alpine

createdb: 
	docker exec -it postgres16 createdb --username=root --owner=root users

migrateup:
	goose --dir=database/migrations postgres "host=localhost password=secret user=root dbname=users sslmode=disable" up

migratedown:
	goose --dir=migrations postgres "host=localhost password=secret user=root dbname=users sslmode=disable" down

dropdb:
	docker exec -it postgres16 dropdb users

server: 
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown server

