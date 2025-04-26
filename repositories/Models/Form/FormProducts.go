package formProducts
type FormProducts struct {
    Id                int64   `json:"id"`
    IdForm            int64   `json:"id_form"`
    IdProduct         int64   `json:"id_product"`
    IdImageAttachment int64   `json:"id_image_attachment"`
    Quantity          float64 `json:"quantity"`
    X                 float64 `json:"x"`
    Y                 float64 `json:"y"`
    Z                 float64 `json:"z"`
}