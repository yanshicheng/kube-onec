package position

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysPositionByUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysPositionByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysPositionByUserLogic {
	return &SearchSysPositionByUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSysPositionByUserLogic) SearchSysPositionByUser(req *types.DefaultIdRequest) (resp []types.PositionSysUser, err error) {
	// todo: add your logic here and delete this line

	return
}
