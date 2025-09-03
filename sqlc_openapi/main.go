package main

import (
	"context"
	"fmt"
	"log"

	"sqlc_openapi/internal/db" // Update with your module name

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// Connect to the database using sqlx
	connStr := "postgres://yseekiaw:Jb220614@2025@localhost:5432/pgdatabase?sslmode=disable"
	dbx, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbx.Close()

	// Use sqlc's generated query functions
	queries := db.New(dbx)
	ctx := context.Background()

	// Create a new user
	user1, err := queries.CreateUser(ctx, db.CreateUserParams{
		Name:  "Alice",
		Email: "alice@example.com",
	})
	if err != nil {
		log.Fatalf("failed to create user: %v", err)
	}
	fmt.Printf("Created user: %+v\n", user1)

	// Get a user by ID
	userByID, err := queries.GetUser(ctx, user1.ID)
	if err != nil {
		log.Fatalf("failed to get user: %v", err)
	}
	fmt.Printf("Fetched user by ID: %+v\n", userByID)

	// List all users
	users, err := queries.ListUsers(ctx)
	if err != nil {
		log.Fatalf("failed to list users: %v", err)
	}
	fmt.Printf("All users: %+v\n", users)
}
