package conn

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func ConnectDb() *sql.DB{
	var db *sql.DB
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("MYSQL_USER")
	cfg.Passwd = os.Getenv("MYSQL_PASSWORD")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("MYSQL_ADDRESS")
	cfg.DBName = os.Getenv("MYSQL_DATABASE")
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db
}
