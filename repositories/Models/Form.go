package repositories

import "richard-project-back/dtos"

type Form struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Company        string `json:"company"`
	NIF            int    `json:"nif"`
	PhoneNumber    string `json:"phone_number"`
	Email          string `json:"email"`
	FiscalAddress  string `json:"fiscal_address"`
	Address        string `json:"address"`
	WithDelivery   bool   `json:"with_delivery"`
	Observations   string `json:"observations"`
	IdIntervention int64  `json:"id_intervention"`
	IdFormProducts int64  `json:"id_form_products"`
}

func NewForm(dto *dtos.Form) *Form {
	return &Form{
		Id:             dto.Id,
		Name:           dto.Name,
		Company:        dto.Company,
		NIF:            dto.NIF,
		PhoneNumber:    dto.PhoneNumber,
		Email:          dto.Email,
		FiscalAddress:  dto.FiscalAddress,
		Address:        dto.Address,
		WithDelivery:   dto.WithDelivery,
		Observations:   dto.Observations,
		IdIntervention: dto.Intervetion.Id,
	}
}
