package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type InventoryService struct {
	inventoryRepo *repositories.InventoryRepository
}

func NewInventoryService(inventoryRepo *repositories.InventoryRepository) *InventoryService {
	return &InventoryService{
		inventoryRepo: inventoryRepo,
	}
}

func (s *InventoryService) CreateInventory(req *models.InventoryCreateRequest) (*models.Inventory, error) {
	inventory := &models.Inventory{
		ItemName:        req.ItemName,
		ItemCode:        req.ItemCode,
		Category:        req.Category,
		Description:     req.Description,
		Quantity:        req.Quantity,
		UnitPrice:       req.UnitPrice,
		ReorderLevel:    req.ReorderLevel,
		SupplierID:      req.SupplierID,
		BatchNumber:     req.BatchNumber,
		ExpiryDate:      req.ExpiryDate,
		Location:        req.Location,
		MinimumQuantity: req.MinimumQuantity,
		MaximumQuantity: req.MaximumQuantity,
		Status:          models.StatusInStock,
		IsActive:        true,
	}

	if err := s.inventoryRepo.Create(inventory); err != nil {
		return nil, err
	}

	return inventory, nil
}

func (s *InventoryService) GetInventoryById(id uint) (*models.Inventory, error) {
	return s.inventoryRepo.FindById(id)
}

func (s *InventoryService) GetAllInventory() ([]models.Inventory, error) {
	return s.inventoryRepo.FindAll()
}

func (s *InventoryService) UpdateInventory(id uint, req *models.InventoryUpdateRequest) error {
	inventory, err := s.inventoryRepo.FindById(id)
	if err != nil {
		return err
	}

	if req.Quantity != nil {
		inventory.Quantity = *req.Quantity
	}
	if req.UnitPrice != nil {
		inventory.UnitPrice = *req.UnitPrice
	}
	if req.ReorderLevel != nil {
		inventory.ReorderLevel = *req.ReorderLevel
	}
	if req.Location != "" {
		inventory.Location = req.Location
	}
	if req.IsActive != nil {
		inventory.IsActive = *req.IsActive
	}
	if req.Status != "" {
		inventory.Status = req.Status
	}

	return s.inventoryRepo.Update(inventory)
}

func (s *InventoryService) DeleteInventory(id uint) error {
	return s.inventoryRepo.Delete(id)
}
