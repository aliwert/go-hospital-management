package models

import (
	"time"
)

const (
	SupplierStatusActive   = "active"
	SupplierStatusInactive = "inactive"
	SupplierStatusBlocked  = "blocked"
)

type Supplier struct {
	ID            uint       `json:"id"`
	Name          string     `json:"name"`
	Code          string     `json:"code"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	Address       string     `json:"address"`
	ContactPerson string     `json:"contact_person"`
	ContactPhone  string     `json:"contact_phone"`
	TaxNumber     string     `json:"tax_number"`
	BankAccount   string     `json:"bank_account"`
	PaymentTerms  string     `json:"payment_terms"`
	DeliveryTerms string     `json:"delivery_terms"`
	Website       string     `json:"website"`
	Rating        float32    `json:"rating"`
	Status        string     `json:"status"`
	Notes         string     `json:"notes"`
	LastOrderDate *time.Time `json:"last_order_date"`
	TotalOrders   int        `json:"total_orders"`
	IsVerified    bool       `json:"is_verified"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"-"`
}

type SupplierCreateRequest struct {
	Name          string `json:"name" validate:"required"`
	Code          string `json:"code" validate:"required"`
	Email         string `json:"email" validate:"required,email"`
	Phone         string `json:"phone" validate:"required"`
	Address       string `json:"address" validate:"required"`
	ContactPerson string `json:"contact_person" validate:"required"`
	ContactPhone  string `json:"contact_phone" validate:"required"`
	TaxNumber     string `json:"tax_number"`
	PaymentTerms  string `json:"payment_terms"`
	DeliveryTerms string `json:"delivery_terms"`
}

type SupplierResponse struct {
	ID            uint    `json:"id"`
	Name          string  `json:"name"`
	Code          string  `json:"code"`
	ContactPerson string  `json:"contact_person"`
	Phone         string  `json:"phone"`
	Email         string  `json:"email"`
	Status        string  `json:"status"`
	Rating        float32 `json:"rating"`
	IsVerified    bool    `json:"is_verified"`
}

// Add after SupplierCreateRequest
type SupplierUpdateRequest struct {
	Name          string   `json:"name,omitempty" validate:"omitempty,min=2"`
	Email         string   `json:"email,omitempty" validate:"omitempty,email"`
	Phone         string   `json:"phone,omitempty"`
	Address       string   `json:"address,omitempty"`
	ContactPerson string   `json:"contact_person,omitempty"`
	ContactPhone  string   `json:"contact_phone,omitempty"`
	TaxNumber     string   `json:"tax_number,omitempty"`
	BankAccount   string   `json:"bank_account,omitempty"`
	PaymentTerms  string   `json:"payment_terms,omitempty"`
	DeliveryTerms string   `json:"delivery_terms,omitempty"`
	Website       string   `json:"website,omitempty"`
	Rating        *float32 `json:"rating,omitempty" validate:"omitempty,min=0,max=5"`
	Status        string   `json:"status,omitempty" validate:"omitempty,oneof=active inactive blocked"`
	Notes         string   `json:"notes,omitempty"`
	IsVerified    *bool    `json:"is_verified,omitempty"`
}
