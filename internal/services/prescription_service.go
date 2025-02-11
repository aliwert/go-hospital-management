package services

import (
	"time"

	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type PrescriptionService struct {
	prescriptionRepo *repositories.PrescriptionRepository
}

func NewPrescriptionService(prescriptionRepo *repositories.PrescriptionRepository) *PrescriptionService {
	return &PrescriptionService{
		prescriptionRepo: prescriptionRepo,
	}
}

func (s *PrescriptionService) CreatePrescription(req *models.PrescriptionCreateRequest) (*models.Prescription, error) {
	prescription := &models.Prescription{
		PatientID:    req.PatientID,
		DoctorID:     req.DoctorID,
		Diagnosis:    req.Diagnosis,
		Notes:        req.Notes,
		IssueDate:    time.Now(),
		ValidUntil:   req.ValidUntil,
		Status:       models.PrescriptionStatusActive,
		MaxRefills:   req.MaxRefills,
		RefillCount:  0,
		IsControlled: false,
		Medications:  req.Medications,
	}

	if err := s.prescriptionRepo.Create(prescription); err != nil {
		return nil, err
	}

	return prescription, nil
}

func (s *PrescriptionService) GetPrescriptionById(id uint) (*models.Prescription, error) {
	return s.prescriptionRepo.FindById(id)
}

func (s *PrescriptionService) GetPatientPrescriptions(patientId uint) ([]models.Prescription, error) {
	return s.prescriptionRepo.FindByPatientId(patientId)
}

func (s *PrescriptionService) GetAllPrescriptions() ([]models.Prescription, error) {
	return s.prescriptionRepo.FindAll()
}

func (s *PrescriptionService) UpdatePrescription(id uint, req *models.PrescriptionUpdateRequest) (*models.Prescription, error) {
	prescription, err := s.prescriptionRepo.FindById(id)
	if err != nil {
		return nil, err
	}

	if req.Diagnosis != "" {
		prescription.Diagnosis = req.Diagnosis
	}
	if req.Notes != "" {
		prescription.Notes = req.Notes
	}
	if req.ValidUntil != nil {
		prescription.ValidUntil = *req.ValidUntil
	}
	if req.Status != "" {
		prescription.Status = req.Status
	}
	if req.MaxRefills != nil {
		prescription.MaxRefills = *req.MaxRefills
	}
	if len(req.Medications) > 0 {
		prescription.Medications = req.Medications
	}

	if err := s.prescriptionRepo.Update(prescription); err != nil {
		return nil, err
	}

	return prescription, nil
}

func (s *PrescriptionService) DeletePrescription(id uint) error {
	return s.prescriptionRepo.Delete(id)
}
