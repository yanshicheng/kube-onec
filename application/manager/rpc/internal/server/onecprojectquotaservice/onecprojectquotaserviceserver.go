// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: manager.proto

package server

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/logic/onecprojectquotaservice"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
)

type OnecProjectQuotaServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedOnecProjectQuotaServiceServer
}

func NewOnecProjectQuotaServiceServer(svcCtx *svc.ServiceContext) *OnecProjectQuotaServiceServer {
	return &OnecProjectQuotaServiceServer{
		svcCtx: svcCtx,
	}
}

// -----------------------项目与集群的对应关系表，记录资源配额和使用情况-----------------------
func (s *OnecProjectQuotaServiceServer) AddOnecProjectQuota(ctx context.Context, in *pb.AddOnecProjectQuotaReq) (*pb.AddOnecProjectQuotaResp, error) {
	l := onecprojectquotaservicelogic.NewAddOnecProjectQuotaLogic(ctx, s.svcCtx)
	return l.AddOnecProjectQuota(in)
}

func (s *OnecProjectQuotaServiceServer) UpdateOnecProjectQuota(ctx context.Context, in *pb.UpdateOnecProjectQuotaReq) (*pb.UpdateOnecProjectQuotaResp, error) {
	l := onecprojectquotaservicelogic.NewUpdateOnecProjectQuotaLogic(ctx, s.svcCtx)
	return l.UpdateOnecProjectQuota(in)
}

func (s *OnecProjectQuotaServiceServer) DelOnecProjectQuota(ctx context.Context, in *pb.DelOnecProjectQuotaReq) (*pb.DelOnecProjectQuotaResp, error) {
	l := onecprojectquotaservicelogic.NewDelOnecProjectQuotaLogic(ctx, s.svcCtx)
	return l.DelOnecProjectQuota(in)
}

func (s *OnecProjectQuotaServiceServer) GetOnecProjectQuota(ctx context.Context, in *pb.GetOnecProjectQuotaReq) (*pb.GetOnecProjectQuotaResp, error) {
	l := onecprojectquotaservicelogic.NewGetOnecProjectQuotaLogic(ctx, s.svcCtx)
	return l.GetOnecProjectQuota(in)
}
