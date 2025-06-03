package iService

import (
	repositories "richard-project-back/repositories/Models"
)

type IProductService interface {
	TesteProduct() string
	GetProduct(Id int64) string
	GetAllProducts() string
	GetProductDetail(Id int64) string
	InsertProduct(product *repositories.Product) bool
	DeleteProduct(Id int64) bool
	UpdateProduct(Id int64, product *repositories.Product) bool
}
