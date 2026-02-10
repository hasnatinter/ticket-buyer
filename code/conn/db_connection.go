package conn

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const (
	dbString = "postgres://hasnat:password@database:5432/ticket"
)

func ConnectDb() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), dbString)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := conn.Ping(context.Background())
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return conn
}
