package sysdictservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysDictLogic {
	return &DelSysDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSysDictLogic) DelSysDict(in *pb.DelSysDictReq) (*pb.DelSysDictResp, error) {
	// 删除数据字典首先查询是否有对应的数据
	dictItems, err := l.svcCtx.SysDictItem.SearchNoPage(l.ctx, "", true, "dict_id = ?", in.Id)
	if err != nil && !errors.Is(err, model.ErrNotFound) {

		l.Logger.Errorf("查询数据字典失败: %v", err)
		return nil, code.FindDictItemsErr
	}
	if len(dictItems) > 0 {
		l.Logger.Errorf("数据字典下有数据，无法删除: %v", err)
		return nil, code.DictHasItemsErr
	}

	// 定义修改 delete_time 实践，和 修改 updatedBy 人 sql
	sqlStr := "UPDATE {table} set `delete_time` = NOW(), `update_by` = ? where `id` = ?"
	_, err = l.svcCtx.SysDict.ExecSql(l.ctx, in.Id, sqlStr, in.UpdatedBy, in.Id)
	if err != nil {
		l.Logger.Errorf("删除字典失败: %v", err)
		return nil, errorx.DatabaseDeleteErr
	}
	return &pb.DelSysDictResp{}, nil
}
