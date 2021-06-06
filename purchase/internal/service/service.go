package service

import "purchase/internal/dao"

type PurchaseService struct {
	pDao      *dao.Dao
}

func New() (s *PurchaseService) {
	return &PurchaseService{
		pDao: dao.NewDao(),
	}
}