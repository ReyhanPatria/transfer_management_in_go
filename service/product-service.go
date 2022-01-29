package service

import (
	"main/entity"
	"main/repository"
	"strings"
)

type ProductService interface {
	FindAll() []entity.Product
	FindByTypeAndKey(type_ string, key string) (entity.Product, error)
	FindMinLimitByTypeAndKey(type_ string, key string) (int, error)
	FindMaxLimitByTypeAndKey(type_ string, key string) (int, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{
		productRepository: productRepository,
	}
}

func (productService *productService) FindAll() []entity.Product {
	return productService.productRepository.FindAll()
}

func (productService *productService) FindByTypeAndKey(type_ string, key string) (entity.Product, error) {
	type_ = strings.ToUpper(type_)
	key = strings.ToUpper(key)
	return productService.productRepository.FindByTypeAndKey(type_, key)
}

func (productService *productService) FindMinLimitByTypeAndKey(type_ string, key string) (int, error) {
	type_ = strings.ToUpper(type_)
	key = strings.ToUpper(key)
	return productService.productRepository.FindMinLimitByTypeAndKey(type_, key)
}

func (productService *productService) FindMaxLimitByTypeAndKey(type_ string, key string) (int, error) {
	type_ = strings.ToUpper(type_)
	key = strings.ToUpper(key)
	return productService.productRepository.FindMaxLimitByTypeAndKey(type_, key)
}
