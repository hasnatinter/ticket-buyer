package main

import (
	"app/code/config"
	"app/code/conn"
	lg "app/code/logger"
	"app/code/router"
	"app/code/server"
	"fmt"
	"log"
	"net/http"

	"gorm.io/gorm"
)

//  @title          APP API
//  @version        1.0
//  @description    This is a sample RESTful API with limited features of ticket master

//  @contact.name   Ahmed Hasnat Safder
//  @contact.email  safder.h@outlook.com

// @host       localhost:8081
// @basePath   /v1
func main() {
	c := config.New()

	l := lg.New(c.Server.Debug)
	var db *gorm.DB = conn.ConnectDb()
	server := server.New(l, db)
	r := router.New(server)

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
