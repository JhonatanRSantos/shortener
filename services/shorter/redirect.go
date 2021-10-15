package shorter

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/JhonatanRSantos/shortener/database/interfaces"
	"github.com/JhonatanRSantos/shortener/models/link"
)

var dbRedirect interfaces.DBConnection

// Find the given URI and Redirects the user to the requested website
// @tags Shortener
// @Summary Find and Redirect
// @Description Find the given URI and Redirects the user to the requested website
// @Param uri path string true "The URI for a given link" default(u43oS8h6)
// @Success 308
// @Failure 404,500
// @Router /{uri} [get]
func Redirect(c *fiber.Ctx) error {
	link := &link.Link{URI: c.Params("uri")}
	output, err := link.Find(dbRedirect)

	log.Println("Redirect", output)
	if err != nil {
		log.Println(err.Error())
		return c.SendStatus(404)
	}

	return c.Status(fiber.StatusPermanentRedirect).Redirect(output.URL)
}

// Creats a new Handler
func NewRedirect(dbConnection interfaces.DBConnection) func(c *fiber.Ctx) error {
	dbRedirect = dbConnection
	return Redirect
}
