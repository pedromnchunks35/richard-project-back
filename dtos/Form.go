package dtos

type Form struct {
	Id            int64        `json:"id"`
	Name          string       `json:"name"`
	Company       string       `json:"company"`
	NIF           int          `json:"nif"`
	PhoneNumber   string       `json:"phone_number"`
	Email         string       `json:"email"`
	FiscalAddress string       `json:"fiscal_address"`
	Address       string       `json:"address"`
	WithDelivery  bool         `json:"with_delivery"`
	Observations  string       `json:"observations"`
	IdIntervetion Intervetion  `json:"intervention"`
	FormProducts  FormProducts `json:"form_products"`
}

type Intervetion struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type FormProducts struct {
	Id              int64   `json:"id"`
	Product         Product `json:"product"`
	ImageAttachment Image   `json:"image_attachment"`
	Quantity        float64 `json:"quantity"`
	X               float64 `json:"x"`
	Y               float64 `json:"y"`
	Z               float64 `json:"z"`
}
