package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Static("/", "./views/index.html") //it's actually a type of app.Get function

	log.Fatal(app.Listen(":6021"))
}
