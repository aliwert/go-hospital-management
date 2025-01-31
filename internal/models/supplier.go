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
	ID            uint       `gorm:"primary_key;auto_increment" json:"id"`
	Name          string     `gorm:"size:100;not null;unique" json:"name"`
	Code          string     `gorm:"size:50;not null;unique" json:"code"`
	Email         string     `gorm:"size:100;not null" json:"email"`
	Phone         string     `gorm:"size:15" json:"phone"`
	Address       string     `gorm:"type:text" json:"address"`
	ContactPerson string     `gorm:"size:100" json:"contact_person"`
	ContactPhone  string     `gorm:"size:15" json:"contact_phone"`
	TaxNumber     string     `gorm:"size:50" json:"tax_number"`
	BankAccount   string     `gorm:"size:50" json:"bank_account"`
	PaymentTerms  string     `gorm:"size:100" json:"payment_terms"`
	DeliveryTerms string     `gorm:"size:100" json:"delivery_terms"`
	Website       string     `gorm:"size:100" json:"website"`
	Rating        float32    `gorm:"default:0" json:"rating"`
	Status        string     `gorm:"size:20;default:'active'" json:"status"`
	Notes         string     `gorm:"type:text" json:"notes"`
	LastOrderDate *time.Time `json:"last_order_date"`
	TotalOrders   int        `gorm:"default:0" json:"total_orders"`
	IsVerified    bool       `gorm:"default:false" json:"is_verified"`
	CreatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time `sql:"index" json:"-"`
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
