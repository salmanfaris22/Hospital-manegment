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
	PatientName string    `gorm:"size:24" json:"patient_name"`
	Age         int       `gorm:"size:2" json:"age"`
	Place       string    `gorm:"size:24" json:"place"`
	Date        time.Time `gorm:"type:date" json:"date"`
	Slot        string    `gorm:"not null" json:"slot"`
	UserID      uint      `gorm:"size:24" json:"user_id"`
	DoctorID    uint      `gorm:"not null" json:"doctor_id"`
	User        User      `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Doctor      Doctor    `gorm:"foreignKey:DoctorID;references:DoctID" json:"doctor"`
}

// Doctor Model
type Doctor struct {
	DoctID    uint   `gorm:"primaryKey:autoIncrement"`
	DoctName  string `gorm:"size:24"`
	Dep       string `gorm:"size:24"`
	TimeSlot1 string `gorm:"size:20"`
	TimeSlot2 string `gorm:"size:20"`
}

// Medicine Model
type Medicine struct {
	MedID   uint   `gorm:"primaryKey:autoIncrement"`
	MedName string `gorm:"size:24"`
	Ilness  string `gorm:"size:24"`
	Price   int    `gorm:"size:4"`
}

type Date struct {
	ID        uint      `gorm:"primaryKey:autoIncrement" json:"id"`
	DateTime  time.Time `gorm:"type:date" json:"date_time"`
	DoctorID  uint      `gorm:"not null" json:"doctor_id"`
	UserId    uint      `gorm:"size:24"`
	Available bool      `gorm:"default:true" json:"available"`
	Slot      string    `gorm:"type:varchar(20)" json:"slot"`
	Doctor    Doctor    `gorm:"foreignKey:DoctorID;references:DoctID" json:"doctor"`
}
