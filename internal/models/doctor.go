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
	Name            string     `json:"name" gorm:"unique"`
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
	Name            string  `json:"name" validate:"required"`
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
	Name            string  `json:"name,omitempty" validate:"omitempty,min=2"`
	Specialization  string  `json:"specialization,omitempty" validate:"omitempty,min=2"`
	Experience      int     `json:"experience,omitempty" validate:"omitempty,min=0"`
	Department      string  `json:"department,omitempty" validate:"omitempty,min=2"`
	ConsultationFee float64 `json:"consultation_fee,omitempty" validate:"omitempty,gt=0"`
	Qualifications  string  `json:"qualifications,omitempty"`
	Languages       string  `json:"languages,omitempty"`
	Biography       string  `json:"biography,omitempty"`
	OfficeNumber    string  `json:"office_number,omitempty"`
	WorkingDays     string  `json:"working_days,omitempty" validate:"omitempty,min=1"`
	WorkingHours    string  `json:"working_hours,omitempty" validate:"omitempty,min=5"`
	MaxPatients     int     `json:"max_patients,omitempty" validate:"omitempty,min=0"`
	Status          string  `json:"status,omitempty" validate:"omitempty,oneof=active inactive on_leave"`
	Availability    *bool   `json:"availability,omitempty"`
	Rating          float32 `json:"rating,omitempty" validate:"omitempty,min=0,max=5"`
	ReviewCount     int     `json:"review_count,omitempty" validate:"omitempty,min=0"`
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
