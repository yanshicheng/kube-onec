package sysdictitemservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckDictItemCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckDictItemCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckDictItemCodeLogic {
	return &CheckDictItemCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckDictItemCodeLogic) CheckDictItemCode(in *pb.CheckDictItemCodeReq) (*pb.CheckDictItemCodeResp, error) {
	_, err := l.svcCtx.SysDictItem.FindOneByDictCodeItemCode(l.ctx, in.DictCode, in.ItemCode)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			l.Logger.Infof("字典项未找到: 字典编码=%s, 字典项编码=%s", in.DictCode, in.ItemCode)
			return nil, err
		}
		l.Logger.Errorf("查询字典项信息失败: 字典编码=%s, 字典项编码=%s, 错误=%v", in.DictCode, in.ItemCode, err)
		return nil, err
	}
	l.Logger.Infof("字典项已存在: 字典编码=%s, 字典项编码=%s", in.DictCode, in.ItemCode)
	return nil, nil
}
