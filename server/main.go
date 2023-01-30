package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := Config{
		host:     "localhost",
		port:     "5432",
		password: "postgres",
		user:     "postgres",
		sslMode:  "disable",
		name:     "simple_notes",
	}
	app := fiber.New()
	db, err := connect(config)
	if err != nil {
		panic("failed to connect")
	}
	if err := migrateAll(db); err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintln("Connected to", config.host, "as", config.user))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server is up!")
	})
	app.Post("/signup", func(c *fiber.Ctx) error {
		return nil
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		return nil
	})
	app.Get("/private", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "private"})
	})
	app.Get("/public", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "public"})
	})
	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}

	/*db.Create(&User{
		Name:     "daniel",
		Email:    "zazu7765@gmail.com",
		Password: "12341234",
		Notebook: Notebook{
			Title: "Daniel's Notebook",
			Notes: []Note{
				{
					Title:   "note 1",
					Content: "content in note",
				},
				{
					Title:   "note 2",
					Content: "content in note 2",
				},
			},
		},
	})
	var users []User
	db.Preload("Notebook.Notes").Preload("Notebook").Find(&users)
	//fmt.Println(user)
	for _, person := range users {
		fmt.Println(person.Name)
		fmt.Println(person.Notebook.Title)
		for _, note := range person.Notebook.Notes {
			fmt.Printf("title %s\n", note.Title)
			fmt.Printf("content %s\n", note.Content)
		}
		fmt.Println("---------")
	}
	*/
}
