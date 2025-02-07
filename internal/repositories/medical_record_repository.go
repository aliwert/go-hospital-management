package repositories

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type MedicalRecordRepository struct {
	db *gorm.DB
}

func NewMedicalRecordRepository(db *gorm.DB) *MedicalRecordRepository {
	return &MedicalRecordRepository{db: db}
}

func (r *MedicalRecordRepository) Create(record *models.MedicalRecord) error {
	return r.db.Create(record).Error
}

func (r *MedicalRecordRepository) FindById(id uint) (*models.MedicalRecord, error) {
	var record models.MedicalRecord
	err := r.db.Preload("Patient").Preload("Doctor").First(&record, id).Error
	return &record, err
}

func (r *MedicalRecordRepository) FindByPatientId(patientId uint) ([]models.MedicalRecord, error) {
	var records []models.MedicalRecord
	err := r.db.Preload("Patient").Preload("Doctor").Where("patient_id = ?", patientId).Find(&records).Error
	return records, err
}

func (r *MedicalRecordRepository) FindAll() ([]models.MedicalRecord, error) {
	var records []models.MedicalRecord
	err := r.db.Preload("Patient").Preload("Doctor").Find(&records).Error
	return records, err
}

func (r *MedicalRecordRepository) Update(record *models.MedicalRecord) error {
	return r.db.Save(record).Error
}

func (r *MedicalRecordRepository) Delete(id uint) error {
	return r.db.Delete(&models.MedicalRecord{}, id).Error
}
