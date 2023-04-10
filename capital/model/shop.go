package model

type Shop struct {
	Id    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
	Desc  string  `json:"desc"`
	Sales int     `json:"sales"`
	Stock int     `json:"stock"`
}
