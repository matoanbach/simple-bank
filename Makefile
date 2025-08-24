postgres:
	docker run -d --name postgres12 -p 49466:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:49466/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:49466/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:49466/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:49466/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

install-migrate:
	curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.3/migrate.linux-amd64.tar.gz | tar xvz

test:
	go test -v -cover ./...

.PHONY: posgres createdb dropdb migrateup migratedown sqlc install-migrate