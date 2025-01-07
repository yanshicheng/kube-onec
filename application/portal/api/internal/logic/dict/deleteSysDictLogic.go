package dict

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteSysDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteSysDictLogic {
	return &DeleteSysDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteSysDictLogic) DeleteSysDict(req *types.DefaultIdRequest) (resp string, err error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok {
		account = "system"
	}
	_, err = l.svcCtx.SysDictRpc.DelSysDict(l.ctx, &pb.DelSysDictReq{
		Id:       req.Id,
		UpdatedBy: account,
	})
	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("删除字典失败: 请求=%+v, 错误=%v", req, err)
		return
	}
	return "数据字典删除成功!", nil
}
