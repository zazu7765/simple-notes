package main

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
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

func checkToken(user *jwt.Token) (uint, error) {
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))
	if time.Now().Unix() > int64(claims["expiration"].(float64)) {
		return userID, errors.New("token Expired")
	}
	return userID, nil
}

func getDefault(c *fiber.Ctx) error {
	return c.SendString("server is up!")
}
func signUpUser(c *fiber.Ctx) error {
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
		Notebook: []Notebook{
			{
				Title: "Untitled Notebook",
				Notes: []Note{
					{
						Title: "My First Note",
					},
				},
			},
		}}

	db.Create(&user)

	token, exp, err := generateJWT(user)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"token": token, "expiration": exp, "user": user})
}
func loginUser(c *fiber.Ctx) error {
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
}
func getUser(c *fiber.Ctx) error {
	var retrievedUser User
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		return err
	}
	if db.Where("id = ?", userID).Find(&retrievedUser).Error != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"success": true,
		"path":    "private",
		"name":    retrievedUser.Name,
		"email":   retrievedUser.Email})
}
func getNote(c *fiber.Ctx) error {
	var notebook Notebook
	var note Note
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "token expired or invalid",
		})
	}
	err = db.Model(&Notebook{}).Where("user_id = ?", userID).Order("created_at").First(&notebook).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "user has no notebook",
		})
	}
	err = db.Model(&Note{}).Where("notebook_id = ?", notebook.ID).Order("created_at").First(&note).Error
	if err != nil {
		return c.JSON(fiber.Map{
			"success": false,
			"error":   "user has no note",
		})
	}
	fmt.Println(note.Title)
	return c.JSON(fiber.Map{
		"success": true,
		"note":    note,
	})

}
