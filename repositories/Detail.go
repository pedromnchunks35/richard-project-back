package repositories

type Detail struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
	IdAgent     int64  `json:"id_agent"`
}

func NewDetail() *Detail {
	return &Detail{}
}

func GetDetail(id int64) *Detail {
	return NewDetail()
}
