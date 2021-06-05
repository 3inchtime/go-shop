package dao

import (
	"github.com/sirupsen/logrus"
	"purchase/internal/model"
)

func (d *Dao) QueryPurchaseByUserID(UserID, Start, Num int) ([]*model.Purchase, error) {
	querySQL, err := d.DB.Prepare("SELECT * FROM purchase WHERE user_id = ? LIMIT ?,?")
	if err != nil {
		logrus.Errorf("Prepare Select SQL Error: %s", err.Error())
		return nil, err
	}

	defer querySQL.Close()
	rows, err := querySQL.Query(UserID, Start, Num)
	pList := make([]*model.Purchase, 0)
	for rows.Next() {
		var p = &model.Purchase{}
		if err = rows.Scan(&p.ID, &p.UserID, &p.TotalPrice, &p.PurchaseDetail, &p.PaidStatus, &p.PaidTime, &p.Status, &p.CreateTime); err != nil {
			logrus.Errorf("rows.Scan error(%v)", err)
			return nil, err
		}
		pList = append(pList, p)
	}
	err = rows.Err()
	if err != nil {
		logrus.Errorf("Query Purchase Error: %s", err.Error())
	}
	return pList, nil
}
