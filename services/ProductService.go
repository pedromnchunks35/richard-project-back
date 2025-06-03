package services

import (
	sql "richard-project-back/configs/connector/postgresql"
		"richard-project-back/helper"
	model "richard-project-back/repositories/Models"
	queryList "richard-project-back/repositories/Queries"
	iService "richard-project-back/services/ServicesInterface"
)

type ProductService struct {
	db *sql.PostgresDbVault
}

func NewProductService(db *sql.PostgresDbVault) iService.IProductService {
	return &ProductService{
		db:db
	}
}

func (s *ProductService) TesteProduct() string {
	return "GOOD"
}

//GET ----

func (s *ProductService) GetProduct(id int64) ([]byte,error) {
	var query = queryList.QueryGetProductByID
		rows, err := s.db.ExecuteQuery(query)
	if err != nil {
		return nil, err
	}

	return helper.RowsToJSON(*rows)

}
func (s *ProductService) GetAllProducts() ([]byte,error) {
	var query = queryList.QueryGetProductByID
	rows, err := s.db.ExecuteQuery(query)
	if err != nil {
		return nil, err
	}

	return helper.RowsToJSON(*rows)
}
func (s *ProductService) GetProductDetail(id int64) ([]byte,error) {
	var query = queryList.QueryGetDetailsByProductID
	rows, err := s.db.ExecuteQuery(query)
	if err != nil {
		return nil, err
	}

	return helper.RowsToJSON(*rows)
}

//INSERT ----

func (s *ProductService) InsertProduct(product *model.Product) bool {
	if product == nil {
		return false
	}
	var query = queryList.QueryInsertProduct
	rowsAffected, err := s.db.ExecuteNonQuery(query)
	if err != nil {
		fmt.Println("Erro ao inserir:", err)
		return false
	}

	return rowsAffected > 0
}

// UPDATE ----
func (s *ProductService) UpdateProduct(id int64, product *model.Product) bool {
	if id == 0 || product == nil {
		return false
	}
	var query = queryList.QueryUpdateProduct
	rowsAffected, err := s.db.ExecuteNonQuery(query)
	if err != nil {
		fmt.Println("Erro ao atualizar:", err)
		return false
	}

	return rowsAffected > 0
}

// DELETE ----
func (s *ProductService) DeleteProduct(id int64) bool {
	if id == 0 {
		return false
	}
	var query = queryList.QueryInsertProduct
	rowsAffected, err := s.db.ExecuteNonQuery(query)
	if err != nil {
		fmt.Println("Erro ao apagar:", err)
		return false
	}

	return rowsAffected > 0
}
