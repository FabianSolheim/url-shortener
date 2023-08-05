package routes

import (
	"encoding/json"
	"net/url"
	"os"
	"url-shortener/models"
	"url-shortener/utils"

	"github.com/gofiber/fiber/v2"
)

func LinkHandler(c *fiber.Ctx) error {
	link := &models.Link{}
	if err := c.BodyParser(link); err != nil {
		return err
	}

	if(link.Alias == "" || link.Link == "") {
		return fiber.NewError(fiber.StatusBadRequest, "Link alias and redirect link cannot be empty")
	}

	if(utils.AlreadyExists(link.Alias)) {
		return fiber.NewError(fiber.StatusBadRequest, "Link alias already exists")
	}


	u, err := url.ParseRequestURI(link.Link)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid redirect link")
	}
	
	link.Link = u.String() //set the original link to the parsed url

	fileContent, err := models.GetLinks()
	if err != nil {
		return err
	}

	fileContent = append(fileContent, *link) //Add link from request to file content

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

	cleanPath := c.Path()[1:]

	for _, value := range links {
		if cleanPath == value.Alias {
			return c.Redirect(value.Link)
		}
	}
	return c.SendString("Link could not be found. Please check the link and try again.")
}
