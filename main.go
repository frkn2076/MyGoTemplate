package main

import (
	"app/MyGoTemplate/db"
	"app/MyGoTemplate/logger"


	"github.com/gofiber/fiber"
	// "github.com/gofiber/fiber/middleware"
	"log"
)

func main() {

	// //Init logger
	// logger.Inıt()

	logger.InfoLog("asdsd")

	// Connect to database
	if err := db.Connect(); err != nil {
		log.Fatal(err)
	  }

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// Middleware
	// app.Use(middleware.Logger())

	// router.SetupRoutes(app)

	// listen on port 3000
	app.Listen(3000)

}

//   func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) {
// 		c.Send("Hello, World 👋!")
// 	})

// 	app.Post("/", func(c *fiber.Ctx) {
// 		// Get raw body from POST request
// 		c.Send(c.Body()) // user=john
// 	})

// 	app.Listen(3000)
// }
