package models

import (
	"time"
)

const (
	StatusActive   = "active"
	StatusInactive = "inactive"
	StatusDeceased = "deceased"

	GenderMale   = "male"
	GenderFemale = "female"
	GenderOther  = "other"
)

type Patient struct {
	ID               uint       `gorm:"primary_key;auto_increment" json:"id"`
	UserID           uint       `gorm:"not null" json:"user_id"`
	User             User       `gorm:"foreignKey:UserID" json:"user"`
	DateOfBirth      time.Time  `gorm:"not null" json:"date_of_birth"`
	Gender           string     `gorm:"size:10;not null" json:"gender"`
	BloodType        string     `gorm:"size:5" json:"blood_type"`
	Address          string     `gorm:"size:255" json:"address"`
	PhoneNumber      string     `gorm:"size:15" json:"phone_number"`
	EmergencyContact string     `gorm:"size:100" json:"emergency_contact"`
	EmergencyPhone   string     `gorm:"size:15" json:"emergency_phone"`
	Insurance        string     `gorm:"size:100" json:"insurance"`
	InsuranceNo      string     `gorm:"size:50" json:"insurance_no"`
	Allergies        string     `gorm:"type:text" json:"allergies"`
	MedicalHistory   string     `gorm:"type:text" json:"medical_history"`
	MaritalStatus    string     `gorm:"size:20" json:"marital_status"`
	Occupation       string     `gorm:"size:100" json:"occupation"`
	Height           float32    `json:"height"`
	Weight           float32    `json:"weight"`
	BMI              float32    `json:"bmi"`
	ChronicDiseases  string     `gorm:"type:text" json:"chronic_diseases"`
	Medications      string     `gorm:"type:text" json:"medications"`
	Status           string     `gorm:"size:20;default:'active'" json:"status"`
	LastVisit        *time.Time `json:"last_visit"`
	NextVisit        *time.Time `json:"next_visit"`
	CreatedAt        time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time `sql:"index" json:"-"`
}

type PatientCreateRequest struct {
	UserID           uint      `json:"user_id" validate:"required"`
	DateOfBirth      time.Time `json:"date_of_birth" validate:"required"`
	Gender           string    `json:"gender" validate:"required,oneof=male female other"`
	BloodType        string    `json:"blood_type"`
	Address          string    `json:"address" validate:"required"`
	PhoneNumber      string    `json:"phone_number" validate:"required"`
	EmergencyContact string    `json:"emergency_contact" validate:"required"`
	EmergencyPhone   string    `json:"emergency_phone" validate:"required"`
	Insurance        string    `json:"insurance"`
	InsuranceNo      string    `json:"insurance_no"`
	Allergies        string    `json:"allergies"`
	MedicalHistory   string    `json:"medical_history"`
}

type PatientResponse struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	Age         int        `json:"age"`
	Gender      string     `json:"gender"`
	BloodType   string     `json:"blood_type"`
	PhoneNumber string     `json:"phone_number"`
	Insurance   string     `json:"insurance"`
	Status      string     `json:"status"`
	LastVisit   *time.Time `json:"last_visit,omitempty"`
	NextVisit   *time.Time `json:"next_visit,omitempty"`
}
