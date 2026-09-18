package service

import (
	"errors"
	"mypos-backend/internal/model"
	"mypos-backend/internal/repository"
)

type ProductService struct {
	unitRepo *repository.ProductRepository
}

func NewProductService(unitRepo *repository.ProductRepository) *ProductService {
	return &ProductService{unitRepo: unitRepo}
}

func (s *ProductService) GetAllUnits() ([]model.Unit, error) {
	return s.unitRepo.GetAllUnits()
}

func (s *ProductService) InsertUnit(req *model.CreateUnitRequest) error {
	unit := &model.CreateUnitRequest{
		Name:        req.Name,
		Description: req.Description,
	}

	return s.unitRepo.InsertUnit(unit)
}

func (s *ProductService) GetUnitByID(id int) (*model.Unit, error) {
	return s.unitRepo.GetUnitByID(id)
}

func (s *ProductService) UpdateUnit(req *model.UpdateUnitRequest) error {
	unit, err := s.unitRepo.GetUnitByID(req.ID)
	if err != nil {
		return err
	}

	if unit == nil {
		return errors.New("unit tidak ditemukan")
	}

	return s.unitRepo.UpdateUnit(req.ID, req)
}

func (s *ProductService) DeleteUnit(id int) error {
	unit, err := s.unitRepo.GetUnitByID(id)
	if err != nil {
		return err
	}

	if unit == nil {
		return errors.New("unit tidak ditemukan")
	}

	return s.unitRepo.DeleteUnit(id)
}
