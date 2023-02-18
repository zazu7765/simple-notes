package main

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

// REQUEST STRUCTS
type signupRequest struct {
	Name     string `json:"name" ,validate:"required"`
	Email    string `json:"email" ,validate:"required"`
	Password string `json:"password" ,validate:"required"`
}
type loginRequest struct {
	Email    string `json:"email" ,validate:"required"`
	Password string `json:"password" ,validate:"required"`
}

type noteUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// RESPONSE STRUCTS
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

// AUTH FUNCTIONS
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
			"Data":    "User already exists!",
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

// USER FUNCTIONS
func getUser(c *fiber.Ctx) error {
	var retrievedUser User
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		c.ClearCookie("jwt")
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
		c.ClearCookie("jwt")
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
func deleteUser(c *fiber.Ctx) error {
	var retrievedUser User
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		c.ClearCookie("jwt")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	if db.Where("id = ?", userID).First(&retrievedUser).Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status":  "error",
			"Message": "User not found",
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
	for _, notebook := range notebooks {
		if err := db.Select(clause.Associations).Unscoped().Delete(&notebook).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"Status":  "error",
				"Message": "Failed to delete notebooks",
				"Data":    err,
			})
		}
	}
	if err := db.Select(clause.Associations).Unscoped().Delete(&User{Model: gorm.Model{ID: userID}}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Failed to delete user",
			"Data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "User deleted",
		"Data":    userID,
	})
}

// NOTES FUNCTIONS
func getNote(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	note := Note{Model: gorm.Model{ID: uint(id)}}
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		c.ClearCookie("jwt")
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
		c.ClearCookie("jwt")
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
func deleteNote(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	_, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		c.ClearCookie("jwt")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	if err := db.Unscoped().Delete(&Note{Model: gorm.Model{ID: uint(id)}}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Note could not be deleted",
			"Data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "User Deleted",
		"Data":    nil,
	})
}
func updateNote(c *fiber.Ctx) error {
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	id, _ := strconv.Atoi(c.Params("id"))
	if err != nil {
		c.ClearCookie("jwt")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JWT Expired or Invalid",
			"Data":    err,
		})
	}
	req := new(noteUpdateRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status":  "error",
			"Message": "JSON Parsing Failed",
			"Data":    err,
		})
	}
	note := Note{Model: gorm.Model{ID: uint(id)}}
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
	if len(req.Title) > 0 {
		note.Title = req.Title
	}
	if len(req.Content) > 0 {
		note.Content = req.Content
	}
	if err = db.Save(&note).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status":  "error",
			"Message": "Unable to update note",
			"Data":    err,
		})
	}
	return c.JSON(fiber.Map{
		"Status":  "success",
		"Message": "Note updated",
		"Data":    nil,
	})
}
func createUser(c *fiber.Ctx) error {
	return nil
}

// NOTEBOOK FUNCTIONS
func getNotebook(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	userID, err := checkToken(c.Locals("user").(*jwt.Token))
	if err != nil {
		c.ClearCookie("jwt")
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
		c.ClearCookie("jwt")
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
func deleteNotebook(c *fiber.Ctx) error {
	return nil
}
func updateNotebook(c *fiber.Ctx) error {
	return nil
}
func createNotebook(c *fiber.Ctx) error {
	return nil
}
