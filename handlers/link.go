package handlers

import (
	"fmt"
	"url-shortener/models"
	"url-shortener/repository"
	"url-shortener/utils"

	"github.com/gofiber/fiber/v2"
)

type LinkHandler struct {
	Repository        *repository.LinkRepository
}

func NewLinkHandler(repository *repository.LinkRepository) *LinkHandler {
	return &LinkHandler{Repository: repository}
}


func (h *LinkHandler) CreateLink(c *fiber.Ctx) error {
	link := &models.Link{}
	if err := c.BodyParser(link); err != nil {
		return err
	}

	a, err := utils.ParseAlias(link.Alias)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid alias")
	}

	l, err := utils.ParseAndValidateUrl(link.Link)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid redirect link")
	}

	link.Link = l.String()
	link.Alias = a

	newLink, err := h.Repository.CreateLink(link)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Alias is already taken")
	}

	return c.SendString("Link added successfully: " + newLink)
}

func (h *LinkHandler) GetLink(c *fiber.Ctx) error {
	alias := c.Path()[1:]
	link, err := h.Repository.GetOneLink(alias)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "Could not find link, please check the alias")
	}

	return c.Redirect(link.Link)
}
