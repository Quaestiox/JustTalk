postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES USER=root -e POSTGRES PASSWORD=123567 -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root justtalk

dropdb:
	docker exec -it postgres dropdb justtalk -f

stop_rm_postgres:
	docker stop postgres
	docker rm postgres

migrateup:
	migrate -path db/migration -database "postgresql://root:123567@localhost:5433/justtalk?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:123567@localhost:5433/justtalk?sslmode=disable" -verbose down

migrateversion:
	migrate -path db/migration -database  "postgresql://root:123567@localhost:5433/justtalk?sslmode=disable" version

migrateforce1:
	migrate -path db/migration -database  "postgresql://root:123567@localhost:5433/justtalk?sslmode=disable" force 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb stop_rm_postgres migrateup migratedown migrateversion migrateforce1 sqlc