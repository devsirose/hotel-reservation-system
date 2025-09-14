postgresdb:
	docker run --name hotel_reservation_db -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_USER=root -e POSTGRES_DB=hotel_reservation -p 5432:5432 -d postgis/postgis:15-3.3-alpine
createdb:
	docker exec -it hotel_reservation_db createdb --username=root --owner=root hotel_reservation
dropdb:
	docker exec -it hotel_reservation_db dropdb hotel_reservation
migrateup:
	migrate -path db/migration/ -database "postgresql://root:mysecretpassword@localhost:5432/hotel_reservation?sslmode=disable" -verbose up
migratedown:
	yes | migrate -path db/migration/ -database "postgresql://root:mysecretpassword@localhost:5432/hotel_reservation?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen --package mockdb -destination db/mock/store.go github.com/devsirose/hotel-reservation/db/sqlc Store
proto:
	rm -f pb/* | \
	protoc --proto_path=proto --go_out=./pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
        --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
        proto/*.proto
evans:
	 evans --host localhost --port 9090 -r repl
docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f postgres

run: docker-up
	@echo "Waiting for database to be ready..."
	@sleep 3
	go run main.go

.PHONY: postgresdb createdb dropdb migrateup migratedown sqlc test server mock proto evans docker-up docker-down docker-logs run