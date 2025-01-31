package models

import (
	"time"
)

type DoctorSchedule struct {
	ID              uint       `json:"id"`
	DoctorID        uint       `json:"doctor_id"`
	Doctor          Doctor     `json:"doctor"`
	WeekDay         int        `json:"week_day"` // 0-6 for Sunday-Saturday
	StartTime       string     `json:"start_time"`
	EndTime         string     `json:"end_time"`
	BreakStartTime  string     `json:"break_start_time"`
	BreakEndTime    string     `json:"break_end_time"`
	SlotDuration    int        `json:"slot_duration"`    // in minutes
	MaxAppointments int        `json:"max_appointments"` // per slot
	IsActive        bool       `json:"is_active"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"-"`
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
