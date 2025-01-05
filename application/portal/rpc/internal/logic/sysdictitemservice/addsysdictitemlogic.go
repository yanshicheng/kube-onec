package sysdictitemservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysDictItemLogic {
	return &AddSysDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------字典数据表-----------------------
func (l *AddSysDictItemLogic) AddSysDictItem(in *pb.AddSysDictItemReq) (*pb.AddSysDictItemResp, error) {

	// 检查数据字典是否存在
	_, err := l.svcCtx.SysDict.FindOneByDictCode(l.ctx, in.DictCode)
	if err != nil {
		l.Logger.Errorf("数据字典不存在: %v", err)
		return nil, code.DictNotExistErr
	}

	_, err = l.svcCtx.SysDictItem.Insert(l.ctx, &model.SysDictItem{
		CreateBy:    in.CreateBy,
		DictCode:    in.DictCode,
		ItemText:    in.ItemText,
		ItemCode:    in.ItemCode,
		Description: in.Description,
		SortOrder:   in.SortOrder,
		UpdateBy:    in.UpdateBy,
	})
	if err != nil {
		l.Logger.Errorf("添加字典数据失败: %v", err)
		return nil, errorx.DatabaseCreateErr
	}
	return &pb.AddSysDictItemResp{}, nil
}
