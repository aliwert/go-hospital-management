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
	ID        uint       `json:"id"`
	Name      string     `json:"name" validate:"required,min=2,max=100"`
	Email     string     `json:"email" validate:"required,email"`
	Password  string     `json:"-" validate:"required,min=6"`
	Role      string     `json:"role" validate:"required,oneof=admin doctor patient"`
	Version   uint       `json:"-"`
	CreatedBy uint       `json:"created_by"`
	UpdatedBy uint       `json:"updated_by"`
	Status    bool       `json:"status"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
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
