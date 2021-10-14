package shorter

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/JhonatanRSantos/shortener/database/interfaces"
	"github.com/JhonatanRSantos/shortener/models/link"
)

var dbRedirect interfaces.DBConnection

// Find Redirects the user to the requested website
// @Param c The current fiber context
func Redirect(c *fiber.Ctx) error {
	link := &link.Link{URI: c.Params("uri")}
	output, err := link.Find(dbRedirect)

	log.Println("Redirect", output)
	if err != nil {
		log.Println(err.Error())
		return c.SendStatus(404)
	}

	return c.Redirect(output.URL)
}

// Creats a new Handler
func NewRedirect(dbConnection interfaces.DBConnection) func(c *fiber.Ctx) error {
	dbRedirect = dbConnection
	return Redirect
}
