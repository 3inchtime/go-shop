package model

type User struct {
	ID int

	UserName   string
	Email      string
	Telephone  string
	Age        int
	CreateTime int

	HashPassword string
}
