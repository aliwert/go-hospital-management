package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type MedicalRecordService struct {
	medicalRecordRepo *repositories.MedicalRecordRepository
}

func NewMedicalRecordService(medicalRecordRepo *repositories.MedicalRecordRepository) *MedicalRecordService {
	return &MedicalRecordService{
		medicalRecordRepo: medicalRecordRepo,
	}
}

func (s *MedicalRecordService) CreateMedicalRecord(req *models.MedicalRecordCreateRequest) (*models.MedicalRecord, error) {
	record := &models.MedicalRecord{
		PatientID:     req.PatientID,
		DoctorID:      req.DoctorID,
		VisitDate:     req.VisitDate,
		Diagnosis:     req.Diagnosis,
		Symptoms:      req.Symptoms,
		Treatment:     req.Treatment,
		Notes:         req.Notes,
		BloodPressure: req.BloodPressure,
		Temperature:   req.Temperature,
		Weight:        req.Weight,
		Height:        req.Height,
		Allergies:     req.Allergies,
		Medications:   req.Medications,
		FollowUpDate:  req.FollowUpDate,
		Status:        "active",
	}

	if err := s.medicalRecordRepo.Create(record); err != nil {
		return nil, err
	}

	return record, nil
}

func (s *MedicalRecordService) GetMedicalRecordById(id uint) (*models.MedicalRecord, error) {
	return s.medicalRecordRepo.FindById(id)
}

func (s *MedicalRecordService) GetAllMedicalRecords() ([]models.MedicalRecord, error) {
	return s.medicalRecordRepo.FindAll()
}

func (s *MedicalRecordService) GetPatientMedicalRecords(patientId uint) ([]models.MedicalRecord, error) {
	return s.medicalRecordRepo.FindByPatientId(patientId)
}

func (s *MedicalRecordService) UpdateMedicalRecord(id uint, record *models.MedicalRecord) error {
	return s.medicalRecordRepo.Update(record)
}

func (s *MedicalRecordService) DeleteMedicalRecord(id uint) error {
	return s.medicalRecordRepo.Delete(id)
}
