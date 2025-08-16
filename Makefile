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
server:
	go run main.go
mock:
	mockgen --package mockdb -destination db/mock/store.go github.com/devsirose/simplebank/db/sqlc Store
proto:
	rm -f pb/* | \
	protoc --proto_path=proto --go_out=./pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
        proto/*.proto
evans:
	 evans --host localhost --port 9090 -r repl
.PHONY: postgresdb createdb dropdb migrateup migratedown sqlc test server mock proto evans