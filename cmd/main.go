package main

import (
	"superindo-test/config"
	"superindo-test/db"
	"superindo-test/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	conf := config.InitConfiguration()
	db.InitDatabase(conf)
	db := db.DB

	app := fiber.New()
	app.Use(logger.New())

	internal.InitializeModule(db, app, conf)

	app.Listen(":8080")
}
