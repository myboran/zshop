package order

import (
	"net/http"

	"github.com/myboran/zshop/app/order/cmd/api/internal/logic/order"
	"github.com/myboran/zshop/app/order/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func AllOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := order.NewAllOrderLogic(r.Context(), svcCtx)
		resp, err := l.AllOrder()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
