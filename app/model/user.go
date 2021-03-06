package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Id        string    `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	Password  string    `gorm:"column:password"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
