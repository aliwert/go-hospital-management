package models

import (
	"time"
)

const (
	CategoryMedicine  = "medicine"
	CategorySupply    = "supply"
	CategoryEquipment = "equipment"

	StatusInStock    = "in_stock"
	StatusLowStock   = "low_stock"
	StatusOutOfStock = "out_of_stock"
	StatusExpired    = "expired"
)

type Inventory struct {
	ID              uint       `gorm:"primary_key;auto_increment" json:"id"`
	ItemName        string     `gorm:"size:100;not null;unique" json:"item_name"`
	ItemCode        string     `gorm:"size:50;not null;unique" json:"item_code"`
	Category        string     `gorm:"size:50;not null" json:"category"`
	Description     string     `gorm:"type:text" json:"description"`
	Quantity        int        `gorm:"not null" json:"quantity"`
	UnitPrice       float64    `gorm:"not null" json:"unit_price"`
	ReorderLevel    int        `gorm:"not null" json:"reorder_level"`
	SupplierID      uint       `gorm:"not null" json:"supplier_id"`
	Supplier        Supplier   `gorm:"foreignKey:SupplierID" json:"supplier"`
	BatchNumber     string     `gorm:"size:50" json:"batch_number"`
	ExpiryDate      time.Time  `json:"expiry_date"`
	Location        string     `gorm:"size:100" json:"location"`
	IsActive        bool       `gorm:"default:true" json:"is_active"`
	Status          string     `gorm:"size:20;default:'in_stock'" json:"status"`
	MinimumQuantity int        `gorm:"default:0" json:"minimum_quantity"`
	MaximumQuantity int        `gorm:"default:0" json:"maximum_quantity"`
	LastOrderDate   time.Time  `json:"last_order_date"`
	LastCountDate   time.Time  `json:"last_count_date"`
	Notes           string     `gorm:"type:text" json:"notes"`
	CreatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt       *time.Time `sql:"index" json:"-"`
}

type InventoryCreateRequest struct {
	ItemName        string    `json:"item_name" validate:"required"`
	ItemCode        string    `json:"item_code" validate:"required"`
	Category        string    `json:"category" validate:"required,oneof=medicine supply equipment"`
	Description     string    `json:"description"`
	Quantity        int       `json:"quantity" validate:"required,min=0"`
	UnitPrice       float64   `json:"unit_price" validate:"required,gt=0"`
	ReorderLevel    int       `json:"reorder_level" validate:"required,min=0"`
	SupplierID      uint      `json:"supplier_id" validate:"required"`
	BatchNumber     string    `json:"batch_number"`
	ExpiryDate      time.Time `json:"expiry_date"`
	Location        string    `json:"location"`
	MinimumQuantity int       `json:"minimum_quantity"`
	MaximumQuantity int       `json:"maximum_quantity"`
}

type InventoryUpdateRequest struct {
	Quantity     *int     `json:"quantity,omitempty" validate:"omitempty,min=0"`
	UnitPrice    *float64 `json:"unit_price,omitempty" validate:"omitempty,gt=0"`
	ReorderLevel *int     `json:"reorder_level,omitempty" validate:"omitempty,min=0"`
	Location     string   `json:"location,omitempty"`
	IsActive     *bool    `json:"is_active,omitempty"`
	Status       string   `json:"status,omitempty" validate:"omitempty,oneof=in_stock low_stock out_of_stock expired"`
	Notes        string   `json:"notes,omitempty"`
}

type InventoryResponse struct {
	ID           uint      `json:"id"`
	ItemName     string    `json:"item_name"`
	ItemCode     string    `json:"item_code"`
	Category     string    `json:"category"`
	Quantity     int       `json:"quantity"`
	UnitPrice    float64   `json:"unit_price"`
	Status       string    `json:"status"`
	SupplierName string    `json:"supplier_name"`
	Location     string    `json:"location"`
	ExpiryDate   time.Time `json:"expiry_date,omitempty"`
}
