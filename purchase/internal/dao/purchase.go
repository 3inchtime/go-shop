package dao

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"purchase/internal/model"
	"time"
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

func (d *Dao) CreatePurchase (p *model.Purchase) (rows int, err error) {
	createTime := time.Now().Unix()
	insertSQL, err := d.DB.Prepare("INSERT INTO purchase (user_id, total_price, purchase_detail, create_time) VALUES (?, ?, ?, ?)")
	if err != nil {
		logrus.Errorf("Prepare Insert SQL Error: %s", err.Error())
	}

	res, err := insertSQL.Exec(p.UserID, p.TotalPrice, p.PurchaseDetail, createTime)
	if err != nil {
		logrus.Errorf("Insert Video Info SQL Error: %s", err.Error())
		return 0, err
	}
	defer insertSQL.Close()
	id, _ := res.LastInsertId()
	return int(id), nil
}

func (d *Dao) PaidPurchase (pID int) (rows int64, err error) {
	paidTime := time.Now().Unix()
	updateSQL, err := d.DB.Prepare("UPDATE purchase SET paid_time=?, paid_status=? WHERE id=?")
	if err != nil {
		logrus.Errorf("Prepare Update SQL Error: %s", err.Error())

	}
	res, err := updateSQL.Exec(paidTime, 1, pID)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	return res.RowsAffected()
}

func (d *Dao) QueryPurchaseDetailByID (id int) (p *model.Purchase, err error) {
	querySQL, err := d.DB.Prepare("SELECT * FROM purchase WHERE id = ?")
	if err != nil {
		logrus.Errorf("Prepare Select SQL Error: %s", err.Error())
		return nil, err
	}
	p = new(model.Purchase)
	err = querySQL.QueryRow(id).Scan(&p.ID, &p.UserID, &p.TotalPrice, &p.PurchaseDetail, &p.PaidTime, &p.PaidStatus, &p.Status, p.CreateTime)
	if err != nil && err != sql.ErrNoRows {
		logrus.Errorf("Query PurcaseDetail Error: %s", err.Error())
	}
	return p, nil
}

