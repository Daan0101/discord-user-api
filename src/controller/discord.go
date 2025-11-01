package controller

import (
	"github.com/gofiber/fiber/v2"

	"github.com/daan0101/discordapi/src/types"
	"github.com/daan0101/discordapi/src/client"
)

func GetDiscordUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user types.User
	client := client.NewClient()
	err := client.Get("/users/"+id, &user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}