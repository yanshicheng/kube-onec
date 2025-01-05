package organization

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrganizationTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrganizationTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrganizationTreeLogic {
	return &GetOrganizationTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrganizationTreeLogic) GetOrganizationTree(req *types.GetOrganizationTreeRequest) (resp *types.GetOrganizationTreeResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
