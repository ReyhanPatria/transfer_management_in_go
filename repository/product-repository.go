package repository

import (
	"main/connect"
	"main/entity"
)

type ProductRepository interface {
	FindAll() []entity.Product
	FindByTypeAndKey(type_ string, key string) (entity.Product, error)
	FindMinLimitByTypeAndKey(type_ string, key string) (int, error)
	FindMaxLimitByTypeAndKey(type_ string, key string) (int, error)
}

type productRepository struct {
	database connect.Database
}

func NewProductRepository(database connect.Database) ProductRepository {
	return &productRepository{
		database: database,
	}
}

func (productRepository *productRepository) FindAll() []entity.Product {
	var products []entity.Product
	productRepository.database.Instance.Find(&products)
	return products
}

func (productRepository *productRepository) FindByTypeAndKey(type_ string, key string) (entity.Product, error) {
	var product entity.Product
	err := productRepository.database.Instance.
		First(&product, entity.Product{
			Key:  key,
			Type: type_,
		})

	return product, err.Error
}

func (productRepository *productRepository) FindMinLimitByTypeAndKey(type_ string, key string) (int, error) {
	var product entity.Product
	err := productRepository.database.Instance.
		Select("`MINLIMIT`").
		First(&product, entity.Product{
			Type: type_,
			Key:  key,
		})

	return product.MinLimit, err.Error
}

func (productRepository *productRepository) FindMaxLimitByTypeAndKey(type_ string, key string) (int, error) {
	var product entity.Product
	err := productRepository.database.Instance.
		Select("`MAXLIMIT`").
		First(&product, entity.Product{
			Type: type_,
			Key:  key,
		})
	return product.MaxLimit, err.Error
}
