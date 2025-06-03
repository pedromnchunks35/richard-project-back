package iService

import (
	repositories "richard-project-back/repositories/Models"
)

type IProductService interface {
	TesteProduct() string
	GetProduct(Id int64) ([]byte, error)
	GetAllProducts() ([]byte, error)
	GetProductDetail(Id int64) ([]byte, error)
	InsertProduct(product *repositories.Product) bool
	DeleteProduct(Id int64) bool
	UpdateProduct(Id int64, product *repositories.Product) bool
}
