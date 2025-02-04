package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type SupplierService struct {
	supplierRepo *repositories.SupplierRepository
}

func NewSupplierService(supplierRepo *repositories.SupplierRepository) *SupplierService {
	return &SupplierService{
		supplierRepo: supplierRepo,
	}
}

func (s *SupplierService) CreateSupplier(req *models.SupplierCreateRequest) (*models.Supplier, error) {
	supplier := &models.Supplier{
		Name:          req.Name,
		Code:          req.Code,
		Email:         req.Email,
		Phone:         req.Phone,
		Address:       req.Address,
		ContactPerson: req.ContactPerson,
		ContactPhone:  req.ContactPhone,
		TaxNumber:     req.TaxNumber,
		PaymentTerms:  req.PaymentTerms,
		DeliveryTerms: req.DeliveryTerms,
		Status:        models.SupplierStatusActive,
		IsVerified:    false,
	}

	if err := s.supplierRepo.Create(supplier); err != nil {
		return nil, err
	}

	return supplier, nil
}

func (s *SupplierService) GetSupplierById(id uint) (*models.Supplier, error) {
	return s.supplierRepo.FindById(id)
}

func (s *SupplierService) GetAllSuppliers() ([]models.Supplier, error) {
	return s.supplierRepo.FindAll()
}

func (s *SupplierService) UpdateSupplier(id uint, supplier *models.Supplier) error {
	return s.supplierRepo.Update(supplier)
}

func (s *SupplierService) DeleteSupplier(id uint) error {
	return s.supplierRepo.Delete(id)
}
