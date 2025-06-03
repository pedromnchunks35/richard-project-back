package repositories

type Product struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	IdCategory int64  `json:"id_category"`
	IdImage    int64  `json:"id_image"`
}
