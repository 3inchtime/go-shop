package service

import (
	"github.com/sirupsen/logrus"
	"purchase/internal/model"
)

func (s *PurchaseService) CreatePurchase(p *model.Purchase) (id int, err error) {
	id, err = s.pDao.CreatePurchase(p)
	if err != nil {
		logrus.Errorf("Create Purchase Error: %s", err.Error())
		return 0, err
	}
	return id, nil
}

func (s *PurchaseService) QueryPurchaseDetail(id int) (p *model.Purchase, err error) {
	p, err = s.pDao.QueryPurchaseDetailByID(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *PurchaseService) QueryUserPurchase(userID int) (pList []*model.Purchase, err error) {
	pList, err = s.pDao.QueryPurchaseByUserID(userID, 0, 10)
	if err != nil {
		return nil, err
	}
	return pList, nil
}