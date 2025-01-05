package sysorganizationservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysOrganizationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysOrganizationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysOrganizationLogic {
	return &AddSysOrganizationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------组织表-----------------------
func (l *AddSysOrganizationLogic) AddSysOrganization(in *pb.AddSysOrganizationReq) (*pb.AddSysOrganizationResp, error) {
	// 添加组织表
	var org model.SysOrganization
	org.Name = in.Name
	org.Description = in.Description
	org.ParentId = in.ParentId

	if in.ParentId != 0 {
		// 如果不等于0 说明有父节点
		// 查询父节点
		parent, err := l.svcCtx.SysOrganization.FindOne(l.ctx, in.ParentId)
		if err != nil {
			if errors.Is(err, model.ErrNotFound) {
				l.Logger.Infof("父节点未找到: 父节点ID=%d", in.ParentId)
				return nil, errorx.DatabaseNotFound
			}
			l.Logger.Errorf("查询父节点信息失败: 父节点ID=%d, 错误=%v", in.ParentId, err)
			return nil, err
		}
		org.Level = parent.Level + 1
		if org.Level > 5 {
			l.Logger.Errorf("组织层级不能超过5级: 父节点ID=%d, 父节点层级=%d", in.ParentId, parent.Level)
			return nil, code.CreateOrganizationErr
		}
	} else {
		org.Level = 1
	}
	_, err := l.svcCtx.SysOrganization.Insert(l.ctx, &org)
	if err != nil {
		l.Logger.Errorf("添加组织失败: %v", err)
		return nil, err
	}
	return &pb.AddSysOrganizationResp{}, nil
}
