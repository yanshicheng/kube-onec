package position

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspositionservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysPositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelSysPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysPositionLogic {
	return &DelSysPositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelSysPositionLogic) DelSysPosition(req *types.DefaultIdRequest) (resp string, err error) {
	res, err := l.svcCtx.SysPositionRpc.DelSysPosition(l.ctx, &syspositionservice.DelSysPositionReq{
		Id: req.Id,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("删除职位失败: %v", err)
		return "", err
	}
	logx.WithContext(l.ctx).Infof("删除职位成功: %v", res)
	return "删除职位成功!", nil
}
