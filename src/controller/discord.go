package controller

import "github.com/gofiber/fiber/v2"

func GetDiscordUser(c *fiber.Ctx) error {
	return c.SendString("Discord User Endpoint")
}