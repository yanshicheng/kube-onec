package onecprojectservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOnecProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOnecProjectLogic {
	return &DelOnecProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelOnecProjectLogic) DelOnecProject(in *pb.DelOnecProjectReq) (*pb.DelOnecProjectResp, error) {

	err := l.svcCtx.ProjectModel.Delete(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("删除项目失败: %v", err)
		return nil, code.DeleteProjectErr
	}
	l.Logger.Infof("删除项目成功: %v", in.Id)
	return &pb.DelOnecProjectResp{}, nil
}
