package model

type TempApoiment struct {
	TokenID   uint   `json:"TokenID" gorm:"primaryKey"`
	Problem   string `json:"problem"`
	PatientID uint   `json:"patientID" gorm:"not null"`
}
