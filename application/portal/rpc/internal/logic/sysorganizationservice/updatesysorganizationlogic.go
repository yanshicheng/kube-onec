package sysorganizationservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysOrganizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateSysOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysOrganizationLogic {
	return &UpdateSysOrganizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateSysOrganizationLogic) UpdateSysOrganization(in *pb.UpdateSysOrganizationReq) (*pb.UpdateSysOrganizationResp, error) {
	// 先查询
	resp, err := l.svcCtx.SysOrganization.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("职位未找到: 职位ID=%d", in.Id)
			return nil, errorx.DatabaseQueryErr
		}
		l.Logger.Errorf("查询职位信息失败: 职位ID=%d, 错误=%v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	if in.Name != "" {
		resp.Name = in.Name
	}
	if in.Description != "" {
		resp.Description = in.Description
	}
	if err := l.svcCtx.SysOrganization.Update(l.ctx, resp); err != nil {
		l.Logger.Errorf("更新职位信息失败: 职位ID=%d, 错误=%v", in.Id, err)
		return nil, errorx.DatabaseUpdateErr
	}
	return &pb.UpdateSysOrganizationResp{}, nil
}
