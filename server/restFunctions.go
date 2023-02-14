package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

// REQUEST/RESPONSE STRUCTS
type signupRequest struct {
	Name     string `json:"name" ,validate:"required"`
	Email    string `json:"email" ,validate:"required"`
	Password string `json:"password" ,validate:"required"`
}
type loginRequest struct {
	Email    string `json:"email" ,validate:"required"`
	Password string `json:"password" ,validate:"required"`
}
type loginResponse struct {
	WebToken   string
	Expiration int64
	Username   string
}
type userResponse struct {
	Name  string
	Email string
}

// AUTH FUNCTIONS
func checkToken(user *jwt.Token) (uint, error) {
	claims := user.Claims.(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))
	if time.Now().Unix() > int64(claims["expiration"].(float64)) {
		return userID, fiber.NewError(fiber.StatusUnauthorized)
	}
	return userID, nil
}
func checkPassword(storedPassword string, checkPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(checkPassword))
	return err == nil
}
func validateLogin(request *loginRequest) []*validator.FieldError {
	var fieldErrors []*validator.FieldError
	if err := validate.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, &err)
		}
	}
	return fieldErrors
}
func validateSignUp(request *signupRequest) []*validator.FieldError {
	var fieldErrors []*validator.FieldError
	if err := validate.Struct(request); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldErrors = append(fieldErrors, &err)
		}
	}
	return fieldErrors
}

// GENERAL FUNCTIONS
func getDefault(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status":  "success",
		"Message": "Server is up",
		"Data":    nil,
	})
}

// USER FUNCTIONS
func signUpUser(c *fiber.Ctx) error {
	req := new(signupRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JSON Parsing Failed",
			"Data":    err,
		})
	}
	if err := validateSignUp(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JSON Validation Failed",
			"Data":    err,
		})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Password hashing failed",
			"Data":    err,
		})
	}
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hash),
		Notebook: []Notebook{
			{
				Title: "Personal Notebook",
				Notes: []Note{
					{
						Title: "My First Note",
					},
					{
						Title: "My Journal",
					},
					{
						Title: "My Grocery List",
					},
					{
						Title: "My To-Do List",
					},
				},
			},
			{
				Title: "Business Notebook",
				Notes: []Note{
					{
						Title: "Meeting Notes",
					},
					{
						Title: "To-Do List",
					},
				},
			},
		}}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Database Error",
			"Data":    "Could not create User",
		})
	}

	token, exp, err := generateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT creation failed",
			"Data":    err,
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status":  "success",
		"Message": "User signup successful",
		"Data": loginResponse{
			token,
			exp,
			user.Name,
		},
	})
}
func loginUser(c *fiber.Ctx) error {
	req := new(loginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Bad Credentials",
			"Data":    err,
		})
	}
	if err := validateLogin(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Bad Credentials",
			"Data":    err,
		})
	}
	var user User
	if err := db.Where("email = ?", req.Email).Order("id").First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Bad Credentials",
			"Data":    "Invalid Email",
		})
	}
	if !checkPassword(user.Password, req.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Bad Credentials",
			"Data":    "Invalid Password",
		})
	}
	token, exp, err := generateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT creation failed",
			"Data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "Login Successful",
		"Data": loginResponse{
			token,
			exp,
			user.Name},
	})
}
func getUser(c *fiber.Ctx) error {
	var retrievedUser User
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	if db.Where("id = ?", userID).First(&retrievedUser).Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Status":  "error",
				"Message": "Database Error",
				"Data":    "User not Found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Database Error",
			"Data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "User retrieved",
		"Data": userResponse{
			retrievedUser.Name,
			retrievedUser.Email,
		},
	})
}
func updateUser(c *fiber.Ctx) error {
	var retrievedUser User
	req := new(signupRequest)
	if err := validateSignUp(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JSON Validation Failed",
			"Data":    err,
		})
	}
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JSON Parsing Failed",
			"Data":    err,
		})
	}
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Password hashing failed",
			"Data":    err,
		})
	}
	if db.Where("id = ?", userID).First(&retrievedUser).Error != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"Status":  "error",
				"Message": "Database Error",
				"Data":    "User not Found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Database Error",
			"Data":    err,
		})
	}
	retrievedUser.Name = req.Name
	retrievedUser.Email = req.Email
	retrievedUser.Password = string(hash)
	if err := db.Save(&retrievedUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Unable to update user",
			"Data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "User updated",
		"Data":    nil,
	})
}

// NOTES FUNCTIONS
func getNote(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	note := Note{Model: gorm.Model{ID: uint(id)}}
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	if err = db.Table("notes").Select("notes.title, notes.content, notebook_id, notebooks.user_id").Joins("inner join notebooks on notebooks.id = notebook_id").Joins("inner join users on users.id = notebooks.user_id").Where("notebooks.user_id = ?", userID).First(&note).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"Status":  "error",
				"Message": "Database Error",
				"Data":    "Note not Found",
			})
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Note not Found",
			"Data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "Note found",
		"Data":    note,
	})

}
func getAllNotes(c *fiber.Ctx) error {
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	var notebooks []Notebook
	if err = db.Preload("Notes").Where("user_id = ?", userID).Find(&notebooks).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Notebook not Found",
			"Data":    err,
		})
	}
	var notes []Note
	for _, notebook := range notebooks {
		notes = append(notes, notebook.Notes...)
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "All User Notes Retrieved",
		"Data":    notes,
	})
}

// NOTEBOOK FUNCTIONS
func getNotebook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	notebook := Notebook{Model: gorm.Model{ID: uint(id)}}
	if err = db.Where("user_id = ?", userID).First(&notebook).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Notebook not Found",
			"Data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "All User Notes Retrieved",
		"Data":    notebook,
	})
}
func getAllNotebooks(c *fiber.Ctx) error {
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	var notebooks []Notebook
	if err = db.Where("user_id = ?", userID).Find(&notebooks).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Notebook Not Found")
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "All User Notes Retrieved",
		"Data":    notebooks,
	})

}
