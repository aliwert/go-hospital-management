package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type PatientService struct {
	patientRepo *repositories.PatientRepository
}

func NewPatientService(patientRepo *repositories.PatientRepository) *PatientService {
	return &PatientService{
		patientRepo: patientRepo,
	}
}

func (s *PatientService) CreatePatient(req *models.PatientCreateRequest) (*models.Patient, error) {
	patient := &models.Patient{
		UserID:           req.UserID,
		DateOfBirth:      req.DateOfBirth,
		Gender:           req.Gender,
		BloodType:        req.BloodType,
		Address:          req.Address,
		PhoneNumber:      req.PhoneNumber,
		EmergencyContact: req.EmergencyContact,
		EmergencyPhone:   req.EmergencyPhone,
		Insurance:        req.Insurance,
		InsuranceNo:      req.InsuranceNo,
		Allergies:        req.Allergies,
		MedicalHistory:   req.MedicalHistory,
		Status:           models.StatusActive,
	}

	if err := s.patientRepo.Create(patient); err != nil {
		return nil, err
	}

	return patient, nil
}

func (s *PatientService) GetPatientById(id uint) (*models.Patient, error) {
	return s.patientRepo.FindById(id)
}

func (s *PatientService) GetAllPatients() ([]models.Patient, error) {
	return s.patientRepo.FindAll()
}

func (s *PatientService) UpdatePatient(id uint, patient *models.Patient) error {
	return s.patientRepo.Update(patient)
}

func (s *PatientService) DeletePatient(id uint) error {
	return s.patientRepo.Delete(id)
}
