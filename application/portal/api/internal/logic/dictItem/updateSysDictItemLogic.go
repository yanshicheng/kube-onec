package dictItem

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysDictItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysDictItemLogic {
	return &UpdateSysDictItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysDictItemLogic) UpdateSysDictItem(req *types.UpdateSysDictItemRequest) (string, error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok || account == "" {
		return "", nil
	}

	_, err := l.svcCtx.SysDictItemRpc.UpdateSysDictItem(l.ctx, &pb.UpdateSysDictItemReq{
		Id:          req.Id,
		Description: req.Description,
		ItemText:    req.ItemText,
		SortOrder:   req.SortOrder,
		UpdateBy:    account,
	})

	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("更新字典失败: 请求=%+v, 错误=%v", req, err)
		return "", err
	}
	return "更新字典成功", nil
}
