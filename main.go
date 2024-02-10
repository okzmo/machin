package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "https://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Set-Cookie",
		AllowCredentials: true,
	}))

	log.Fatal(app.Listen(":3000"))
}