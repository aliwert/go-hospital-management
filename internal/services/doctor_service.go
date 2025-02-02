package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type DoctorService struct {
	doctorRepo *repositories.DoctorRepository
}

func NewDoctorService(doctorRepo *repositories.DoctorRepository) *DoctorService {
	return &DoctorService{
		doctorRepo: doctorRepo,
	}
}

func (s *DoctorService) CreateDoctor(req *models.DoctorCreateRequest) (*models.Doctor, error) {
	doctor := &models.Doctor{
		UserID:          req.UserID,
		Name:            req.Name,
		Specialization:  req.Specialization,
		LicenseNumber:   req.LicenseNumber,
		Experience:      req.Experience,
		Department:      req.Department,
		ConsultationFee: req.ConsultationFee,
		Education:       req.Education,
		Qualifications:  req.Qualifications,
		Languages:       req.Languages,
		Biography:       req.Biography,
		WorkingDays:     req.WorkingDays,
		WorkingHours:    req.WorkingHours,
		MaxPatients:     req.MaxPatients,
		Status:          models.DoctorStatusActive,
		Availability:    true,
	}

	if err := s.doctorRepo.Create(doctor); err != nil {
		return nil, err
	}

	return doctor, nil
}

func (s *DoctorService) GetDoctorById(id uint) (*models.Doctor, error) {
	return s.doctorRepo.FindById(id)
}

func (s *DoctorService) GetAllDoctors() ([]models.Doctor, error) {
	return s.doctorRepo.FindAll()
}

func (s *DoctorService) UpdateDoctor(id uint, req *models.DoctorUpdateRequest) error {
	doctor, err := s.doctorRepo.FindById(id)
	if err != nil {
		return err
	}

	// Update all possible fields
	if req.Specialization != "" {
		doctor.Specialization = req.Specialization
	}
	if req.Experience > 0 {
		doctor.Experience = req.Experience
	}
	if req.Department != "" {
		doctor.Department = req.Department
	}
	if req.ConsultationFee > 0 {
		doctor.ConsultationFee = req.ConsultationFee
	}
	if req.Rating > 0 {
		doctor.Rating = req.Rating
	}
	if req.ReviewCount > 0 {
		doctor.ReviewCount = req.ReviewCount
	}
	if req.Qualifications != "" {
		doctor.Qualifications = req.Qualifications
	}
	if req.Languages != "" {
		doctor.Languages = req.Languages
	}
	if req.Biography != "" {
		doctor.Biography = req.Biography
	}
	if req.OfficeNumber != "" {
		doctor.OfficeNumber = req.OfficeNumber
	}
	if req.WorkingDays != "" {
		doctor.WorkingDays = req.WorkingDays
	}
	if req.WorkingHours != "" {
		doctor.WorkingHours = req.WorkingHours
	}
	if req.MaxPatients > 0 {
		doctor.MaxPatients = req.MaxPatients
	}
	if req.Status != "" {
		doctor.Status = req.Status
	}
	if req.Availability != nil {
		doctor.Availability = *req.Availability
	}

	return s.doctorRepo.Update(doctor)
}

func (s *DoctorService) DeleteDoctor(id uint) error {
	return s.doctorRepo.Delete(id)
}
