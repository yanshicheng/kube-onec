package errorx

import (
	"github.com/yanshicheng/kube-onec/common/handler/errorx/types"
	"net/http"
)

func ErrHandler(err error) (int, any) {
	code := CodeFromError(err)
	return http.StatusOK, types.Status{
		Code:    int32(code.Code()),
		Message: code.Message(),
	}
}
