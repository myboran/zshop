package logic

import (
	"context"

	"github.com/myboran/zshop/app/order/cmd/rpc/internal/svc"
	"github.com/myboran/zshop/app/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOrderLogic {
	return &UserOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOrderLogic) UserOrder(in *pb.UserOrderReq) (*pb.UserOrderResp, error) {
	uOrders, err := l.svcCtx.OrderInfoModel.FindOrderByUser(l.ctx, in.User)
	if err != nil {
		return nil, err
	}
	resp := &pb.UserOrderResp{Orders: make([]*pb.OrderInfo, len(uOrders))}
	for i, v := range uOrders {
		resp.Orders[i] = new(pb.OrderInfo)
		resp.Orders[i].Id = v.Id
		resp.Orders[i].OrderSn = v.OrderSn
		resp.Orders[i].User = v.User
		resp.Orders[i].Status = v.Statue
	}
	return resp, nil
}
