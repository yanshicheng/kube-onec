package user

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/logic/user"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/common/verify"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GlobalUpdateSysUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateGloabSysUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		// 设置默认值
		defaults.SetDefaults(&req)
		// validator验证
		if err := svcCtx.Validator.Validate.StructCtx(r.Context(), &req); err != nil {
			strErr := verify.RemoveTopSaStr(err.(validator.ValidationErrors), svcCtx.Validator.Translator)
			httpx.ErrorCtx(r.Context(), w, errorx.New(40020, strErr))
			return
		}
		l := user.NewGlobalUpdateSysUserLogic(r.Context(), svcCtx)
		resp, err := l.GlobalUpdateSysUser(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
