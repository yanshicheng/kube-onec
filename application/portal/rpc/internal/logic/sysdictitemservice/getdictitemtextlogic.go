package sysdictitemservicelogic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDictItemTextLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDictItemTextLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDictItemTextLogic {
	return &GetDictItemTextLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDictItemTextLogic) GetDictItemText(in *pb.GetDictItemNameReq) (*pb.GetDictItemTextResp, error) {
	res, err := l.svcCtx.SysDictItem.FindOneByDictCodeItemCode(l.ctx, in.DictCode, in.ItemCode)
	if err != nil {
		if errors.Is(err, sqlc.ErrNotFound) {
			l.Logger.Infof("字典项未找到: 字典编码=%s, 字典项编码=%s", in.DictCode, in.ItemCode)
			return nil, err
		}
		l.Logger.Errorf("查询字典项信息失败: 字典编码=%s, 字典项编码=%s, 错误=%v", in.DictCode, in.ItemCode, err)
		return nil, err
	}
	l.Logger.Infof("字典项已找到: 字典编码=%s, 字典项编码=%s, 字典项名称=%s", in.DictCode, in.ItemCode, res.ItemText)
	return &pb.GetDictItemTextResp{
		ItemText: res.ItemText,
	}, nil
}
