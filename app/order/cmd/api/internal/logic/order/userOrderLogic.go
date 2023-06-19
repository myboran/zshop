package order

import (
	"context"
	"github.com/myboran/zshop/app/order/cmd/rpc/order"

	"github.com/myboran/zshop/app/order/cmd/api/internal/svc"
	"github.com/myboran/zshop/app/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOrderLogic {
	return &UserOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserOrderLogic) UserOrder(req *types.UserOrderReq) (resp *types.OrderResp, err error) {
	orders, err := l.svcCtx.OrderRpc.UserOrder(l.ctx, &order.UserOrderReq{User: req.UserId})
	if err != nil {
		return nil, err
	}
	resp = &types.OrderResp{Orders: make([]types.Order, len(orders.Orders))}
	for i, v := range orders.Orders {
		resp.Orders[i].Id = v.Id
		resp.Orders[i].OrderSn = v.OrderSn
		resp.Orders[i].Status = v.Status
		resp.Orders[i].User = v.User
	}
	return resp, nil
}
