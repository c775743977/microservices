package model

type User struct {
	Id     int64  `db:"id"`
	Name   string `db:"name"`
	Password string `db:"password"`
	Gender string `db:"gender"`
}

func(this *User) TableName() string {
	return "user"
}