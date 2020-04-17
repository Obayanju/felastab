package main

import (
	"fmt"
	"log"

	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	jsoniter "github.com/json-iterator/go"
	"github.com/obayanju/felastab/search"
)

func main() {
	type CodeData struct {
		Code     string
		FilePath string
	}

	type Response struct {
		data []string
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Post("/parse", func(c *fiber.Ctx) {
		var codeData CodeData
		fmt.Println("parse request")
		if err := c.BodyParser(&codeData); err != nil {
			log.Fatal(err)
		}
		tokens := []string{}
		search.Start(codeData.Code, &tokens, codeData.FilePath)
		response := Response{}
		for _, token := range tokens {
			response.data = append(response.data, token)
		}
		toks, _ := jsoniter.Marshal(response.data)
		c.Write(string(toks))
	})

	app.Listen(3000)
}
