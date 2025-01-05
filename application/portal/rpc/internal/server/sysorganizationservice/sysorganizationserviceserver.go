// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: portal.proto

package server

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/logic/sysorganizationservice"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
)

type SysOrganizationServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSysOrganizationServiceServer
}

func NewSysOrganizationServiceServer(svcCtx *svc.ServiceContext) *SysOrganizationServiceServer {
	return &SysOrganizationServiceServer{
		svcCtx: svcCtx,
	}
}

// -----------------------组织表-----------------------
func (s *SysOrganizationServiceServer) AddSysOrganization(ctx context.Context, in *pb.AddSysOrganizationReq) (*pb.AddSysOrganizationResp, error) {
	l := sysorganizationservicelogic.NewAddSysOrganizationLogic(ctx, s.svcCtx)
	return l.AddSysOrganization(in)
}

func (s *SysOrganizationServiceServer) UpdateSysOrganization(ctx context.Context, in *pb.UpdateSysOrganizationReq) (*pb.UpdateSysOrganizationResp, error) {
	l := sysorganizationservicelogic.NewUpdateSysOrganizationLogic(ctx, s.svcCtx)
	return l.UpdateSysOrganization(in)
}

func (s *SysOrganizationServiceServer) DelSysOrganization(ctx context.Context, in *pb.DelSysOrganizationReq) (*pb.DelSysOrganizationResp, error) {
	l := sysorganizationservicelogic.NewDelSysOrganizationLogic(ctx, s.svcCtx)
	return l.DelSysOrganization(in)
}

func (s *SysOrganizationServiceServer) GetSysOrganizationById(ctx context.Context, in *pb.GetSysOrganizationByIdReq) (*pb.GetSysOrganizationByIdResp, error) {
	l := sysorganizationservicelogic.NewGetSysOrganizationByIdLogic(ctx, s.svcCtx)
	return l.GetSysOrganizationById(in)
}

func (s *SysOrganizationServiceServer) SearchSysOrganization(ctx context.Context, in *pb.SearchSysOrganizationReq) (*pb.SearchSysOrganizationResp, error) {
	l := sysorganizationservicelogic.NewSearchSysOrganizationLogic(ctx, s.svcCtx)
	return l.SearchSysOrganization(in)
}
