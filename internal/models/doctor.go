package models

import (
	"time"
)

const (
	DoctorStatusActive   = "active"
	DoctorStatusInactive = "inactive"
	DoctorStatusOnLeave  = "on_leave"
)

type Doctor struct {
	ID              uint       ` json:"id"`
	UserID          uint       `json:"user_id"`
	User            User       `json:"user"`
	Specialization  string     `json:"specialization"`
	LicenseNumber   string     `json:"license_number"`
	Experience      int        `json:"experience"`
	Department      string     `json:"department"`
	Availability    bool       `json:"availability"`
	ConsultationFee float64    `json:"consultation_fee"`
	Status          string     `json:"status"`
	Education       string     `json:"education"`
	Qualifications  string     `json:"qualifications"`
	Languages       string     `json:"languages"`
	Biography       string     `json:"biography"`
	Rating          float32    `json:"rating"`
	ReviewCount     int        `json:"review_count"`
	OfficeNumber    string     `json:"office_number"`
	WorkingDays     string     `json:"working_days"`  // e.g., "1,2,3,4,5"
	WorkingHours    string     `json:"working_hours"` // e.g., "09:00-17:00"
	MaxPatients     int        `json:"max_patients"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"-"`
}

type DoctorCreateRequest struct {
	UserID          uint    `json:"user_id" validate:"required"`
	Specialization  string  `json:"specialization" validate:"required"`
	LicenseNumber   string  `json:"license_number" validate:"required"`
	Experience      int     `json:"experience" validate:"required,min=0"`
	Department      string  `json:"department" validate:"required"`
	ConsultationFee float64 `json:"consultation_fee" validate:"required,gt=0"`
	Education       string  `json:"education"`
	Qualifications  string  `json:"qualifications"`
	Languages       string  `json:"languages"`
	Biography       string  `json:"biography"`
	WorkingDays     string  `json:"working_days" validate:"required"`
	WorkingHours    string  `json:"working_hours" validate:"required"`
	MaxPatients     int     `json:"max_patients" validate:"required,min=0"`
}

type DoctorUpdateRequest struct {
	Specialization  string  `json:"specialization,omitempty"`
	ConsultationFee float64 `json:"consultation_fee,omitempty" validate:"omitempty,gt=0"`
	Status          string  `json:"status,omitempty" validate:"omitempty,oneof=active inactive on_leave"`
	Availability    *bool   `json:"availability,omitempty"`
	WorkingDays     string  `json:"working_days,omitempty"`
	WorkingHours    string  `json:"working_hours,omitempty"`
	MaxPatients     int     `json:"max_patients,omitempty" validate:"omitempty,min=0"`
}

type DoctorResponse struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	Specialization  string  `json:"specialization"`
	Department      string  `json:"department"`
	Status          string  `json:"status"`
	Availability    bool    `json:"availability"`
	ConsultationFee float64 `json:"consultation_fee"`
	Rating          float32 `json:"rating"`
	WorkingDays     string  `json:"working_days"`
	WorkingHours    string  `json:"working_hours"`
}
