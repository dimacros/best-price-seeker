package main

import (
	. "dimacros/best-price-seeker/app/desired-product"
	. "dimacros/best-price-seeker/app/desired-product/handlers"
	"os"

	"github.com/ghodss/yaml"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	app := fiber.New()
	api := app.Group("api")

	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("public/index.html")
	})

	app.Get("/openapi.json", func(c *fiber.Ctx) error {
		openApiYaml, err := os.ReadFile("openapi.yaml")

		if err != nil {
			return err
		}

		openApiJson, err := yaml.YAMLToJSON(openApiYaml)

		return c.Type("json").SendString(string(openApiJson))
	})

	api.Post("/desired-products", func(c *fiber.Ctx) error {
		createDesiredProductCommand := new(CreateDesiredProductCommand)

		if err := c.BodyParser(createDesiredProductCommand); err != nil {
			return err
		}

		return c.JSON(CreateDesiredProduct(createDesiredProductCommand, db))
	})

	api.Get("/desired-products", func(c *fiber.Ctx) error {
		return c.JSON(GetAllDesiredProducts(db))
	})

	api.Get("/desired-products", func(c *fiber.Ctx) error {
		return c.JSON(GetAllDesiredProducts(db))
	})

	api.Delete("/desired-products/:id", func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			return err
		}

		DeleteDesiredProduct(id, db)

		return c.SendStatus(204)
	})

	app.Listen(":3000")
}
