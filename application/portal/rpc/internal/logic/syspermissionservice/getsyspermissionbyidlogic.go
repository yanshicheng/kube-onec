package syspermissionservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysPermissionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysPermissionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysPermissionByIdLogic {
	return &GetSysPermissionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysPermissionByIdLogic) GetSysPermissionById(in *pb.GetSysPermissionByIdReq) (*pb.GetSysPermissionByIdResp, error) {
	resp, err := l.svcCtx.SysPermission.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("权限不存在: %v", err)
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("查询权限失败: %v", err)
		return nil, err
	}
	return &pb.GetSysPermissionByIdResp{
		Data: &pb.SysPermission{
			Id:         resp.Id,
			ParentId:   resp.ParentId,
			Name:       resp.Name,
			Uri:        resp.Uri,
			Action:     resp.Action,
			Level:      resp.Level,
			CreateTime: resp.CreateTime.Unix(), // 将 time.Time 转换为 Unix 时间戳
			UpdateTime: resp.UpdateTime.Unix(), // 将 time.Time 转换为 Unix 时间戳
		},
	}, nil
}
