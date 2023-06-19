package logic

import (
	"context"

	"github.com/myboran/zshop/app/order/cmd/rpc/internal/svc"
	"github.com/myboran/zshop/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllOrderInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAllOrderInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllOrderInfoLogic {
	return &AllOrderInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AllOrderInfoLogic) AllOrderInfo(in *pb.OrderEmpty) (*pb.AllOrderResp, error) {
	allOrders, err := l.svcCtx.OrderInfoModel.FindOrders(l.ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.AllOrderResp{Orders: make([]*pb.OrderInfo, len(allOrders))}
	for i, v := range allOrders {
		resp.Orders[i] = new(pb.OrderInfo)
		resp.Orders[i].Id = v.Id
		resp.Orders[i].OrderSn = v.OrderSn
		resp.Orders[i].User = v.User
		resp.Orders[i].Status = v.Statue
	}
	return resp, nil
}
