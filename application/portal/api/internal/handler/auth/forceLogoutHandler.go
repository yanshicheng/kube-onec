package auth

import (
	"net/http"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/logic/auth"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ForceLogoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := auth.NewForceLogoutLogic(r.Context(), svcCtx)
		resp, err := l.ForceLogout()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
