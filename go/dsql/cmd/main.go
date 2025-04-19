package main

import (
	"context"
	"fmt"

	_ "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/tusmasoma/go-microservice-k8s/go/pkg/log"
	"github.com/tusmasoma/samples/go/dsql"
)

type Owner struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Telephone string `json:"telephone"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Info("No .env file found", log.Ferror(err))
	}
	ctx := context.Background()
	// Establish connection
	conn, err := dsql.NewPostgresConn(ctx)
	if err != nil {
		return
	}

	// Create owner table
	_, err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS owner (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name VARCHAR(255),
			city VARCHAR(255),
			telephone VARCHAR(255)
		)
	`)
	if err != nil {
		return
	}

	// insert data
	query := `INSERT INTO owner (id, name, city, telephone) VALUES ($1, $2, $3, $4)`
	_, err = conn.Exec(ctx, query, uuid.New(), "John Doe", "Anytown", "555-555-0150")

	if err != nil {
		return
	}

	owners := []Owner{}
	// Define the SQL query to insert a new owner record.
	query = `SELECT id, name, city, telephone FROM owner where name='John Doe'`

	rows, err := conn.Query(ctx, query)
	defer rows.Close()

	owners, err = pgx.CollectRows(rows, pgx.RowToStructByName[Owner])
	fmt.Println(owners)
	if err != nil || owners[0].Name != "John Doe" || owners[0].City != "Anytown" {
		panic("Error retrieving data")
	}

	// Delete some data
	_, err = conn.Exec(ctx, `DELETE FROM owner where name='John Doe'`)
	if err != nil {
		return
	}

	defer conn.Close(ctx)

	return
}
