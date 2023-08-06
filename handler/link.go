package handler

import (
	"fmt"
	"net/url"
	"url-shortener/models"
	"url-shortener/repository"

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

	if link.Alias == "" || link.Link == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Link alias or redirect link cannot be empty")
	}

	u, err := url.ParseRequestURI(link.Link)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid redirect link")
	}

	link.Link = u.String()

	newLink, err := h.Repository.CreateLink(link)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Alias is already taken")
	}

	return c.SendString("Link added successfully: " + string(newLink))
}

func (h *LinkHandler) GetLink(c *fiber.Ctx) error {

	if c.Path() == "/favicon.ico" { //Better way to do this?
		return c.SendStatus(fiber.StatusNoContent)
	}

	alias := c.Path()[1:]
	link, err := h.Repository.GetOneLink(alias)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(fiber.StatusBadRequest, "Could not find link, please check the alias")
	}

	return c.Redirect(link.Link)
}
