package database

import (
	"fmt"
	"time"

	"github.com/jin-wk/fiber-mysql/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB       *gorm.DB
	host     = config.Config("DB_HOST")
	port     = config.Config("DB_PORT")
	username = config.Config("DB_USERNAME")
	password = config.Config("DB_PASSWORD")
	database = config.Config("DB_DATABASE")
)

// Connect : connect to DB
func Connect() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", username, password, host, port, database)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	con, err := DB.DB()
	if err != nil {
		return err
	}
	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(50)
	con.SetConnMaxLifetime(time.Hour)

	return nil
}
