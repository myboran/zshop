package goods

import (
	"context"
	"github.com/myboran/zshop/app/goods/cmd/rpc/goods"

	"github.com/myboran/zshop/app/goods/cmd/api/internal/svc"
	"github.com/myboran/zshop/app/goods/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List(req *types.ListReq) (resp *types.ListResp, err error) {
	gs, err := l.svcCtx.GoodsRpc.GoodsList(l.ctx, &goods.GoodsListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	resp = &types.ListResp{Goods: make([]types.Goods, len(gs.GetGoods()))}
	for i, v := range gs.Goods {
		resp.Goods[i].Id = v.Id
		resp.Goods[i].Price = float64(v.Price)
		resp.Goods[i].Name = v.Name
	}
	return resp, nil
}
