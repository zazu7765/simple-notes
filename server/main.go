package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
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
	app      = fiber.New()
	validate = validator.New()
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
	user.Put("/", updateUser)
	user.Delete("/", deleteUser)

	notebooks := app.Group("/notebooks")
	notebooks.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))
	notebooks.Get("/notebook/:id", getNotebook)
	notebooks.Get("/all", getAllNotebooks)
	//notebooks.Post("/")
	//notebooks.Put("")
	//notebooks.Delete("/")

	notes := app.Group("/notes")
	notes.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))
	notes.Get("/note/:id", getNote)
	notes.Get("/all", getAllNotes)

	//notes.Post("/")
	//notes.Put("/")
	notes.Delete("/:id", deleteNote)

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
