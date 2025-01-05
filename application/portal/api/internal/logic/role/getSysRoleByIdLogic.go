package role

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysroleservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysRoleByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysRoleByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysRoleByIdLogic {
	return &GetSysRoleByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysRoleByIdLogic) GetSysRoleById(req *types.DefaultIdRequest) (resp *types.GetSysRoleByIdResponse, err error) {
	resp = &types.GetSysRoleByIdResponse{}
	res, err := l.svcCtx.SysRoleRpc.GetSysRoleById(l.ctx, &sysroleservice.GetSysRoleByIdReq{
		Id: req.Id,
	})
	logx.WithContext(l.ctx).Infof("获取角色: %v", res)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("获取角色失败: %v", err)
		return nil, err
	}
	err = copier.Copy(&resp, res.Data)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("数据拷贝失败: 错误=%v", err)
		return nil, err
	}
	logx.WithContext(l.ctx).Infof("获取角色成功: %v", resp)
	return resp, nil
}
