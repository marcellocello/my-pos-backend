package service

import (
	"mypos-backend/internal/model"
	"mypos-backend/internal/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) GetProduct(statusStok string, categoryId int, keyword string) ([]model.Product, error) {
	return s.productRepo.GetProduct(statusStok, categoryId, keyword)
}

func (s *ProductService) GetProductByID(id int) (*model.Product, error) {
	return s.productRepo.GetProductByID(id)
}

func (s *ProductService) InsertProduct(product *model.CreateProductRequest) error {
	return s.productRepo.InsertProduct(product)
}
