package model

import "time"

// Patient Model
////type Patient struct {
//	ID    uint   `gorm:"primaryKey:autoIncrement"`
//	Name  string `gorm:"size:24"`
//	Phone string `gorm:"size:24"`
//	Age   int    `gorm:"size:2"`
//	Place string `gorm:"size:24"`
////}

// Appointment Model
type Appointment struct {
	TokenID     uint      `gorm:"primaryKey:autoIncrement"`
	PatientName uint      `gorm:"size:24"`
	Age         int       `gorm:"size:2"`
	Place       string    `gorm:"size:24"`
	Date        time.Time `gorm:"type:date"`
	UserID      uint      `gorm:"not null"`
	DoctorID    uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
	Doctor      Doctor    `gorm:"foreignKey:DoctorID;references:DoctID"`
}

// Doctor Model
type Doctor struct {
	DoctID   uint   `gorm:"primaryKey:autoIncrement"`
	DoctName string `gorm:"size:24"`
	Dep      string `gorm:"size:24"`
}

// Medicine Model
type Medicine struct {
	MedID   uint   `gorm:"primaryKey:autoIncrement"`
	MedName string `gorm:"size:24"`
	Ilness  string `gorm:"size:24"`
	Price   int    `gorm:"size:4"`
}
