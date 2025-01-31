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
	ID              uint       `gorm:"primary_key;auto_increment" json:"id"`
	UserID          uint       `gorm:"not null" json:"user_id"`
	User            User       `gorm:"foreignKey:UserID" json:"user"`
	Specialization  string     `gorm:"size:100;not null" json:"specialization"`
	LicenseNumber   string     `gorm:"size:50;not null;unique" json:"license_number"`
	Experience      int        `gorm:"not null" json:"experience"`
	Department      string     `gorm:"size:100;not null" json:"department"`
	Availability    bool       `gorm:"default:true" json:"availability"`
	ConsultationFee float64    `gorm:"not null" json:"consultation_fee"`
	Status          string     `gorm:"type:varchar(20);default:'active'" json:"status"`
	Education       string     `gorm:"type:text" json:"education"`
	Qualifications  string     `gorm:"type:text" json:"qualifications"`
	Languages       string     `gorm:"size:100" json:"languages"`
	Biography       string     `gorm:"type:text" json:"biography"`
	Rating          float32    `gorm:"default:0" json:"rating"`
	ReviewCount     int        `gorm:"default:0" json:"review_count"`
	OfficeNumber    string     `gorm:"size:20" json:"office_number"`
	WorkingDays     string     `gorm:"size:50" json:"working_days"`  // e.g., "1,2,3,4,5"
	WorkingHours    string     `gorm:"size:50" json:"working_hours"` // e.g., "09:00-17:00"
	MaxPatients     int        `gorm:"default:0" json:"max_patients"`
	CreatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       *time.Time `sql:"index" json:"-"`
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
