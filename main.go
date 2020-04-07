package main

import (
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	jsoniter "github.com/json-iterator/go"
	"github.com/obayanju/felastab/search"
)

func main() {
	type CodeData struct {
		Code string
	}

	type Response struct {
		data []string
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/", func(c *fiber.Ctx) {
		var codeData CodeData
		if err := c.BodyParser(&codeData); err != nil {
			log.Fatal(err)
		}
		tokens := []string{}
		err := search.Start(codeData.Code, &tokens)
		if err != nil {
			log.Fatalf("Lexing failed -> %v", err)
		}
		response := Response{}
		for _, token := range tokens {
			response.data = append(response.data, token)
		}
		toks, _ := jsoniter.Marshal(response.data)
		c.Write(toks)
	})

	app.Listen(3000)
}
