package sysdictitemservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelSysDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelSysDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelSysDictItemLogic {
	return &DelSysDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelSysDictItemLogic) DelSysDictItem(in *pb.DelSysDictItemReq) (*pb.DelSysDictItemResp, error) {
	// 定义修改 delete_time 实践，和 修改 updatedBy 人 sql
	sqlStr := "UPDATE {table} set `delete_time` = NOW(), `update_by` = ? where `id` = ?"
	_, err := l.svcCtx.SysDictItem.ExecSql(l.ctx, in.Id, sqlStr, in.UpdatedBy, in.Id)
	if err != nil {
		l.Logger.Errorf("删除字典数据失败: %v", err)
		return nil, errorx.DatabaseDeleteErr
	}

	return &pb.DelSysDictItemResp{}, nil
}
