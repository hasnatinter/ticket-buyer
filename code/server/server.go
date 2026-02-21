package server

import (
	"app/code/logger"

	"gorm.io/gorm"
)

type Server struct {
	logger *logger.Logger
	db     *gorm.DB
}

func New(logger *logger.Logger, db *gorm.DB) *Server {
	return &Server{logger: logger, db: db}
}

func (server *Server) Logger() *logger.Logger {
	return server.logger
}

func (server *Server) DB() *gorm.DB {
	return server.db
}
