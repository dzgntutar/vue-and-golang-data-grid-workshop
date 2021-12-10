package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		ctx.SendString("Hello from fiber")
		return nil
	})

	app.Listen(":3000")
}