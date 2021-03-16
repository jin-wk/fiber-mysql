package routes

import (
	fiberSwagger "github.com/arsmn/fiber-swagger"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware/cors"
	"github.com/gofiber/fiber/middleware/logger"
	"github.com/gofiber/template/html"
	_ "github.com/jin-wk/fiber-mysql/docs"
	"github.com/jin-wk/fiber-mysql/handlers"
)

// New : create an instance
func New() *fiber.App {
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
		Format:     "${yellow}[${time}] ${blue}[${path}] ${green}[${method}] ${white}${body} ${white}${resBody} > ${yellow}${status}\n",
		TimeFormat: "2006-01-02 15:04:05.000",
		TimeZone:   "Asia/Seoul",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Fiber-MySQL",
		})
	})
	app.Get("/swagger/*", fiberSwagger.Handler)
	api := app.Group(("/api"))
	api.Get("/user/:id", handlers.GetUser)

	return app
}
