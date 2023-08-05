package routes

import (
	"encoding/json"
	"net/url"
	"os"
	"url-shortener/models"

	"github.com/gofiber/fiber/v2"
)

func LinkHandler(c *fiber.Ctx) error {
	link := &models.Link{}
	if err := c.BodyParser(link); err != nil {
		return err
	}

	if(link.ShortLink == "" || link.OriginalLink == "") {
		return fiber.NewError(fiber.StatusBadRequest, "Link alias and redirect link cannot be empty")
	}

	u, err := url.ParseRequestURI(link.OriginalLink)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid redirect link")
	}
	
	link.OriginalLink = u.String()

	fileContent, err := models.GetLinks()
	if err != nil {
		return err
	}

	fileContent = append(fileContent, *link)

	marshalledFileContent, err := json.MarshalIndent(fileContent, "", "    ")
	if err != nil {
		return err
	}

	json.MarshalIndent(marshalledFileContent, "", "    ")

	tempFile, err := os.CreateTemp("", "temp_links.json")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(marshalledFileContent); err != nil {
		tempFile.Close()
		return err
	}

	if err := tempFile.Close(); err != nil {
		return err
	}

	if err := os.Rename(tempFile.Name(), "links.json"); err != nil {
		return err
	}

	return c.SendString("Link added successfully!")
}

func MapHandler(c *fiber.Ctx) error {
	links, err := models.GetLinks()
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Error while fetching links")
	}

	for _, value := range links {
		if c.Path() == value.ShortLink {
			return c.Redirect(value.OriginalLink)
		}
	}
	return c.SendString("Link could not be found. Please check the link and try again.")
}
