package order

import (
	"net/http"

	"github.com/myboran/zshop/app/order/cmd/api/internal/logic/order"
	"github.com/myboran/zshop/app/order/cmd/api/internal/svc"
	"github.com/myboran/zshop/app/order/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderGoodsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := order.NewOrderGoodsLogic(r.Context(), svcCtx)
		resp, err := l.OrderGoods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
