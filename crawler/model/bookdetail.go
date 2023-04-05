package model

type BookDetail struct {
	Author   string `json:"author"`
	Press    string `json:"press"`
	BookPage string `json:"bookPage"`
	Price    string `json:"price"`
	Score    string `json:"score"`
	Content  string `json:"content"`
}

func (b BookDetail) String() string {
	return "Author:" + b.Author + "出版社：" + b.Press + "页数：" + b.BookPage + "定价：" + b.Price + "\n简介：" + b.Content
}
