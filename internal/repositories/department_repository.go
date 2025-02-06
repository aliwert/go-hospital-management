package repositories

import (
	"errors"

	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) Create(department *models.Department) error {
	// check if department with same name exists
	var count int64
	if err := r.db.Model(&models.Department{}).Where("name = ?", department.Name).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("department with this name already exists")
	}

	return r.db.Create(department).Error
}

func (r *DepartmentRepository) FindById(id uint) (*models.Department, error) {
	var department models.Department
	err := r.db.Preload("HeadDoctor").First(&department, id).Error
	return &department, err
}

func (r *DepartmentRepository) FindAll() ([]models.Department, error) {
	var departments []models.Department
	err := r.db.Preload("HeadDoctor").Find(&departments).Error
	return departments, err
}

func (r *DepartmentRepository) Update(department *models.Department) error {
	return r.db.Save(department).Error
}

func (r *DepartmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Department{}, id).Error
}
