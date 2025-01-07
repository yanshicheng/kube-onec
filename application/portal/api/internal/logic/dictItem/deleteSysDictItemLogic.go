package dictItem

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDictItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDictItemLogic {
	return &DeleteSysDictItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDictItemLogic) DeleteSysDictItem(req *types.DefaultIdRequest) (resp string, err error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok || account == "" {
		account = "system"
	}
	_, err = l.svcCtx.SysDictItemRpc.DelSysDictItem(l.ctx, &pb.DelSysDictItemReq{
		Id:       req.Id,
		UpdatedBy: account,
	})

	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("删除字典失败: 请求=%+v, 错误=%v", req, err)
		return "", err
	}
	return "删除成功", nil
}
