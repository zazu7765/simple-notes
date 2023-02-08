package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"time"
)

var (
	config = Config{
		host:     "simple-notes-db-1",
		port:     "5432",
		password: "postgres",
		user:     "postgres",
		sslMode:  "disable",
		name:     "simple_notes",
	}
	db = func() *gorm.DB {
		db, err := connect(config)
		if err != nil {
			panic(err)
		}
		if err := migrateAll(db); err != nil {
			panic(err)
		}
		return db
	}()
	app = fiber.New()
)

func main() {
	fmt.Println(fmt.Sprintln("Connected to", config.host, "as", config.user))

	app.Get("/", getDefault)

	sessions := app.Group("/session")
	sessions.Post("/create", signUpUser)
	sessions.Post("/login", loginUser)

	user := app.Group("/user")
	user.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))
	user.Get("/", getUser)
	//user.Put("/")
	//user.Delete("/")

	notebooks := app.Group("/notebook")
	notebooks.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))
	//notebooks.Get("/")
	//notebooks.Post("/")
	//notebooks.Put("")
	//notebooks.Delete("/")

	notes := app.Group("/note")
	notes.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))
	notes.Get("/", getNote)
	//notes.Post("/")
	//notes.Put("/")
	//notes.Delete("/")

	public := app.Group("/public")
	public.Get("/", func(c *fiber.Ctx) error {
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

func generateJWT(user User) (string, int64, error) {
	exp := time.Now().Add(time.Minute * 30).Unix()
	JWTToken := jwt.New(jwt.SigningMethodHS256)
	claim := JWTToken.Claims.(jwt.MapClaims)
	claim["user_id"] = user.ID
	claim["expiration"] = exp
	t, err := JWTToken.SignedString([]byte("bananas"))
	if err != nil {
		return "", 0, err
	}
	return t, exp, nil
}
