postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createbd: 
	sudo docker exec -it postgres createdb --username=root --owner=root user
dropdb:
	sudo docker exec -it postgres dropdb  user

migrateup: 
	migrate -path schema/ -database "postgresql://root:secret@localhost:5432/user?sslmode=disable" -verbose up

migratedown:
	migrate -path schema/ -database "postgresql://root:secret@localhost:5432/user?sslmode=disable" -verbose down
.PHONY: postgres createbd dropdb migrateup migratedown

