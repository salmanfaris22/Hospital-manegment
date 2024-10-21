package model

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	FirstName string    `gorm:"size:50;not null" json:"first_name" validate:"required"`
	LastName  string    `gorm:"size:50;not null" json:"last_name" validate:"required"`
	Password  string    `gorm:"size:100;not null" json:"password" validate:"required,min=6"`
	Email     string    `gorm:"size:100;unique;not null" json:"email" validate:"required,email"`
	Phone     string    `gorm:"size:15" json:"phone" validate:"omitempty,len=10,numeric"` // Ensuring phone is 10 digits
	Token     string    `gorm:"size:255"`
	UserType  string    `gorm:"size:20;default:user"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
