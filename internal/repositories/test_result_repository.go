package repositories

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type TestResultRepository struct {
	db *gorm.DB
}

func NewTestResultRepository(db *gorm.DB) *TestResultRepository {
	return &TestResultRepository{db: db}
}

func (r *TestResultRepository) Create(result *models.TestResult) error {
	return r.db.Create(result).Error
}

func (r *TestResultRepository) FindById(id uint) (*models.TestResult, error) {
	var result models.TestResult
	err := r.db.First(&result, id).Error
	return &result, err
}

func (r *TestResultRepository) FindByMedicalRecordId(medicalRecordId uint) ([]models.TestResult, error) {
	var results []models.TestResult
	err := r.db.Where("medical_record_id = ?", medicalRecordId).Find(&results).Error
	return results, err
}

func (r *TestResultRepository) Update(result *models.TestResult) error {
	return r.db.Save(result).Error
}

func (r *TestResultRepository) Delete(id uint) error {
	return r.db.Delete(&models.TestResult{}, id).Error
}
