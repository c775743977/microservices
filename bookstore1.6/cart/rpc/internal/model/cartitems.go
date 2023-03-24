package model

type CartItem struct {
	// ID int
	Book *Book `gorm:"-"`
	Amount float64
	Num int32
	CartID string `gorm:"column:cart_id"`
	BookID string `gorm:"column:book_id"`
	IsThis bool `gorm:"-"`
}