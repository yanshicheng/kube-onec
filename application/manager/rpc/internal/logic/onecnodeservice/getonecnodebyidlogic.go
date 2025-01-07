package onecnodeservicelogic

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOnecNodeByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOnecNodeByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnecNodeByIdLogic {
	return &GetOnecNodeByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOnecNodeByIdLogic) GetOnecNodeById(in *pb.GetOnecNodeByIdReq) (*pb.GetOnecNodeByIdResp, error) {
	res, err := l.svcCtx.NodeModel.FindOne(l.ctx, in.Id)
	if err != nil {
		l.Logger.Errorf("获取节点信息失败: %v", err)
		return nil, err
	}
	data := ConvertDbModelToPbModel(res)
	return &pb.GetOnecNodeByIdResp{
		Data: data,
	}, nil
}
