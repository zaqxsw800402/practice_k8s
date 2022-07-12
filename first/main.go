package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("start server")

	go saveToFile()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	err := app.Listen(":6666")
	if err != nil {
		log.Fatal(err)
		return
	}
}

func saveToFile() {
	// check tmp folder
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		os.Mkdir("tmp", 0755)
	}

	// save to file
	file, err := os.Create("tmp/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for i := 0; ; i++ {
		// write some text to file
		s := fmt.Sprintf("This is line %d\n", i)
		file.WriteString(s)

		// save changes
		err = file.Sync()
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(3 * time.Second)
	}
}
