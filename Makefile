postgresdb:
	sudo docker run --name hotel-reservation-system-db -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=root -e POSTGRES_DB=hotel-reservation-system -p 5432:5432 -d postgis/postgis:16-3.4
createdb:
	sudo docker exec -it hotel-reservation-system-db createdb --username=root --owner=root hotel-reservation-system
dropdb:
	sudo docker exec -it hotel-reservation-system-db dropdb hotel-reservation-system
migrateup:
	migrate -path db/migration/ -database "postgresql://root:mysecretpassword@localhost:5432/hotel-reservation-system?sslmode=disable" -verbose up
migratedown:
	yes | migrate -path db/migration/ -database "postgresql://root:mysecretpassword@localhost:5432/hotel-reservation-system?sslmode=disable" -verbose down
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