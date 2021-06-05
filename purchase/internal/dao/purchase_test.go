package dao

import (
	"fmt"
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
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", pList)
}
