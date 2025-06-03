package repositories

type Product struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	IdCategory int64  `json:"id_category"`
	IdDetail   int64  `json:"id_detail"`
	IdImage    int64  `json:"id_image"`
}
