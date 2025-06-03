package repositories

type ProductRepositoy interface {
	GetProduct() string
	InsertProduct() bool
	RemoveProduct() bool
	UpdateProduct() bool
	Test() string
}
type ImageRepositoy interface {
	GetImage() string
}

type FormRepository interface {
}
