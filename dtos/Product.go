package dtos

type Product struct {
	Id       uint     `json:"id"`
	Name     string   `json:"name"`
	Category Category `json:"category"`
	Image    Image    `json:"image"`
	Detatil  Detail   `json:"detail"`
}

type Category struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type Image struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Base64 string `json:"base64"`
}

type Detail struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Agent       Agent  `json:"agent"`
}

type Agent struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
