package main

import (
	"app/code/config"
	"app/code/conn"
	"app/code/router"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func main() {
	var db *pgx.Conn = conn.ConnectDb()
	r := router.New(db)
	c := config.New()

	s := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server start failed")
	}
}
