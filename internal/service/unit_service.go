package service

import (
	"errors"
	"mypos-backend/internal/model"
	"mypos-backend/internal/repository"
)

type UnitService struct {
	unitRepo *repository.UnitRepository
}

func NewUnitService(unitRepo *repository.UnitRepository) *UnitService {
	return &UnitService{unitRepo: unitRepo}
}

func (s *UnitService) GetAllUnits() ([]model.Unit, error) {
	return s.unitRepo.GetAllUnits()
}

func (s *UnitService) InsertUnit(req *model.CreateUnitRequest) error {
	unit := &model.CreateUnitRequest{
		Name:        req.Name,
		Description: req.Description,
	}

	return s.unitRepo.InsertUnit(unit)
}

func (s *UnitService) GetUnitByID(id int) (*model.Unit, error) {
	return s.unitRepo.GetUnitByID(id)
}

func (s *UnitService) UpdateUnit(req *model.UpdateUnitRequest) error {
	unit, err := s.unitRepo.GetUnitByID(req.ID)
	if err != nil {
		return err
	}

	if unit == nil {
		return errors.New("unit tidak ditemukan")
	}

	return s.unitRepo.UpdateUnit(req.ID, req)
}

func (s *UnitService) DeleteUnit(id int) error {
	unit, err := s.unitRepo.GetUnitByID(id)
	if err != nil {
		return err
	}

	if unit == nil {
		return errors.New("unit tidak ditemukan")
	}

	return s.unitRepo.DeleteUnit(id)
}
