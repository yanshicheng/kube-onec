package position

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspositionservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysPositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysPositionLogic {
	return &AddSysPositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysPositionLogic) AddSysPosition(req *types.AddSysPositionRequest) (resp string, err error) {
	res, err := l.svcCtx.SysPositionRpc.AddSysPosition(l.ctx, &syspositionservice.AddSysPositionReq{
		Name: req.Name,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("添加职位失败: %v", err)
		return "", err
	}
	logx.WithContext(l.ctx).Infof("添加职位成功: %v", res)
	return "添加职位成功!", nil
}
