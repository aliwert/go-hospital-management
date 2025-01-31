package models

import (
	"time"
)

const (
	DepartmentStatusActive      = "active"
	DepartmentStatusInactive    = "inactive"
	DepartmentStatusMaintenance = "maintenance"
)

type Department struct {
	ID           uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name         string     `gorm:"size:100;not null;unique;index" json:"name"`
	Description  string     `gorm:"type:text" json:"description"`
	HeadDoctorID uint       `gorm:"not null;index" json:"head_doctor_id"`
	HeadDoctor   Doctor     `gorm:"foreignKey:HeadDoctorID" json:"head_doctor"`
	Location     string     `gorm:"size:100;not null" json:"location"`
	FloorNumber  int        `gorm:"not null" json:"floor_number"`
	PhoneNumber  string     `gorm:"size:15" json:"phone_number"`
	Email        string     `gorm:"size:100" json:"email"`
	StaffCount   int        `gorm:"default:0" json:"staff_count"`
	IsActive     bool       `gorm:"default:true" json:"is_active"`
	Status       string     `gorm:"type:varchar(20);default:'active'" json:"status"`
	OpenTime     string     `gorm:"size:5" json:"open_time"`  // Format: "HH:MM"
	CloseTime    string     `gorm:"size:5" json:"close_time"` // Format: "HH:MM"
	Capacity     int        `gorm:"default:0" json:"capacity"`
	CreatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt    *time.Time `sql:"index" json:"-"`
}

type DepartmentCreateRequest struct {
	Name         string `json:"name" validate:"required"`
	Description  string `json:"description"`
	HeadDoctorID uint   `json:"head_doctor_id" validate:"required"`
	Location     string `json:"location" validate:"required"`
	FloorNumber  int    `json:"floor_number" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	OpenTime     string `json:"open_time" validate:"required,time"`
	CloseTime    string `json:"close_time" validate:"required,time"`
	Capacity     int    `json:"capacity" validate:"required,min=0"`
}

type DepartmentUpdateRequest struct {
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
	HeadDoctorID uint   `json:"head_doctor_id,omitempty"`
	Location     string `json:"location,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty"`
	Email        string `json:"email,omitempty" validate:"omitempty,email"`
	Status       string `json:"status,omitempty" validate:"omitempty,oneof=active inactive maintenance"`
	OpenTime     string `json:"open_time,omitempty" validate:"omitempty,time"`
	CloseTime    string `json:"close_time,omitempty" validate:"omitempty,time"`
	Capacity     int    `json:"capacity,omitempty" validate:"omitempty,min=0"`
}

type DepartmentResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	HeadDoctor  string `json:"head_doctor"`
	Location    string `json:"location"`
	Status      string `json:"status"`
	StaffCount  int    `json:"staff_count"`
	OpenTime    string `json:"open_time"`
	CloseTime   string `json:"close_time"`
	PhoneNumber string `json:"phone_number"`
}
