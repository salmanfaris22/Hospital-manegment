package model

import "time"

type Patient struct {
	ID           uint          `gorm:"primaryKey"`
	Name         string        `gorm:"size:24"`
	Phone        string        `gorm:"size:24"`
	Age          int           `gorm:"size:2"`
	Place        string        `gorm:"size:24"`
	Appointments []Appointment `gorm:"foreignKey:PatientID"`
}

type Appointment struct {
	TokenID   uint      `gorm:"primaryKey"`
	Date      time.Time `gorm:"type:date"`
	PatientID uint      `gorm:"not null"`
	DoctorID  uint      `gorm:"not null"`
	Patient   Patient   `gorm:"foreignKey:PatientID"`
	Doctor    Doctor    `gorm:"foreignKey:DoctorID"`
}

type Doctor struct {
	DoctID       uint          `gorm:"primaryKey"`
	DoctName     string        `gorm:"size:24"`
	Dep          string        `gorm:"size:24"`
	Appointments []Appointment `gorm:"foreignKey:DoctorID"`
	Medicines    []Medicine    `gorm:"foreignKey:MedID"`
}

type Medicine struct {
	MedID   uint   `gorm:"primaryKey"`
	MedName string `gorm:"size:24"`
}
