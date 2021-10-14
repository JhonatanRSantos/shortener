package shorter

import (
	"github.com/gofiber/fiber/v2"

	"github.com/JhonatanRSantos/shortener/database/interfaces"
	"github.com/JhonatanRSantos/shortener/models/link"
	"github.com/JhonatanRSantos/shortener/utils"
)

const randomStringSize = 8

var dbCreate interfaces.DBConnection

type CreateParams struct {
	URL string `json:"url" validate:"required,min=1"`
}

// Generate a shorter URL
// @Param c The current fiber context
func Create(c *fiber.Ctx) error {
	params := &CreateParams{}
	if err := utils.ParseAndValidateBodyParams(c.BodyParser, params); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	date, err := utils.GetCurrenteDate()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	link := link.NewLink(
		params.URL, utils.RandomString(randomStringSize), date, date,
	)
	link.Save(dbCreate)

	return c.SendString(c.BaseURL() + "/" + link.URI)
}

// Creates a new Handler
func NewCreate(dbConnection interfaces.DBConnection) func(c *fiber.Ctx) error {
	dbCreate = dbConnection
	return Create
}
