package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

)

func main() {
	fmt.Println("abc")
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Chao em Minh Thu")
	})

	app.Get("/:param", func(ctx *fiber.Ctx) error {
		return ctx.SendString("param: " + ctx.Params("param"))
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("POST aha")
	})

	app.Static("/static", "./public")
	// http://localhost:3000/static/foo.txt
	app.Static("/files", "./files")

	log.Fatal(app.Listen(":3000"))
	// http://localhost:3000/

	app.Get("/love/:param", func (c *fiber.Ctx) error {
		return c.SendString("I love : " + c.Params("param"))
	})

	// http://localhost:3000/love/minh%20thu
	// http://localhost:3000/love/phuong%20ly
	
	
}

// C:\Users\admin>netstat -ano | findstr ":3000"
//  TCP    0.0.0.0:3000           0.0.0.0:0              LISTENING       18836
// taskkill /PID 18836 /F

