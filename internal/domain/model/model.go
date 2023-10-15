package model

// Order model.
type Order struct {
	Item   int         `json:"item"`
	Packet map[int]int `json:"packets"`
}
