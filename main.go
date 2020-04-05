package main

import (
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/obayanju/felastab/search"
)

func main() {
	type CodeData struct {
		Code string
	}
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/", func(c *fiber.Ctx) {
		var codeData CodeData
		if err := c.BodyParser(&codeData); err != nil {
			log.Fatal(err)
		}
		search.Start(codeData.Code)
	})

	app.Listen(3001)
}
