package routes

import (

	"github.com/daan0101/discordapi/src/controller"
)

var ApiRoutes = []Route{
	{Method: "GET", Path: "/discord/:id", Handler: controller.GetDiscordUser},
}