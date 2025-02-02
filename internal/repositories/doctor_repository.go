package repositories

import (
	"errors"

	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type DoctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) Create(doctor *models.Doctor) error {
	// Check if doctor with same name exists
	var count int64
	if err := r.db.Model(&models.Doctor{}).Where("name = ?", doctor.Name).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("doctor with this name already exists")
	}

	// Check if license number is unique
	if err := r.db.Model(&models.Doctor{}).Where("license_number = ?", doctor.LicenseNumber).Count(&count).Error; err != nil {
		return err
	}

	if count > 0 {
		return errors.New("license number already exists")
	}

	return r.db.Create(doctor).Error
}

func (r *DoctorRepository) FindById(id uint) (*models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.Preload("User").First(&doctor, id).Error
	return &doctor, err
}

func (r *DoctorRepository) FindAll() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := r.db.Preload("User").Find(&doctors).Error
	return doctors, err
}

func (r *DoctorRepository) Update(doctor *models.Doctor) error {
	return r.db.Save(doctor).Error
}

func (r *DoctorRepository) Delete(id uint) error {
	return r.db.Delete(&models.Doctor{}, id).Error
}
