package models

import "gorm.io/gorm"

// import "mo"

type User struct {
	gorm.Model
	Email    string `gorm:"unique"`
	Password string
	Posts    []Post `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
