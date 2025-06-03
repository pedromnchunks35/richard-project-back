package services

import (
	model "richard-project-back/repositories/Models"
	queryList "richard-project-back/repositories/Queries"
	iService "richard-project-back/services/ServicesInterface"
)

type ProductService struct {
}

func NewProductService() iService.IProductService {
	return &ProductService{}
}

func (s *ProductService) TesteProduct() string {
	return "GOOD"
}

//GET ----

func (s *ProductService) GetProduct(id int64) string {
	var query = queryList.QueryGetProductByID
	var json = ExecuteQuery(query)
	return json
}
func (s *ProductService) GetAllProducts() string {
	var query = queryList.QueryGetProductByID
	var json = ExecuteQuery(query)
	return json
}
func (s *ProductService) GetProductDetail(id int64) string {
	var query = queryList.QueryGetDetailsByProductID
	var json = ExecuteQuery(query)
	return json
}

//INSERT ----

func (s *ProductService) InsertProduct(product *model.Product) bool {
	if product == nil {
		return false
	}
	var query = queryList.QueryInsertProduct
	var json = ExecuteQuery(query)

	return json == ""
}

// UPDATE ----
func (s *ProductService) UpdateProduct(id int64, product *model.Product) bool {
	if id == 0 || product == nil {
		return false
	}
	var query = queryList.QueryUpdateProduct
	var json = ExecuteQuery(query)

	return json == ""
}

// DELETE ----
func (s *ProductService) DeleteProduct(id int64) bool {
	if id == 0 {
		return false
	}
	var query = queryList.QueryInsertProduct
	var json = ExecuteQuery(query)
	return json == ""
}
