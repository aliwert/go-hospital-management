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
	ID               uint       `json:"id"`
	UserID           uint       `json:"user_id"`
	User             User       `json:"user"`
	DateOfBirth      time.Time  `json:"date_of_birth"`
	Gender           string     `json:"gender"`
	BloodType        string     `json:"blood_type"`
	Address          string     `json:"address"`
	PhoneNumber      string     `json:"phone_number"`
	EmergencyContact string     `json:"emergency_contact"`
	EmergencyPhone   string     `json:"emergency_phone"`
	Insurance        string     `json:"insurance"`
	InsuranceNo      string     `json:"insurance_no"`
	Allergies        string     `json:"allergies"`
	MedicalHistory   string     `json:"medical_history"`
	MaritalStatus    string     `json:"marital_status"`
	Occupation       string     `json:"occupation"`
	Height           float32    `json:"height"`
	Weight           float32    `json:"weight"`
	BMI              float32    `json:"bmi"`
	ChronicDiseases  string     `json:"chronic_diseases"`
	Medications      string     `json:"medications"`
	Status           string     `json:"status"`
	LastVisit        *time.Time `json:"last_visit"`
	NextVisit        *time.Time `json:"next_visit"`
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        time.Time  `json:"updated_at"`
	DeletedAt        *time.Time `json:"-"`
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
