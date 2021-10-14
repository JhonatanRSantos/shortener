package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/JhonatanRSantos/shortener/database/memory"
	"github.com/JhonatanRSantos/shortener/services/shorter"
)

var db memory.MemoryDB
var app *fiber.App

func main() {
	app = fiber.New()

	app.Use(cors.New())
	app.Use(recover.New())

	db = memory.NewMemoryDB()

	app.Get("/:uri", shorter.NewRedirect(&db))
	app.Post("/", shorter.NewCreate(&db))

	app.Listen(":5000")
}
