package models



type Content struct {
	Title  string `json:"title"`
	Introduction string `json:"introduction"`
	Examples     string `json:"examples"`
}

type Lesson struct {
	ID int `json:"id"`
	Level string `json:"level"`
	Lesson int `json:"lesson"`
	Content Content `json:"content"`
}