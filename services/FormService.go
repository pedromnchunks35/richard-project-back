package services

import (
	"fmt"
	sql "richard-project-back/configs/connector/postgresql"
	"richard-project-back/dtos"
	"richard-project-back/helper"
	queryList "richard-project-back/repositories/Queries"
	iService "richard-project-back/services/ServicesInterface"
)

type FormService struct {
	db *sql.PostgresDbVault
}

func NewFormService(db *sql.PostgresDbVault) iService.IFormService {
	return &FormService{db: db}
}

func (s *FormService) TesteForm() string {
	return "GOOD"
}

func (s *FormService) CreateForm(form *dtos.Form) bool {
	if form == nil {
		return false
	}
	query := queryList.QueryInsertForm
	rowsAffected, err := s.db.ExecuteNonQuery(query)
	if err != nil {
		fmt.Println("Erro ao inserir form:", err)
		return false
	}

	return rowsAffected > 0
}

func (s *FormService) UpdateForm(id int64, form *dtos.Form) bool {
	if id == 0 || form == nil {
		return false
	}
	query := queryList.QueryUpdateForm
	rowsAffected, err := s.db.ExecuteNonQuery(query)
	if err != nil {
		fmt.Println("Erro ao inserir form:", err)
		return false
	}

	return rowsAffected > 0
}

func (s *FormService) DeleteForm(id int64) bool {
	if id == 0 {
		return false
	}
	query := queryList.QueryDeleteForm
	rowsAffected, err := s.db.ExecuteNonQuery(query)
	if err != nil {
		fmt.Println("Erro ao inserir form:", err)
		return false
	}

	return rowsAffected > 0
}

func (s *FormService) GetForm(id int64) ([]byte, error) {
	query := queryList.QueryGetFormWithProducts
	rows, err := s.db.ExecuteQuery(query)
	if err != nil {
		return nil, err
	}

	return helper.RowsToJSON(*rows)

}
