package repositories

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{db: db}
}

func (r *AppointmentRepository) Create(appointment *models.Appointment) error {
	return r.db.Create(appointment).Error
}

func (r *AppointmentRepository) FindById(id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.db.Preload("Patient").Preload("Doctor").First(&appointment, id).Error
	return &appointment, err
}

func (r *AppointmentRepository) FindAll() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Patient").Preload("Doctor").Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) FindByPatientId(patientId uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Doctor").Where("patient_id = ?", patientId).Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) FindByDoctorId(doctorId uint) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.db.Preload("Patient").Where("doctor_id = ?", doctorId).Find(&appointments).Error
	return appointments, err
}

func (r *AppointmentRepository) Update(appointment *models.Appointment) error {
	return r.db.Save(appointment).Error
}

func (r *AppointmentRepository) Delete(id uint) error {
	return r.db.Delete(&models.Appointment{}, id).Error
}
