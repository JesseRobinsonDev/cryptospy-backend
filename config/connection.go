package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

// Creates a connection to the PostgreSQL database on Heroku
func DatabaseConnect() *pgx.Conn {
	
	databaseUrl := os.Getenv("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	return conn
}

// Permanent connection for less strenuous tasks 
var Conn *pgx.Conn = DatabaseConnect()
