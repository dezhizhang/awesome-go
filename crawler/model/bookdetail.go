package model

type BookDetail struct {
	Author   string `json:"author"`
	Press    string `json:"press"`
	BookPage string `json:"bookPage"`
	Price    string `json:"price"`
	Score    string `json:"score"`
	Content  string `json:"content"`
}
