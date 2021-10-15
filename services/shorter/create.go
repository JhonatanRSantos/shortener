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
	URL string `json:"url" validate:"required,min=1" example:"www.google.com.br"`
}

type CreateResponse struct {
	BaseURL string `json:"base_url" example:"http://localhost:5000/"`
	URI     string `json:"uri" example:"u43oS8h6"`
}

// Generates a shorter URL
// @tags Shortener
// @Summary shortener
// @Description Generates a shorter URL
// @Param data body CreateParams true "The request body"
// @Success 200 {object} CreateResponse "A shorten url"
// @Failure 400,500
// @Router / [post]
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

	return c.JSON(CreateResponse{
		BaseURL: c.BaseURL() + "/",
		URI:     link.URI,
	})
}

// Creates a new Handler
func NewCreate(dbConnection interfaces.DBConnection) func(c *fiber.Ctx) error {
	dbCreate = dbConnection
	return Create
}
