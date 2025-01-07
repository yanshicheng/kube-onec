package position

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/syspositionservice"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysPositionByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSysPositionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysPositionByIdLogic {
	return &GetSysPositionByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSysPositionByIdLogic) GetSysPositionById(req *types.DefaultIdRequest) (resp *types.SysPosition, err error) {
	res, err := l.svcCtx.SysPositionRpc.GetSysPositionById(l.ctx, &syspositionservice.GetSysPositionByIdReq{
		Id: req.Id,
	})
	if err != nil {
		logx.Errorf("查询职位信息失败: 职位ID=%d, 错误=%v", req.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	resp = &types.SysPosition{
		CreatedAt: res.Data.CreatedAt,
		Id:        res.Data.Id,
		Name:      res.Data.Name,
		UpdatedAt: res.Data.UpdatedAt,
	}
	return
}
