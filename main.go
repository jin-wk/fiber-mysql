package main

import (
	"log"

	"github.com/jin-wk/fiber-mysql/database"
	"github.com/jin-wk/fiber-mysql/models"
	"github.com/jin-wk/fiber-mysql/routes"
)

// @title Fiber-MySQL
// @version 1.0
// @description Basic Application by Fiber with MySQL
// @contact.name jin-wk
// @contact.url https://github.com/jin-wk
// @contact.email note@kakao.com
// @host localhost:3000
// @BasePath /
func main() {
	if err := database.Connect(); err != nil {
		log.Panic("Can't connect database: ", err.Error())
	}
	database.DB.AutoMigrate(&models.User{})
	app := routes.New()
	log.Fatal(app.Listen(":3000"))
}
