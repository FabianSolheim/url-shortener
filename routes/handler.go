package routes

import (
	"net/url"
	"url-shortener/models"
	"url-shortener/repository"

	"github.com/gofiber/fiber/v2"
)

func LinkHandler(c *fiber.Ctx) error {
	link := &models.Link{}
	if err := c.BodyParser(link); err != nil {
		return err
	}

	if link.Alias == "" || link.Link == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Link alias or redirect link cannot be empty")
	}

	_, err := repository.GetLink(link.Alias)
	if err == nil {
		return fiber.NewError(fiber.StatusBadRequest, "Link alias already exists")
	}

	u, err := url.ParseRequestURI(link.Link)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid redirect link")
	}

	link.Link = u.String()

	newLink, err := repository.CreateLink(link)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Error while creating link")
	}

	return c.SendString("Link added successfully: " + newLink.Alias)
}

func MapHandler(c *fiber.Ctx) error {
	alias := c.Path()[1:]
	link, err := repository.GetLink(alias)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Could not find link, please check the alias")
	}

	return c.Redirect(link.Link)
}
