package models

type Books struct {
	Books []Book `json:"books"`
}
type Book struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}