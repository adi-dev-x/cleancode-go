package model

import "gorm.io/gorm"

type AdminRegister struct {
	gorm.Model
	Username string `json:"username" gorm:"not null;unique"`
	Name     string `json:"name" gorm:"not null;unique"`
	Email    string `json:"email" gorm:"not null;unique"`
	Phone    string `json:"phone" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
}
