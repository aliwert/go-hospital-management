package models

import (
	"time"
)

const (
	RoleAdmin   = "admin"
	RoleDoctor  = "doctor"
	RolePatient = "patient"
)

type User struct {
	ID        uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name      string     `gorm:"size:100;not null;index" json:"name" validate:"required,min=2,max=100"`
	Email     string     `gorm:"size:100;not null;unique;index" json:"email" validate:"required,email"`
	Password  string     `gorm:"size:100;not null;" json:"-" validate:"required,min=6"`
	Role      string     `gorm:"type:varchar(10);not null;index" json:"role" validate:"required,oneof=admin doctor patient"`
	Version   uint       `gorm:"default:1" json:"-"`
	CreatedBy uint       `gorm:"not null" json:"created_by"`
	UpdatedBy uint       `gorm:"not null" json:"updated_by"`
	Status    bool       `gorm:"default:true" json:"status"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin doctor patient"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateRequest struct {
	Name     string `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
	Email    string `json:"email,omitempty" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6"`
	Status   *bool  `json:"status,omitempty"`
}

type UserResponse struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Role      string     `json:"role"`
	Status    bool       `json:"status"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
}
