package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	fmt.Println("Hello worlda")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg ": "Hello World"})

	})
	log.Fatal(app.Listen(":5000"))
}
