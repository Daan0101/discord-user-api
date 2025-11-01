package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Route struct {
	Method  string
	Path    string
	Handler fiber.Handler
}

func RegisterRoutes(app *fiber.App, routes []Route) {
	for _, r := range routes {
		handlers := make([]fiber.Handler, 0)
		handlers = append(handlers, r.Handler)

		switch r.Method {
		case "GET":
			app.Get(r.Path, handlers...)
		case "POST":
			app.Post(r.Path, handlers...)
		case "PUT":
			app.Put(r.Path, handlers...)
		case "DELETE":
			app.Delete(r.Path, handlers...)
		default:
			panic("Unsupported method: " + r.Method)
		}
	}
}

func SetupAppRoutes(app *fiber.App) {
	RegisterRoutes(app, ApiRoutes)
}
