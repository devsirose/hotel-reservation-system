postgresdb:
	sudo docker run --name postgres-local -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=root -e POSTGRES_DB=simple-bank -p 5430:5432 -d postgres
createdb:
	sudo docker exec -it postgres-local createdb --username=root --owner=root simple_bank
dropdb:
	sudo docker exec -it postgres-local dropdb simple_bank
migrateup:
	migrate -path db/migration/ -database "postgresql://root:mysecretpassword@localhost:5430/simple_bank?sslmode=disable" -verbose up
migratedown:
	yes | migrate -path db/migration/ -database "postgresql://root:mysecretpassword@localhost:5430/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
.PHONY: postgresdb createdb dropdb migrateup migratedown sqlc test