package services

import "richard-project-back/repositories"

type ProductService interface {
	Teste() string
	GetProduct(Id int64) repositories.Product
	InsertProduct(product repositories.Product) bool
	RemoveProduct(Id int64) bool
	UpdateProduct(Id int64,product repositories.Product) bool
}
