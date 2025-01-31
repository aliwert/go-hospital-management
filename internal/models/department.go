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
	ID           uint       `json:"id"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	HeadDoctorID uint       `json:"head_doctor_id"`
	HeadDoctor   Doctor     `json:"head_doctor"`
	Location     string     `json:"location"`
	FloorNumber  int        `json:"floor_number"`
	PhoneNumber  string     `json:"phone_number"`
	Email        string     `json:"email"`
	StaffCount   int        `json:"staff_count"`
	IsActive     bool       `json:"is_active"`
	Status       string     `json:"status"`
	OpenTime     string     `json:"open_time"`  // Format: "HH:MM"
	CloseTime    string     `json:"close_time"` // Format: "HH:MM"
	Capacity     int        `json:"capacity"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"-"`
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
