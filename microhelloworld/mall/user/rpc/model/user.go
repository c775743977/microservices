package model

type User struct {
	Id     string  `db:"number"`
	Name   string `db:"name"`
	Password string `db:"password"`
	Gender string `db:"gender"`
}

func(this *User) TableName() string {
	return "user"
}