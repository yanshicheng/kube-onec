package organization

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysorganizationservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysOrganizationByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysOrganizationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysOrganizationByIdLogic {
	return &GetSysOrganizationByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysOrganizationByIdLogic) GetSysOrganizationById(req *types.DefaultIdRequest) (resp *types.SysOrganization, err error) {
	res, err := l.svcCtx.SysOrganizationRpc.GetSysOrganizationById(l.ctx, &sysorganizationservice.GetSysOrganizationByIdReq{
		Id: req.Id,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("获取机构失败: %v", err)
		return nil, err
	}
	resp = &types.SysOrganization{
		Id:          res.Data.Id,
		Name:        res.Data.Name,
		ParentId:    res.Data.ParentId,
		Description: res.Data.Description,
		Level:       res.Data.Level,
		UpdatedAt:   res.Data.UpdatedAt,
		CreatedAt:   res.Data.CreatedAt,
	}

	return
}
