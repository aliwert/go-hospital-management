package repositories

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"gorm.io/gorm"
)

type DoctorScheduleRepository struct {
	db *gorm.DB
}

func NewDoctorScheduleRepository(db *gorm.DB) *DoctorScheduleRepository {
	return &DoctorScheduleRepository{db: db}
}

func (r *DoctorScheduleRepository) Create(schedule *models.DoctorSchedule) error {
	return r.db.Create(schedule).Error
}

func (r *DoctorScheduleRepository) FindById(id uint) (*models.DoctorSchedule, error) {
	var schedule models.DoctorSchedule
	err := r.db.Preload("Doctor").First(&schedule, id).Error
	return &schedule, err
}

func (r *DoctorScheduleRepository) FindByDoctorId(doctorId uint) ([]models.DoctorSchedule, error) {
	var schedules []models.DoctorSchedule
	err := r.db.Where("doctor_id = ?", doctorId).Find(&schedules).Error
	return schedules, err
}

func (r *DoctorScheduleRepository) Update(schedule *models.DoctorSchedule) error {
	return r.db.Save(schedule).Error
}

func (r *DoctorScheduleRepository) Delete(id uint) error {
	return r.db.Delete(&models.DoctorSchedule{}, id).Error
}
