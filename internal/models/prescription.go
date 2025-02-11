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
	ID           uint                     `json:"id"`
	PatientID    uint                     `json:"patient_id"`
	Patient      Patient                  `json:"patient"`
	DoctorID     uint                     `json:"doctor_id"`
	Doctor       Doctor                   `json:"doctor"`
	Diagnosis    string                   `json:"diagnosis"`
	Notes        string                   `json:"notes"`
	IssueDate    time.Time                `json:"issue_date"`
	ValidUntil   time.Time                `json:"valid_until"`
	Status       string                   `json:"status"`
	PharmacyID   *uint                    `json:"pharmacy_id,omitempty"`
	FilledDate   *time.Time               `json:"filled_date,omitempty"`
	RefillCount  int                      `json:"refill_count"`
	MaxRefills   int                      `json:"max_refills"`
	IsControlled bool                     `json:"is_controlled"`
	Medications  []PrescriptionMedication `json:"medications"`
	CreatedAt    time.Time                `json:"created_at"`
	UpdatedAt    time.Time                `json:"updated_at"`
	DeletedAt    *time.Time               `json:"-"`
}

type PrescriptionMedication struct {
	ID             uint      `json:"id"`
	PrescriptionID uint      `json:"prescription_id"`
	MedicineName   string    `json:"medicine_name"`
	Dosage         string    `json:"dosage"`
	Frequency      string    `json:"frequency"`
	Duration       string    `json:"duration"`
	Instructions   string    `json:"instructions"`
	Quantity       int       `json:"quantity"`
	Substitution   bool      `json:"substitution"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
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

type PrescriptionUpdateRequest struct {
	Diagnosis   string                   `json:"diagnosis,omitempty"`
	Notes       string                   `json:"notes,omitempty"`
	ValidUntil  *time.Time               `json:"valid_until,omitempty"`
	Status      string                   `json:"status,omitempty" validate:"omitempty,oneof=active expired voided filled"`
	MaxRefills  *int                     `json:"max_refills,omitempty" validate:"omitempty,min=0"`
	Medications []PrescriptionMedication `json:"medications,omitempty" validate:"omitempty,min=1"`
}
