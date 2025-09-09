DB_URL=postgres://root:secret@localhost:5432/postgres?sslmode=disable

network:
	docker network create bank-network
	
postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=postgres -d postgres:15-alpine
	
mysql:
	docker run --name mysql18 -p 3306:3306 -e MYSQL_ROOT_PASSWORD=secret -d mysql:18

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

init-migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path  db/migration -database $(DB_URL) -verbose up

migrateup1:
	migrate -path  db/migration -database $(DB_URL) -verbose up 1

migrateup2:
	migrate -path  db/migration -database $(DB_URL) -verbose up 2
	
migratedown:
	migrate -path db/migration -database $(DB_URL) -verbose down

migratedown1:
	migrate -path db/migration -database $(DB_URL) -verbose down 1

migratedown2:
	migrate -path db/migration -database $(DB_URL) -verbose down 2

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

db_docs:
	dbdocs build docs/db.dbml
	
db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml
	
sqlc:
	sqlc generate



