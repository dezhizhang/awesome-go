package model

type BookDetail struct {
	Name     string `json:"name"`
	Author   string `json:"author"`
	Press    string `json:"press"`
	BookPage string `json:"bookPage"`
	Price    string `json:"price"`
	Score    string `json:"score"`
	Content  string `json:"content"`
}

func (b BookDetail) String() string {
	return "书名" + b.Name + "作者:" + b.Author + "出版社：" + b.Press + "页数：" + b.BookPage + "定价：" + b.Price +
		"评分:" + b.Score + "简介：" + b.Content
}
