package dict

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysdictservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysDictLogic {
	return &AddSysDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSysDictLogic) AddSysDict(req *types.AddSysDictRequest) (resp string, err error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok || account == "" {
		account = "system"
	}
	_, err = l.svcCtx.SysDictRpc.AddSysDict(l.ctx, &sysdictservice.AddSysDictReq{
		CreatedBy:    account,
		UpdatedBy:    account,
		DictCode:    req.DictCode,
		DictName:    req.DictName,
		Description: req.Description,
	})
	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("添加字典失败: 请求=%+v, 错误=%v", req, err)
		return
	}
	return "数据字典添加成功!", nil
}
