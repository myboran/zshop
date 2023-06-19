package logic

import (
	"context"

	"github.com/myboran/zshop/app/inventory/cmd/rpc/internal/svc"
	"github.com/myboran/zshop/app/inventory/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InventoryDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInventoryDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InventoryDetailLogic {
	return &InventoryDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InventoryDetailLogic) InventoryDetail(in *pb.InventoryDetailReq) (*pb.InventoryDetailResp, error) {
	inventory, err := l.svcCtx.InventoryModel.FindOneByGoods(l.ctx, in.Goods)
	if err != nil {
		return nil, err
	}
	return &pb.InventoryDetailResp{
		Id:     inventory.Id,
		Goods:  inventory.Goods,
		Stocks: inventory.Stocks,
		Freeze: inventory.Freeze,
	}, nil
}
