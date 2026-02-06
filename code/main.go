package main

import (
	"database/sql"
	"net/http"

	"app/code/conn"
	"app/code/router"
)

func main() {
	var db *sql.DB = conn.ConnectDb()
	r := router.New(db)

	http.ListenAndServe("0.0.0.0:8080", r)
}
