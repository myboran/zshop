package goods

import (
	"net/http"

	"github.com/myboran/zshop/app/goods/cmd/api/internal/logic/goods"
	"github.com/myboran/zshop/app/goods/cmd/api/internal/svc"
	"github.com/myboran/zshop/app/goods/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GoodsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := goods.NewGoodsLogic(r.Context(), svcCtx)
		resp, err := l.Goods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
