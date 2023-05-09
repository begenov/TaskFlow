postgres:
	sudo docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createbd: 
	sudo docker exec -it postgres12 createdb --username=root --owner=root task
dropdb:
	sudo docker exec -it postgres12 dropdb  task

migrateup: 
	migrate -path schema/ -database "postgresql://root:secret@localhost:5432/task?sslmode=disable" -verbose up

migratedown:
	migrate -path schema/ -database "postgresql://root:secret@localhost:5432/task?sslmode=disable" -verbose down
.PHONY: postgres createbd dropdb migrateup migratedown

