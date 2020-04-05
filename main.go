package main

import (
	"log"

	"github.com/gofiber/fiber"
	"github.com/obayanju/felastab/search"
)

func main() {
	type Code struct {
		Text string
	}
	app := fiber.New()

	app.Post("/", func(c *fiber.Ctx) {
		var code Code
		if err := c.BodyParser(&code); err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("%v\n", code.Text)
		search.Start(code.Text)
	})

	app.Listen(3001)
}
