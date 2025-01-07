package sysdictservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSysDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSysDictLogic {
	return &AddSysDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------字典表-----------------------
func (l *AddSysDictLogic) AddSysDict(in *pb.AddSysDictReq) (*pb.AddSysDictResp, error) {
	_, err := l.svcCtx.SysDict.Insert(l.ctx, &model.SysDict{
		DictCode:    in.DictCode,
		DictName:    in.DictName,
		Description: in.Description,
		CreatedBy:    in.CreatedBy,
		UpdatedBy:    in.UpdatedBy,
	})
	if err != nil {
		l.Logger.Errorf("添加字典失败: %v", err)
		return nil, errorx.DatabaseCreateErr
	}
	l.Logger.Infof("添加字典成功: %v", in)
	return &pb.AddSysDictResp{}, nil
}
