package models

import (
	"time"
)

const (
	PrescriptionStatusActive  = "active"
	PrescriptionStatusExpired = "expired"
	PrescriptionStatusVoided  = "voided"
	PrescriptionStatusFilled  = "filled"
)

type Prescription struct {
	ID           uint                     `gorm:"primary_key;auto_increment" json:"id"`
	PatientID    uint                     `gorm:"not null;index" json:"patient_id"`
	Patient      Patient                  `gorm:"foreignKey:PatientID" json:"patient"`
	DoctorID     uint                     `gorm:"not null;index" json:"doctor_id"`
	Doctor       Doctor                   `gorm:"foreignKey:DoctorID" json:"doctor"`
	Diagnosis    string                   `gorm:"type:text;not null" json:"diagnosis"`
	Notes        string                   `gorm:"type:text" json:"notes"`
	IssueDate    time.Time                `gorm:"not null" json:"issue_date"`
	ValidUntil   time.Time                `gorm:"not null" json:"valid_until"`
	Status       string                   `gorm:"type:varchar(20);default:'active'" json:"status"`
	PharmacyID   *uint                    `json:"pharmacy_id,omitempty"`
	FilledDate   *time.Time               `json:"filled_date,omitempty"`
	RefillCount  int                      `gorm:"default:0" json:"refill_count"`
	MaxRefills   int                      `gorm:"default:0" json:"max_refills"`
	IsControlled bool                     `gorm:"default:false" json:"is_controlled"`
	Medications  []PrescriptionMedication `gorm:"foreignKey:PrescriptionID" json:"medications"`
	CreatedAt    time.Time                `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time                `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time               `sql:"index" json:"-"`
}

type PrescriptionMedication struct {
	ID             uint      `gorm:"primary_key;auto_increment" json:"id"`
	PrescriptionID uint      `gorm:"not null" json:"prescription_id"`
	MedicineName   string    `gorm:"size:100;not null" json:"medicine_name"`
	Dosage         string    `gorm:"size:50;not null" json:"dosage"`
	Frequency      string    `gorm:"size:50;not null" json:"frequency"`
	Duration       string    `gorm:"size:50;not null" json:"duration"`
	Instructions   string    `gorm:"type:text" json:"instructions"`
	Quantity       int       `gorm:"not null" json:"quantity"`
	Substitution   bool      `gorm:"default:true" json:"substitution"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type PrescriptionCreateRequest struct {
	PatientID   uint                     `json:"patient_id" validate:"required"`
	DoctorID    uint                     `json:"doctor_id" validate:"required"`
	Diagnosis   string                   `json:"diagnosis" validate:"required"`
	Notes       string                   `json:"notes"`
	ValidUntil  time.Time                `json:"valid_until" validate:"required,gtfield=IssueDate"`
	MaxRefills  int                      `json:"max_refills" validate:"min=0"`
	Medications []PrescriptionMedication `json:"medications" validate:"required,min=1"`
}

type PrescriptionResponse struct {
	ID          uint                     `json:"id"`
	PatientName string                   `json:"patient_name"`
	DoctorName  string                   `json:"doctor_name"`
	IssueDate   time.Time                `json:"issue_date"`
	ValidUntil  time.Time                `json:"valid_until"`
	Status      string                   `json:"status"`
	Medications []PrescriptionMedication `json:"medications"`
}
