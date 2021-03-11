package database

import (
	"database/sql"
	"fmt"
	"log"

	// import mysql engine
	_ "github.com/go-sql-driver/mysql"
	"github.com/jin-wk/fiber-mysql/config"
)

// DB : MySQL connection
var DB = Connect()

var (
	host     = config.Config("DB_HOST")
	port     = config.Config("DB_PORT")
	username = config.Config("DB_USERNAME")
	password = config.Config("DB_PASSWORD")
	database = config.Config("DB_DATABASE")
)

// Connect : connect to DB
func Connect() *sql.DB {
	var err error
	client, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database))
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	if err = client.Ping(); err != nil {
		log.Fatal(err.Error())
	}
	return client
}
