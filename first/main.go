package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	log.Println("start server")

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":6666")
}
