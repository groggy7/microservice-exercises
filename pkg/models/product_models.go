package models

type Product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Type  string `json:"type"`
	Price int    `json:"price"`
}
