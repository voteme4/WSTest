package main

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("You shouldn't be here!")
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	client := make(map[*websocket.Conn]bool)

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		id := c.Params("id")
		fmt.Println("Client connected with id:", id)
		log.Println(c.Locals("allowed"))
		log.Println(c.Params("id"))
		log.Println(c.Query("v"))
		log.Println(c.Cookies("session"))

		client[c] = true

		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("Error reading message:", err)
				break
			}
			fmt.Println("Recieved message:", string(msg))

			for client := range client {
				if client != c {
					err = client.WriteMessage(mt, msg)
					if err != nil {
						log.Println("Error writing message:", err)
					}
				}
			}
		}
	}))

	log.Fatal(app.Listen(":5000"))
}
