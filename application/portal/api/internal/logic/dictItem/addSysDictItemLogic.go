package dictItem

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysdictservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysDictItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysDictItemLogic {
	return &AddSysDictItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysDictItemLogic) AddSysDictItem(req *types.AddSysDictItemRequest) (string, error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok || account == "" {
		account = "system"
	}
	_, err := l.svcCtx.SysDictItemRpc.AddSysDictItem(l.ctx, &sysdictservice.AddSysDictItemReq{
		CreateBy:    account,
		UpdateBy:    account,
		DictCode:    req.DictCode,
		ItemText:    req.ItemText,
		ItemCode:    req.ItemCode,
		Description: req.Description,
		SortOrder:   req.SortOrder,
	})
	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("添加字典失败: 请求=%+v, 错误=%v", req, err)
		return "", err
	}

	return "添加数据字典成功", nil
}
