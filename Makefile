postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root Go-bank

dropdb:
	docker exec -it postgres17 dropdb Go-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/Go-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/Go-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate


.PHONY: postgres createdb dropdb migrateup migratedown sqlc