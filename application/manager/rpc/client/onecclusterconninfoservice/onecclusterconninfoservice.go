// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: manager.proto

package onecclusterconninfoservice

import (
	"context"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddOnecClusterConnInfoReq         = pb.AddOnecClusterConnInfoReq
	AddOnecClusterConnInfoResp        = pb.AddOnecClusterConnInfoResp
	AddOnecClusterReq                 = pb.AddOnecClusterReq
	AddOnecClusterResp                = pb.AddOnecClusterResp
	AddOnecNodeAnnotationReq          = pb.AddOnecNodeAnnotationReq
	AddOnecNodeAnnotationResp         = pb.AddOnecNodeAnnotationResp
	AddOnecNodeLabelReq               = pb.AddOnecNodeLabelReq
	AddOnecNodeLabelResp              = pb.AddOnecNodeLabelResp
	AddOnecNodeReq                    = pb.AddOnecNodeReq
	AddOnecNodeResp                   = pb.AddOnecNodeResp
	AddOnecNodeTaintReq               = pb.AddOnecNodeTaintReq
	AddOnecNodeTaintResp              = pb.AddOnecNodeTaintResp
	AddOnecProjectAdminReq            = pb.AddOnecProjectAdminReq
	AddOnecProjectAdminResp           = pb.AddOnecProjectAdminResp
	AddOnecProjectApplicationReq      = pb.AddOnecProjectApplicationReq
	AddOnecProjectApplicationResp     = pb.AddOnecProjectApplicationResp
	AddOnecProjectQuotaReq            = pb.AddOnecProjectQuotaReq
	AddOnecProjectQuotaResp           = pb.AddOnecProjectQuotaResp
	AddOnecProjectReq                 = pb.AddOnecProjectReq
	AddOnecProjectResp                = pb.AddOnecProjectResp
	CancelForbidOnecNodeReq           = pb.CancelForbidOnecNodeReq
	CancelForbidOnecNodeResp          = pb.CancelForbidOnecNodeResp
	DelOnecClusterConnInfoReq         = pb.DelOnecClusterConnInfoReq
	DelOnecClusterConnInfoResp        = pb.DelOnecClusterConnInfoResp
	DelOnecClusterReq                 = pb.DelOnecClusterReq
	DelOnecClusterResp                = pb.DelOnecClusterResp
	DelOnecNodeAnnotationReq          = pb.DelOnecNodeAnnotationReq
	DelOnecNodeAnnotationResp         = pb.DelOnecNodeAnnotationResp
	DelOnecNodeLabelReq               = pb.DelOnecNodeLabelReq
	DelOnecNodeLabelResp              = pb.DelOnecNodeLabelResp
	DelOnecNodeReq                    = pb.DelOnecNodeReq
	DelOnecNodeResp                   = pb.DelOnecNodeResp
	DelOnecNodeTaintReq               = pb.DelOnecNodeTaintReq
	DelOnecNodeTaintResp              = pb.DelOnecNodeTaintResp
	DelOnecProjectAdminReq            = pb.DelOnecProjectAdminReq
	DelOnecProjectAdminResp           = pb.DelOnecProjectAdminResp
	DelOnecProjectApplicationReq      = pb.DelOnecProjectApplicationReq
	DelOnecProjectApplicationResp     = pb.DelOnecProjectApplicationResp
	DelOnecProjectQuotaReq            = pb.DelOnecProjectQuotaReq
	DelOnecProjectQuotaResp           = pb.DelOnecProjectQuotaResp
	DelOnecProjectReq                 = pb.DelOnecProjectReq
	DelOnecProjectResp                = pb.DelOnecProjectResp
	ForbidOnecNodeReq                 = pb.ForbidOnecNodeReq
	ForbidOnecNodeResp                = pb.ForbidOnecNodeResp
	GetOnecClusterByIdReq             = pb.GetOnecClusterByIdReq
	GetOnecClusterByIdResp            = pb.GetOnecClusterByIdResp
	GetOnecClusterConnInfoByIdReq     = pb.GetOnecClusterConnInfoByIdReq
	GetOnecClusterConnInfoByIdResp    = pb.GetOnecClusterConnInfoByIdResp
	GetOnecNodeByIdReq                = pb.GetOnecNodeByIdReq
	GetOnecNodeByIdResp               = pb.GetOnecNodeByIdResp
	GetOnecProjectAdminByIdReq        = pb.GetOnecProjectAdminByIdReq
	GetOnecProjectAdminByIdResp       = pb.GetOnecProjectAdminByIdResp
	GetOnecProjectApplicationByIdReq  = pb.GetOnecProjectApplicationByIdReq
	GetOnecProjectApplicationByIdResp = pb.GetOnecProjectApplicationByIdResp
	GetOnecProjectByIdReq             = pb.GetOnecProjectByIdReq
	GetOnecProjectByIdResp            = pb.GetOnecProjectByIdResp
	GetOnecProjectQuotaByIdReq        = pb.GetOnecProjectQuotaByIdReq
	GetOnecProjectQuotaByIdResp       = pb.GetOnecProjectQuotaByIdResp
	OnecCluster                       = pb.OnecCluster
	OnecClusterConnInfo               = pb.OnecClusterConnInfo
	OnecNode                          = pb.OnecNode
	OnecProject                       = pb.OnecProject
	OnecProjectAdmin                  = pb.OnecProjectAdmin
	OnecProjectApplication            = pb.OnecProjectApplication
	OnecProjectQuota                  = pb.OnecProjectQuota
	SearchOnecClusterConnInfoReq      = pb.SearchOnecClusterConnInfoReq
	SearchOnecClusterConnInfoResp     = pb.SearchOnecClusterConnInfoResp
	SearchOnecClusterReq              = pb.SearchOnecClusterReq
	SearchOnecClusterResp             = pb.SearchOnecClusterResp
	SearchOnecNodeReq                 = pb.SearchOnecNodeReq
	SearchOnecNodeResp                = pb.SearchOnecNodeResp
	SearchOnecProjectAdminReq         = pb.SearchOnecProjectAdminReq
	SearchOnecProjectAdminResp        = pb.SearchOnecProjectAdminResp
	SearchOnecProjectApplicationReq   = pb.SearchOnecProjectApplicationReq
	SearchOnecProjectApplicationResp  = pb.SearchOnecProjectApplicationResp
	SearchOnecProjectQuotaReq         = pb.SearchOnecProjectQuotaReq
	SearchOnecProjectQuotaResp        = pb.SearchOnecProjectQuotaResp
	SearchOnecProjectReq              = pb.SearchOnecProjectReq
	SearchOnecProjectResp             = pb.SearchOnecProjectResp
	SyncOnecClusterReq                = pb.SyncOnecClusterReq
	SyncOnecClusterResp               = pb.SyncOnecClusterResp
	SyncOnecNodeReq                   = pb.SyncOnecNodeReq
	SyncOnecNodeResp                  = pb.SyncOnecNodeResp
	UpdateOnecClusterConnInfoReq      = pb.UpdateOnecClusterConnInfoReq
	UpdateOnecClusterConnInfoResp     = pb.UpdateOnecClusterConnInfoResp
	UpdateOnecClusterReq              = pb.UpdateOnecClusterReq
	UpdateOnecClusterResp             = pb.UpdateOnecClusterResp
	UpdateOnecProjectAdminReq         = pb.UpdateOnecProjectAdminReq
	UpdateOnecProjectAdminResp        = pb.UpdateOnecProjectAdminResp
	UpdateOnecProjectApplicationReq   = pb.UpdateOnecProjectApplicationReq
	UpdateOnecProjectApplicationResp  = pb.UpdateOnecProjectApplicationResp
	UpdateOnecProjectQuotaReq         = pb.UpdateOnecProjectQuotaReq
	UpdateOnecProjectQuotaResp        = pb.UpdateOnecProjectQuotaResp
	UpdateOnecProjectReq              = pb.UpdateOnecProjectReq
	UpdateOnecProjectResp             = pb.UpdateOnecProjectResp

	OnecClusterConnInfoService interface {
		// -----------------------通用的服务连接信息表，动态支持多个服务-----------------------
		AddOnecClusterConnInfo(ctx context.Context, in *AddOnecClusterConnInfoReq, opts ...grpc.CallOption) (*AddOnecClusterConnInfoResp, error)
		UpdateOnecClusterConnInfo(ctx context.Context, in *UpdateOnecClusterConnInfoReq, opts ...grpc.CallOption) (*UpdateOnecClusterConnInfoResp, error)
		DelOnecClusterConnInfo(ctx context.Context, in *DelOnecClusterConnInfoReq, opts ...grpc.CallOption) (*DelOnecClusterConnInfoResp, error)
		GetOnecClusterConnInfoById(ctx context.Context, in *GetOnecClusterConnInfoByIdReq, opts ...grpc.CallOption) (*GetOnecClusterConnInfoByIdResp, error)
		SearchOnecClusterConnInfo(ctx context.Context, in *SearchOnecClusterConnInfoReq, opts ...grpc.CallOption) (*SearchOnecClusterConnInfoResp, error)
	}

	defaultOnecClusterConnInfoService struct {
		cli zrpc.Client
	}
)

func NewOnecClusterConnInfoService(cli zrpc.Client) OnecClusterConnInfoService {
	return &defaultOnecClusterConnInfoService{
		cli: cli,
	}
}

// -----------------------通用的服务连接信息表，动态支持多个服务-----------------------
func (m *defaultOnecClusterConnInfoService) AddOnecClusterConnInfo(ctx context.Context, in *AddOnecClusterConnInfoReq, opts ...grpc.CallOption) (*AddOnecClusterConnInfoResp, error) {
	client := pb.NewOnecClusterConnInfoServiceClient(m.cli.Conn())
	return client.AddOnecClusterConnInfo(ctx, in, opts...)
}

func (m *defaultOnecClusterConnInfoService) UpdateOnecClusterConnInfo(ctx context.Context, in *UpdateOnecClusterConnInfoReq, opts ...grpc.CallOption) (*UpdateOnecClusterConnInfoResp, error) {
	client := pb.NewOnecClusterConnInfoServiceClient(m.cli.Conn())
	return client.UpdateOnecClusterConnInfo(ctx, in, opts...)
}

func (m *defaultOnecClusterConnInfoService) DelOnecClusterConnInfo(ctx context.Context, in *DelOnecClusterConnInfoReq, opts ...grpc.CallOption) (*DelOnecClusterConnInfoResp, error) {
	client := pb.NewOnecClusterConnInfoServiceClient(m.cli.Conn())
	return client.DelOnecClusterConnInfo(ctx, in, opts...)
}

func (m *defaultOnecClusterConnInfoService) GetOnecClusterConnInfoById(ctx context.Context, in *GetOnecClusterConnInfoByIdReq, opts ...grpc.CallOption) (*GetOnecClusterConnInfoByIdResp, error) {
	client := pb.NewOnecClusterConnInfoServiceClient(m.cli.Conn())
	return client.GetOnecClusterConnInfoById(ctx, in, opts...)
}

func (m *defaultOnecClusterConnInfoService) SearchOnecClusterConnInfo(ctx context.Context, in *SearchOnecClusterConnInfoReq, opts ...grpc.CallOption) (*SearchOnecClusterConnInfoResp, error) {
	client := pb.NewOnecClusterConnInfoServiceClient(m.cli.Conn())
	return client.SearchOnecClusterConnInfo(ctx, in, opts...)
}
