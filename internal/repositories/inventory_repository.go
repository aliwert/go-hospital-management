package repositories

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (r *InventoryRepository) Create(inventory *models.Inventory) error {
	return r.db.Create(inventory).Error
}

func (r *InventoryRepository) FindById(id uint) (*models.Inventory, error) {
	var inventory models.Inventory
	err := r.db.Preload("Supplier").First(&inventory, id).Error
	return &inventory, err
}

func (r *InventoryRepository) FindAll() ([]models.Inventory, error) {
	var inventories []models.Inventory
	err := r.db.Preload("Supplier").Find(&inventories).Error
	return inventories, err
}

func (r *InventoryRepository) Update(inventory *models.Inventory) error {
	return r.db.Save(inventory).Error
}

func (r *InventoryRepository) Delete(id uint) error {
	return r.db.Delete(&models.Inventory{}, id).Error
}
