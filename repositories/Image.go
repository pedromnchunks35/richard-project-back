package repositories

type Image struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Base64 string `json:"base64"`
}

func NewImage() Image {
	return Image{}
}
