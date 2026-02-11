package conn

import (
	"app/code/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

const (
	fmtDBString = "postgres://%s:%s@%s:%d/%s"
)

func ConnectDb() *pgx.Conn {
	c := config.NewDB()
	dbString := fmt.Sprintf(fmtDBString, c.Username, c.Password, c.Host, c.Port, c.DBName)
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
