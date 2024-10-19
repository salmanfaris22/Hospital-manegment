package model

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"size:50;not null"`
	LastName  string `gorm:"size:50;not null"`
	Password  string `gorm:"size:100;not null"`
	Email     string `gorm:"size:100;unique;not null"`
	Phone     string `gorm:"size:15"`
	Token     string `gorm:"size:255"`
	UserType  string `gorm:"size:20"`
	//RefreshToken string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	UserID    string    `gorm:"size:50"`
}
