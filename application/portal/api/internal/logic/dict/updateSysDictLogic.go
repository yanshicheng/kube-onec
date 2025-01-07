package dict

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateSysDictLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateSysDictLogic {
	return &UpdateSysDictLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateSysDictLogic) UpdateSysDict(req *types.UpdateSysDictRequest) (resp string, err error) {
	account, ok := l.ctx.Value("account").(string)
	if !ok {
		return "", errorx.Unauthorized
	}
	_, err = l.svcCtx.SysDictRpc.UpdateSysDict(l.ctx, &pb.UpdateSysDictReq{
		Id:          req.Id,
		DictName:    req.DictName,
		Description: req.Description,
		UpdatedBy:    account,
	})
	if err != nil {
		l.Logger.Errorf("数据字典更新失败: %v", err)
		return "", errorx.DatabaseUpdateErr
	}
	return "更新成功", nil
}
