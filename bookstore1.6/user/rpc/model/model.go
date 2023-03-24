package model

type User struct {
	ID int64
	Name string
	Password string
	Email string
	Privilege string
}

type Cart struct {
	ID int64
	Num int64
	Amount float64
	UserID int64
}