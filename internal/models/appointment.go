package models

import (
	"time"
)

const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusCancelled = "cancelled"
	StatusCompleted = "completed"

	PaymentStatusUnpaid = "unpaid"
	PaymentStatusPaid   = "paid"
)

type Appointment struct {
	ID              uint       `json:"id"`
	PatientID       uint       `json:"patient_id"`
	Patient         User       `json:"patient"`
	DoctorID        uint       `json:"doctor_id"`
	Doctor          Doctor     `json:"doctor"`
	AppointmentDate time.Time  `json:"appointment_date"`
	Status          string     `json:"status"`
	Description     string     `json:"description"`
	Fee             float64    `json:"fee"`
	PaymentStatus   string     `json:"payment_status"`
	PaymentDate     *time.Time `json:"payment_date,omitempty"`
	CancelledAt     *time.Time `json:"cancelled_at,omitempty"`
	CancelReason    string     `json:"cancel_reason,omitempty"`
	Notes           string     `json:"notes,omitempty"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"-"`
}

type AppointmentCreateRequest struct {
	PatientID       uint      `json:"patient_id" validate:"required"`
	DoctorID        uint      `json:"doctor_id" validate:"required"`
	AppointmentDate time.Time `json:"appointment_date" validate:"required,future"`
	Description     string    `json:"description"`
	Fee             float64   `json:"fee" validate:"required,gt=0"`
}

type AppointmentUpdateRequest struct {
	Status        string `json:"status,omitempty" validate:"omitempty,oneof=pending confirmed cancelled completed"`
	Description   string `json:"description,omitempty"`
	PaymentStatus string `json:"payment_status,omitempty" validate:"omitempty,oneof=unpaid paid"`
	CancelReason  string `json:"cancel_reason,omitempty"`
	Notes         string `json:"notes,omitempty"`
}

type AppointmentResponse struct {
	ID              uint      `json:"id"`
	PatientName     string    `json:"patient_name"`
	DoctorName      string    `json:"doctor_name"`
	AppointmentDate time.Time `json:"appointment_date"`
	Status          string    `json:"status"`
	Fee             float64   `json:"fee"`
	PaymentStatus   string    `json:"payment_status"`
	CreatedAt       time.Time `json:"created_at"`
}
