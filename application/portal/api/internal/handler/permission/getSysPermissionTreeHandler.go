package permission

import (
	"net/http"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/logic/permission"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetSysPermissionTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := permission.NewGetSysPermissionTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetSysPermissionTree()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
