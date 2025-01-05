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

type GetSysOrganizationByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysOrganizationByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysOrganizationByIdLogic {
	return &GetSysOrganizationByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysOrganizationByIdLogic) GetSysOrganizationById(in *pb.GetSysOrganizationByIdReq) (*pb.GetSysOrganizationByIdResp, error) {
	resp, err := l.svcCtx.SysOrganization.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("根据ID: %d查询数据不存在", in.Id)
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("根据ID: %d查询数据失败: %v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	// 转换时间字段
	data := &pb.SysOrganization{}
	data.Id = resp.Id
	data.Name = resp.Name
	data.ParentId = resp.ParentId
	data.Level = resp.Level
	data.Description = resp.Description
	data.UpdateTime = resp.UpdateTime.Unix()
	data.CreateTime = resp.CreateTime.Unix()
	return &pb.GetSysOrganizationByIdResp{
		Data: data,
	}, nil
}
