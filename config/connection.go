package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

func DatabaseConnect() *pgx.Conn {
	//err := godotenv.Load(".env")
	
	//if err != nil {
	//		log.Fatalf("Some error occured. Err: %s", err)
	//}
	
	databaseUrl := os.Getenv("DATABASE_URL")
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	
	return conn
}

var Conn *pgx.Conn = DatabaseConnect()
