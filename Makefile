postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123567 -d postgres

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

server:
	go run main.go

proto-w:
	del /F /Q pb\*.go
	del /F /Q doc\swagger\*.swagger.json
	protoc --proto_path=proto --proto_path=include --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=justtalk \
	proto/*.proto


proto-l:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --proto_path=include --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger \
	proto/*.proto

evans:
	evans --host localhost --port 9101 -r repl

client:
	go run ./client/main.go

redis:
	docker run --name redis -p 6380:6379 -d redis:latest

rabbitmq:
	docker run -it --name rabbitmq \
  	-p 5672:5672 -p 15672:15672 \
	-v rabbitmq_data:/var/lib/rabbitmq \
	rabbitmq:4.0-management
simple-pub:
	go run ./mq_test/simple_mod/simple_publish/publish.go

simple-rec:
	go run ./mq_test/simple_mod/simple_recieve/recieve.go

work-pub:
	go run ./mq_test/work_mod/producer/producer.go

work-rec:
	go run ./mq_test/work_mod/worker/worker.go

grpcurl:
	grpcurl -plaintext localhost:9101 list pb.JustTalk


.PHONY: postgres createdb dropdb stop_rm_postgres migrateup migratedown migrateversion migrateforce1 \
	sqlc proto-l proto-w evans client redis rabbitmq\
	simple-pub simple-rec