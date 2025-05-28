package repositories

type Category struct {
	Id   int64  `json:"Id"`
	Name string `json:"Name"`
}

func NewCategory() *Category {
	return &Category{}
}
func GetCategory(id int64) *Category {
	return NewCategory()
}
