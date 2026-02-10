package main

import (
	"app/code/conn"
	"app/code/router"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {
	var db *pgx.Conn = conn.ConnectDb()
	r := router.New(db)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		IdleTimeout:  5 * time.Second,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server start failed")
	}
}
