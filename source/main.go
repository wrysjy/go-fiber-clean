package main

import (
	"gofiber/controller/web"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	preFork := os.Getenv("PREFORK") == "true"
	app := fiber.New(fiber.Config{
		Prefork: preFork,
	})
	app.Use(logger.New())
	webRouter(app)

	log.Fatal(app.Listen(":3000"))
}

func webRouter(app *fiber.App) {
	app.Get("/", web.Home)
}
