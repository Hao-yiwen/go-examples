// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/logic"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/svc"
	"github.com/haoyiwen/go-examples/go-zero/user-api/internal/types"
)

func ListUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewListUserLogic(r.Context(), svcCtx)
		resp, err := l.ListUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
