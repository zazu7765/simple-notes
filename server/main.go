package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type signupRequest struct {
	Name     string
	Email    string
	Password string
}
type loginRequest struct {
	Email    string
	Password string
}

func main() {
	config := Config{
		host:     "simple-notes-db-1",
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
		req := new(signupRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		if req.Name == "" || req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "bad signup credentials")
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		user := User{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hash),
			Notebook: Notebook{
				Title: "Untitled"}}

		db.Create(&user)

		token, exp, err := generateJWT(user)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"token": token, "expiration": exp, "user": user})
	})
	app.Post("/login", func(c *fiber.Ctx) error {
		req := new(loginRequest)
		if err := c.BodyParser(req); err != nil {
			return err
		}
		if req.Email == "" || req.Password == "" {
			return fiber.NewError(fiber.StatusBadRequest, "bad login credentials")
		}
		var user User
		if db.Where("email = ?", req.Email).Order("id").Find(&user).Error != nil {
			return fiber.NewError(fiber.StatusBadRequest, "user does not exist or bad login credentials")
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			return err
		}
		token, exp, err := generateJWT(user)
		if err != nil {
			return err
		}
		return c.JSON(fiber.Map{
			"token": token, "expiration": exp, "user": user})
	})
	private := app.Group("/private")
	private.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("bananas"),
	}))
	private.Get("/", func(c *fiber.Ctx) error {
		var retrievedUser User
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		if time.Now().Unix() > int64(claims["expiration"].(float64)) {
			return c.JSON(fiber.Map{
				"success": false,
				"path":    "private",
				"error":   "token"})
		}
		userID := claims["user_id"].(float64)
		if db.Where("id = ?", userID).Find(&retrievedUser).Error != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(fiber.Map{
			"success": true,
			"path":    "private",
			"name":    retrievedUser.Name,
			"email":   retrievedUser.Email})
	})
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
	exp := time.Now().Add(time.Minute * 5).Unix()
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
