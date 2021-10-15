package main

import (
	"log"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/JhonatanRSantos/shortener/database/memory"
	"github.com/JhonatanRSantos/shortener/docs"
	"github.com/JhonatanRSantos/shortener/services/shorter"
)

const (
	API_HOST        = "localhost:5000"
	API_TITLE       = "Shortner API"
	API_VERSION     = "1.0.0"
	API_BASE_PATH   = "/"
	API_DESCRIPTION = "This is a url shortener made in Golang"
)

var db memory.MemoryDB
var app *fiber.App

func init() {
	app = fiber.New(fiber.Config{
		AppName: "Shortener API",
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			log.Println(e.Error())
			return c.Status(fiber.StatusInternalServerError).SendString(e.Error())
		},
		// StrictRouting: true,
	})

	app.Use(cors.New())
	app.Use(recover.New())

	db = memory.NewMemoryDB()

	docs.SwaggerInfo.Host = API_HOST
	docs.SwaggerInfo.Title = API_TITLE
	docs.SwaggerInfo.Version = API_VERSION
	docs.SwaggerInfo.BasePath = API_BASE_PATH
	docs.SwaggerInfo.Description = API_DESCRIPTION
}

// @contact.name API Jhonatan R. Santos
// @contact.url https://github.com/JhonatanRSantos/shortener

// @license.name MIT
// @license.url https://github.com/JhonatanRSantos/shortener/blob/main/LICENSE
func main() {
	app.Get("/docs/*", swagger.Handler)

	app.Get("/:uri", shorter.NewRedirect(&db))
	app.Post("/", shorter.NewCreate(&db))

	app.Listen(":5000")
}
