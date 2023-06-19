package goods

import (
	"context"
	"github.com/myboran/zshop/app/goods/cmd/rpc/goods"
	"github.com/myboran/zshop/app/inventory/cmd/rpc/inventory"

	"github.com/myboran/zshop/app/goods/cmd/api/internal/svc"
	"github.com/myboran/zshop/app/goods/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodsLogic {
	return &GoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GoodsLogic) Goods(req *types.GoodsReq) (resp *types.GoodsResp, err error) {
	goodsInfo, err := l.svcCtx.GoodsRpc.GoodsDetail(l.ctx, &goods.GoodDetailReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	iv, err := l.svcCtx.InventoryRpc.InventoryDetail(l.ctx, &inventory.InventoryDetailReq{
		Goods: goodsInfo.GetId(),
	})
	if err != nil {
		return nil, err
	}
	return &types.GoodsResp{
		Id:     goodsInfo.Id,
		Name:   goodsInfo.Name,
		Price:  float64(goodsInfo.Price),
		Stocks: iv.Stocks,
	}, nil
}
