package onecprojectquotaservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecProjectQuotaLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecProjectQuotaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecProjectQuotaLogic {
	return &GetOnecProjectQuotaLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecProjectQuotaLogic) GetOnecProjectQuota(in *pb.GetOnecProjectQuotaReq) (*pb.GetOnecProjectQuotaResp, error) {
	res, err := l.svcCtx.ProjectQuotaModel.FindOneByClusterUuidProjectId(l.ctx, in.ClusterUuid, in.ProjectId)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("项目资源不存在: %v", in.ProjectId)
			return &pb.GetOnecProjectQuotaResp{}, nil
		}
		l.Logger.Errorf("获取项目资源失败: %v", err)
		return nil, err
	}
	l.Logger.Infof("获取项目资源成功: %v", res)

	return &pb.GetOnecProjectQuotaResp{
		Data: shared.ConvertToPBQuota(*res),
	}, nil
}
