package model

// Order model
type Order struct {
	Item   uint
	Packet map[uint]uint
}
