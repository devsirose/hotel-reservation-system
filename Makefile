postgresdb:
	sudo docker run --name account-service-db -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=root -e POSTGRES_DB=account-service -p 5432:5432 -d postgres:15
createdb:
	sudo docker exec -it account-service-db createdb --username=root --owner=root account-service
dropdb:
	sudo docker exec -it account-service-db dropdb account-service
migrateup:
	migrate -path db/migration/ -database "postgresql://root:mysecretpassword@localhost:5432/account-service?sslmode=disable" -verbose up
migratedown:
	yes | migrate -path db/migration/ -database "postgresql://root:mysecretpassword@localhost:5432/account-service?sslmode=disable" -verbose down
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