package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/graphql-go/graphql"
)

func main() {
	app := fiber.New()

	// sample field
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type:    graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) { return "Hello", nil },
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemConf := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemConf)
	fmt.Println(schema)

	if err != nil {
		log.Fatal("Failed to create a new schema with error, %v", err)
	}

	app.Post("/graphql", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		fmt.Println(c.Body())
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
