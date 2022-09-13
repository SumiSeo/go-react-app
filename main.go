package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:rootroot@/yt_go_auth"), &gorm.Config{})

	if err!= nil {
		panic("could not connect to the database")
	}

	app := fiber.New() // create a new Fiber instance
  
	// Create a new endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world!")
	})
  
	// Start server on port 3000
	app.Listen(":3000")
  }