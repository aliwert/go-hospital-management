package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type DepartmentService struct {
	departmentRepo *repositories.DepartmentRepository
}

func NewDepartmentService(departmentRepo *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{
		departmentRepo: departmentRepo,
	}
}

func (s *DepartmentService) CreateDepartment(req *models.DepartmentCreateRequest) (*models.Department, error) {
	department := &models.Department{
		Name:         req.Name,
		Description:  req.Description,
		HeadDoctorID: req.HeadDoctorID,
		Location:     req.Location,
		FloorNumber:  req.FloorNumber,
		PhoneNumber:  req.PhoneNumber,
		Email:        req.Email,
		OpenTime:     req.OpenTime,
		CloseTime:    req.CloseTime,
		Capacity:     req.Capacity,
		Status:       models.DepartmentStatusActive,
		IsActive:     true,
	}

	if err := s.departmentRepo.Create(department); err != nil {
		return nil, err
	}

	return department, nil
}

func (s *DepartmentService) GetDepartmentById(id uint) (*models.Department, error) {
	return s.departmentRepo.FindById(id)
}

func (s *DepartmentService) GetAllDepartments() ([]models.Department, error) {
	return s.departmentRepo.FindAll()
}

func (s *DepartmentService) UpdateDepartment(id uint, req *models.DepartmentUpdateRequest) error {
	department, err := s.departmentRepo.FindById(id)
	if err != nil {
		return err
	}

	if req.Name != "" {
		department.Name = req.Name
	}
	if req.Description != "" {
		department.Description = req.Description
	}
	if req.HeadDoctorID != 0 {
		department.HeadDoctorID = req.HeadDoctorID
	}
	if req.Location != "" {
		department.Location = req.Location
	}
	if req.PhoneNumber != "" {
		department.PhoneNumber = req.PhoneNumber
	}
	if req.Email != "" {
		department.Email = req.Email
	}
	if req.Status != "" {
		department.Status = req.Status
	}
	if req.OpenTime != "" {
		department.OpenTime = req.OpenTime
	}
	if req.CloseTime != "" {
		department.CloseTime = req.CloseTime
	}
	if req.Capacity != 0 {
		department.Capacity = req.Capacity
	}

	return s.departmentRepo.Update(department)
}

func (s *DepartmentService) DeleteDepartment(id uint) error {
	return s.departmentRepo.Delete(id)
}
