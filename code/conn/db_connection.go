package conn

import (
	"app/code/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	fmtDBString = "postgres://%s:%s@%s:%d/%s"
)

func ConnectDb() *gorm.DB {
	c := config.NewDB()
	var logLevel gormlogger.LogLevel
	if c.Debug {
		logLevel = gormlogger.Info
	} else {
		logLevel = gormlogger.Error
	}
	dbString := fmt.Sprintf(fmtDBString, c.Username, c.Password, c.Host, c.Port, c.DBName)
	conn, err := gorm.Open(postgres.Open(dbString), &gorm.Config{
		Logger: gormlogger.Default.LogMode(logLevel),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		}})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
	return conn
}
