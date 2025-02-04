package repositories

import (
	"errors"

	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type SupplierRepository struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

func (r *SupplierRepository) Create(supplier *models.Supplier) error {
	var count int64
	if err := r.db.Model(&models.Supplier{}).Where("code = ?", supplier.Code).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("supplier code already exists")
	}

	return r.db.Create(supplier).Error
}

func (r *SupplierRepository) FindById(id uint) (*models.Supplier, error) {
	var supplier models.Supplier
	err := r.db.First(&supplier, id).Error
	return &supplier, err
}

func (r *SupplierRepository) FindAll() ([]models.Supplier, error) {
	var suppliers []models.Supplier
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *SupplierRepository) Update(supplier *models.Supplier) error {
	return r.db.Save(supplier).Error
}

func (r *SupplierRepository) Delete(id uint) error {
	return r.db.Delete(&models.Supplier{}, id).Error
}
