package service

import "purchase/internal/dao"

type Service struct {
	pDao      *dao.Dao
}

func New() (s *Service) {
	return &Service{
		pDao: dao.NewDao(),
	}
}