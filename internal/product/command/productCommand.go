package command

type CreateProductCommand struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Price    int32  `json:"price"`
}