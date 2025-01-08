package onecprojectservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecProjectLogic {
	return &AddOnecProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------项目表，记录项目信息-----------------------
func (l *AddOnecProjectLogic) AddOnecProject(in *pb.AddOnecProjectReq) (*pb.AddOnecProjectResp, error) {
	_, err := l.svcCtx.ProjectModel.Insert(l.ctx, &model.OnecProject{
		Name:        in.Name,
		Identifier:  in.Identifier,
		Description: in.Description,
		CreatedBy:   in.CreatedBy,
		UpdatedBy:   in.UpdatedBy,
	})
	if err != nil {
		l.Logger.Errorf("添加项目失败: %v", err)
		return nil, code.CreateProjectErr
	}
	l.Logger.Infof("添加项目成功: %v", in.Name)
	return &pb.AddOnecProjectResp{}, nil
}
