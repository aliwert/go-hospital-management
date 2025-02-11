package repositories

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type PrescriptionRepository struct {
	db *gorm.DB
}

func NewPrescriptionRepository(db *gorm.DB) *PrescriptionRepository {
	return &PrescriptionRepository{db: db}
}

func (r *PrescriptionRepository) Create(prescription *models.Prescription) error {
	return r.db.Create(prescription).Error
}

func (r *PrescriptionRepository) FindById(id uint) (*models.Prescription, error) {
	var prescription models.Prescription
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Medications").First(&prescription, id).Error
	return &prescription, err
}

func (r *PrescriptionRepository) FindByPatientId(patientId uint) ([]models.Prescription, error) {
	var prescriptions []models.Prescription
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Medications").
		Where("patient_id = ?", patientId).Find(&prescriptions).Error
	return prescriptions, err
}

func (r *PrescriptionRepository) FindAll() ([]models.Prescription, error) {
	var prescriptions []models.Prescription
	err := r.db.Preload("Patient").Preload("Doctor").Preload("Medications").Find(&prescriptions).Error
	return prescriptions, err
}

func (r *PrescriptionRepository) Update(prescription *models.Prescription) error {
	return r.db.Save(prescription).Error
}

func (r *PrescriptionRepository) Delete(id uint) error {
	return r.db.Delete(&models.Prescription{}, id).Error
}
