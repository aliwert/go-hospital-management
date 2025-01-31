package models

import (
	"time"
)

type MedicalRecord struct {
	ID              uint         `gorm:"primary_key;auto_increment" json:"id"`
	PatientID       uint         `gorm:"not null;index" json:"patient_id"`
	Patient         Patient      `gorm:"foreignKey:PatientID" json:"patient"`
	DoctorID        uint         `gorm:"not null;index" json:"doctor_id"`
	Doctor          Doctor       `gorm:"foreignKey:DoctorID" json:"doctor"`
	VisitDate       time.Time    `gorm:"not null;index" json:"visit_date"`
	Diagnosis       string       `gorm:"type:text;not null" json:"diagnosis"`
	Symptoms        string       `gorm:"type:text" json:"symptoms"`
	Treatment       string       `gorm:"type:text" json:"treatment"`
	Notes           string       `gorm:"type:text" json:"notes"`
	PrescriptionID  uint         `json:"prescription_id"`
	Prescription    Prescription `gorm:"foreignKey:PrescriptionID" json:"prescription"`
	TestResults     []TestResult `gorm:"foreignKey:MedicalRecordID" json:"test_results"`
	BloodPressure   string       `gorm:"size:20" json:"blood_pressure"`
	Temperature     float32      `json:"temperature"`
	Weight          float32      `json:"weight"`
	Height          float32      `json:"height"`
	BMI             float32      `json:"bmi"`
	PulseRate       int          `json:"pulse_rate"`
	RespiratoryRate int          `json:"respiratory_rate"`
	Allergies       string       `gorm:"type:text" json:"allergies"`
	Medications     string       `gorm:"type:text" json:"medications"`
	FollowUpDate    *time.Time   `json:"follow_up_date"`
	IsFollowUp      bool         `gorm:"default:false" json:"is_follow_up"`
	Attachments     string       `gorm:"type:text" json:"attachments"` // JSON string of file paths
	Status          string       `gorm:"size:20;default:'active'" json:"status"`
	CreatedAt       time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time    `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       *time.Time   `sql:"index" json:"-"`
}

type MedicalRecordCreateRequest struct {
	PatientID     uint       `json:"patient_id" validate:"required"`
	DoctorID      uint       `json:"doctor_id" validate:"required"`
	VisitDate     time.Time  `json:"visit_date" validate:"required"`
	Diagnosis     string     `json:"diagnosis" validate:"required"`
	Symptoms      string     `json:"symptoms"`
	Treatment     string     `json:"treatment"`
	Notes         string     `json:"notes"`
	BloodPressure string     `json:"blood_pressure"`
	Temperature   float32    `json:"temperature"`
	Weight        float32    `json:"weight"`
	Height        float32    `json:"height"`
	Allergies     string     `json:"allergies"`
	Medications   string     `json:"medications"`
	FollowUpDate  *time.Time `json:"follow_up_date"`
}

type MedicalRecordResponse struct {
	ID           uint       `json:"id"`
	PatientName  string     `json:"patient_name"`
	DoctorName   string     `json:"doctor_name"`
	VisitDate    time.Time  `json:"visit_date"`
	Diagnosis    string     `json:"diagnosis"`
	Treatment    string     `json:"treatment"`
	FollowUpDate *time.Time `json:"follow_up_date,omitempty"`
	Status       string     `json:"status"`
}
