package models

import (
	"time"
)

type MedicalRecord struct {
	ID              uint         `json:"id"`
	PatientID       uint         `json:"patient_id"`
	Patient         Patient      `json:"patient"`
	DoctorID        uint         `json:"doctor_id"`
	Doctor          Doctor       `json:"doctor"`
	VisitDate       time.Time    `json:"visit_date"`
	Diagnosis       string       `json:"diagnosis"`
	Symptoms        string       `json:"symptoms"`
	Treatment       string       `json:"treatment"`
	Notes           string       `json:"notes"`
	PrescriptionID  uint         `json:"prescription_id"`
	Prescription    Prescription `json:"prescription"`
	TestResults     []TestResult `json:"test_results"`
	BloodPressure   string       `json:"blood_pressure"`
	Temperature     float32      `json:"temperature"`
	Weight          float32      `json:"weight"`
	Height          float32      `json:"height"`
	BMI             float32      `json:"bmi"`
	PulseRate       int          `json:"pulse_rate"`
	RespiratoryRate int          `json:"respiratory_rate"`
	Allergies       string       `json:"allergies"`
	Medications     string       `json:"medications"`
	FollowUpDate    *time.Time   `json:"follow_up_date"`
	IsFollowUp      bool         `json:"is_follow_up"`
	Attachments     string       `json:"attachments"`
	Status          string       `json:"status"`
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
