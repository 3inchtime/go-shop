package dao

import (
	"fmt"
	"purchase/internal/model"
	"testing"
)

var d = NewDao()

func TestDao_QueryPurchaseByUserID(t *testing.T) {
	var (
		UserID = 1
		Start  = 0
		Num    = 20
	)
	pList, err := d.QueryPurchaseByUserID(UserID, Start, Num)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%+v\n", pList)
}

func TestDao_CreatePurchase(t *testing.T) {
	p := &model.Purchase{
		UserID: 1,
		TotalPrice: 1000,
		PurchaseDetail: "adflkadfka",
	}
	rowID, err := d.CreatePurchase(p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Row ID: %d", rowID)
}

func TestDao_PaidPurchase(t *testing.T) {
	var pID = 1
	rowID, err := d.PaidPurchase(pID)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Row ID: %d", rowID)
}
