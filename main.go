package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {

	fmt.Println("Hello worlda")
	app := fiber.New()

	todos := []Todo{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg ": "Hello World"})

	})

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Json body is requierd"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		var x int = 5

		var p *int = &x

		fmt.Println(p)  //0X0001
		fmt.Println(*p) //5

		return c.Status(201).JSON(todo)

	})
	log.Fatal(app.Listen(":5000"))
}
