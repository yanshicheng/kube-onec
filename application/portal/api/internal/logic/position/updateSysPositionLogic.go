package position

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspositionservice"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysPositionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysPositionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysPositionLogic {
	return &UpdateSysPositionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysPositionLogic) UpdateSysPosition(req *types.UpdateSysPositionRequest) (resp string, err error) {
	res, err := l.svcCtx.SysPositionRpc.UpdateSysPosition(l.ctx, &syspositionservice.UpdateSysPositionReq{
		Id:   req.Id,
		Name: req.Name,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("更新职位失败: %v", err)
		return "", errorx.DatabaseUpdateErr
	}
	logx.WithContext(l.ctx).Infof("更新职位成功: %v", res)
	return "更新职位成功!", nil
}
