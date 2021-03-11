package api

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/template/html"
	"github.com/jin-wk/fiber-mysql/database"
	"github.com/jin-wk/fiber-mysql/routes"
)

// Init : app
func Init() {
	defer database.DB.Close()
	app := fiber.New(fiber.Config{
		// Prefork: true,
		Views: html.New("./views", ".html"),
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] [${path}] [${method}] ${body} ${resBody} > ${status}\n",
		TimeFormat: "2006-01-02 15:04:05.000",
		TimeZone:   "Asia/Seoul",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "FC-SIA",
		})
	})
	app.Post("/profile", routes.GetProfile)
	log.Fatal(app.Listen(":3000"))
}
