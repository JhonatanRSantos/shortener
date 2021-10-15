package link

import (
	"errors"
	"strings"
	"time"

	"github.com/JhonatanRSantos/shortener/database/interfaces"
)

type Link struct {
	URL       string    `database:"url"`
	URI       string    `database:"uri,primary_key"`
	CreatedAt time.Time `database:"created_at"`
	UpdatedAt time.Time `database:"updated_at"`
}

// Save the current link
func (l *Link) Save(dbConnection interfaces.DBConnection) error {
	return dbConnection.SaveData("links", l)
}

// Find a link
func (l *Link) Find(dbConnection interfaces.DBConnection) (Link, error) {
	ouput, err := dbConnection.FindData("links", l)

	if err != nil {
		return Link{}, err
	}

	switch link := ouput.(type) {
	case *Link:
		return *link, nil
	default:
		return Link{}, errors.New("failed to parse the output")
	}
}

// Create a new link
func NewLink(url string, uri string, createdAt time.Time, updatedAt time.Time) Link {
	if !strings.Contains(url, "http://") || !strings.Contains(url, "https://") {
		url = "https://" + url
	}

	return Link{
		URL: url, URI: uri, CreatedAt: createdAt, UpdatedAt: updatedAt,
	}
}
