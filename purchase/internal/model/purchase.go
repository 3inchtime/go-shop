package model

type Purchase struct {
	ID int64
	UserID int
	TotalPrice int
	PurchaseDetail string
	PaidTime int
	PaidStatus int
	Status int
	CreateTime int
}