package main

import (
	"app/code/config"
	"app/code/conn"
	"app/code/router"
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"
)

//  @title          APP API
//  @version        1.0
//  @description    This is a sample RESTful API with limited features of ticket master

//  @contact.name   Ahmed Hasnat Safder
//  @contact.email  safder.h@outlook.com

// @host       localhost:8081
// @basePath   /v1
func main() {
	var db *pgx.Conn = conn.ConnectDb()
	r := router.New(db)
	c := config.New()

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Server.Port),
		Handler:      r,
		ReadTimeout:  c.Server.TimeoutRead,
		WriteTimeout: c.Server.TimeoutWrite,
		IdleTimeout:  c.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server start failed")
	}
}
