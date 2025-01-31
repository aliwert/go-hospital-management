package models

import (
	"time"
)

type DoctorSchedule struct {
	ID              uint       `gorm:"primary_key;auto_increment" json:"id"`
	DoctorID        uint       `gorm:"not null;index" json:"doctor_id"`
	Doctor          Doctor     `gorm:"foreignKey:DoctorID" json:"doctor"`
	WeekDay         int        `gorm:"not null" json:"week_day"` // 0-6 for Sunday-Saturday
	StartTime       string     `gorm:"not null" json:"start_time"`
	EndTime         string     `gorm:"not null" json:"end_time"`
	BreakStartTime  string     `gorm:"size:5" json:"break_start_time"`
	BreakEndTime    string     `gorm:"size:5" json:"break_end_time"`
	SlotDuration    int        `gorm:"default:30" json:"slot_duration"`   // in minutes
	MaxAppointments int        `gorm:"default:1" json:"max_appointments"` // per slot
	IsActive        bool       `gorm:"default:true" json:"is_active"`
	CreatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       *time.Time `sql:"index" json:"-"`
}

type ScheduleCreateRequest struct {
	DoctorID        uint   `json:"doctor_id" validate:"required"`
	WeekDay         int    `json:"week_day" validate:"required,min=0,max=6"`
	StartTime       string `json:"start_time" validate:"required,time"`
	EndTime         string `json:"end_time" validate:"required,time"`
	BreakStartTime  string `json:"break_start_time,omitempty" validate:"omitempty,time"`
	BreakEndTime    string `json:"break_end_time,omitempty" validate:"omitempty,time"`
	SlotDuration    int    `json:"slot_duration,omitempty" validate:"omitempty,min=15"`
	MaxAppointments int    `json:"max_appointments,omitempty" validate:"omitempty,min=1"`
}

type ScheduleUpdateRequest struct {
	StartTime       string `json:"start_time,omitempty" validate:"omitempty,time"`
	EndTime         string `json:"end_time,omitempty" validate:"omitempty,time"`
	BreakStartTime  string `json:"break_start_time,omitempty" validate:"omitempty,time"`
	BreakEndTime    string `json:"break_end_time,omitempty" validate:"omitempty,time"`
	SlotDuration    int    `json:"slot_duration,omitempty" validate:"omitempty,min=15"`
	MaxAppointments int    `json:"max_appointments,omitempty" validate:"omitempty,min=1"`
	IsActive        *bool  `json:"is_active,omitempty"`
}

type ScheduleResponse struct {
	ID              uint   `json:"id"`
	DoctorName      string `json:"doctor_name"`
	WeekDay         int    `json:"week_day"`
	StartTime       string `json:"start_time"`
	EndTime         string `json:"end_time"`
	BreakStartTime  string `json:"break_start_time,omitempty"`
	BreakEndTime    string `json:"break_end_time,omitempty"`
	SlotDuration    int    `json:"slot_duration"`
	MaxAppointments int    `json:"max_appointments"`
	IsActive        bool   `json:"is_active"`
}
