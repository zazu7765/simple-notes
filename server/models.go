package main

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null" json:"-"`
	Notebook []Notebook
}
type Notebook struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	Title  string `gorm:"default:Untitled Notebook"`
	Notes  []Note
}
type Note struct {
	gorm.Model
	NotebookID uint   `gorm:"not null"`
	Title      string `gorm:"default:Untitled Note"`
	Content    string `gorm:"default:Lorem Ipsum Dolor Sit Amet"`
}

func migrateAll(db *gorm.DB) error {
	err := db.AutoMigrate(User{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(Notebook{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(Note{})
	if err != nil {
		panic(err)
	}
	return err
}
