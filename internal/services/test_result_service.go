package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type TestResultService struct {
	testResultRepo *repositories.TestResultRepository
}

func NewTestResultService(testResultRepo *repositories.TestResultRepository) *TestResultService {
	return &TestResultService{
		testResultRepo: testResultRepo,
	}
}

func (s *TestResultService) CreateTestResult(result *models.TestResult) error {
	return s.testResultRepo.Create(result)
}

func (s *TestResultService) GetTestResultById(id uint) (*models.TestResult, error) {
	return s.testResultRepo.FindById(id)
}

func (s *TestResultService) GetTestResultsByMedicalRecordId(medicalRecordId uint) ([]models.TestResult, error) {
	return s.testResultRepo.FindByMedicalRecordId(medicalRecordId)
}

func (s *TestResultService) UpdateTestResult(result *models.TestResult) error {
	return s.testResultRepo.Update(result)
}

func (s *TestResultService) DeleteTestResult(id uint) error {
	return s.testResultRepo.Delete(id)
}
