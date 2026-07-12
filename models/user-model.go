package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm: "size:255;not null"`
	Email string `gorm: "size:255;not null;unique"`
	Age uint `gorm: "check:age >= 18;not null"`
}
