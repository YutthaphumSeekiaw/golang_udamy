# Install sqlc
brew install sqlc
or
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Install golang-migrate CLI
brew install golang-migrate
or
go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate create -ext sql -dir ./database/migrations -seq create_users_table
migrate -path ./database/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" up

migrate -path ./database/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down 1 # Revert one migration
migrate -path ./database/migrations -database "postgres://user:password@localhost:5432/dbname?sslmode=disable" down # Revert all migrations


# For your Go project
go get github.com/jmoiron/sqlx
go get github.com/golang-migrate/migrate/v4/database/postgres
go get github.com/golang-migrate/migrate/v4/source/file

//คำสั่งนี้สร้าง file up and down
#migrate create -ext sql -dir db/migrations -seq create_users_table

# Ensure you have set the correct database connection string
export DATABASE_URL="postgres://yseekiaw:Jb220614@2025@localhost:5432/pgdatabase?sslmode=disable" && migrate -path db/migrations -database "$DATABASE_URL" up

# Generate code
sqlc generate

# Clean Lib not use
go mod tidy

# Run
go run .
