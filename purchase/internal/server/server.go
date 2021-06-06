package server

import (
	"context"
	"purchase/internal/service"
	pb "purchase/proto"
)

type PurchaseServer struct {
	PurchaseService  *service.PurchaseService
}

// QueryPurchaseDetail(context.Context, *QueryPurchaseDetailRequest, *PurchaseDetail) error
func (s *PurchaseServer) QueryPurchaseDetail(ctx context.Context, q *pb.QueryPurchaseDetailRequest, r *pb.PurchaseDetail) (error) {
	detail, err := s.PurchaseService.QueryPurchaseDetail(int(q.Id))
	if err != nil {
		return err
	}
	response := &pb.PurchaseDetail{}
	response.Id = int32(detail.ID)
	response.UserId = int32(detail.UserID)
	response.TotalPrice = int32(detail.TotalPrice)
	r = response

	return nil
}

// QueryUserPurchase(context.Context, *QueryPurchaseRequest, *QueryPurchaseResponse) error
func (s *PurchaseServer) QueryUserPurchase(ctx context.Context, q *pb.QueryPurchaseRequest, r *pb.QueryPurchaseResponse) error {
	pList, err := s.PurchaseService.QueryUserPurchase(int(q.UserId))
	if err != nil {
		return err
	}
	response := &pb.QueryPurchaseResponse{}
	for _, p := range pList{
		tmp := &pb.PurchaseDetail{}
		tmp.Id = int32(p.ID)
		tmp.UserId = int32(p.UserID)
		tmp.TotalPrice = int32(p.TotalPrice)
		response.Purchases = append(response.Purchases, tmp)
	}
	r = response
	return nil
}