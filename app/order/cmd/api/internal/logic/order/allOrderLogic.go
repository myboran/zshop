package order

import (
	"context"
	"github.com/myboran/zshop/app/order/cmd/rpc/pb"

	"github.com/myboran/zshop/app/order/cmd/api/internal/svc"
	"github.com/myboran/zshop/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AllOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAllOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AllOrderLogic {
	return &AllOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AllOrderLogic) AllOrder() (resp *types.OrderResp, err error) {
	orders, err := l.svcCtx.OrderRpc.AllOrderInfo(l.ctx, &pb.OrderEmpty{})
	if err != nil {
		return nil, err
	}
	resp = &types.OrderResp{Orders: make([]types.Order, len(orders.Orders))}
	for i, v := range orders.Orders {
		resp.Orders[i].Id = v.Id
		resp.Orders[i].OrderSn = v.OrderSn
		resp.Orders[i].User = v.User
		resp.Orders[i].Status = v.Status
	}
	return
}
