// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: manager.proto

package onecprojectquotaservice

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
	EnableScheduledNodeReq            = pb.EnableScheduledNodeReq
	EnableScheduledNodeResp           = pb.EnableScheduledNodeResp
	EvictNodePodReq                   = pb.EvictNodePodReq
	EvictNodePodResp                  = pb.EvictNodePodResp
	ForbidScheduledReq                = pb.ForbidScheduledReq
	ForbidScheduledResp               = pb.ForbidScheduledResp
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
	NodeAnnotations                   = pb.NodeAnnotations
	NodeLabels                        = pb.NodeLabels
	NodeTaints                        = pb.NodeTaints
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
	SearchOnecNodeAnnotationListReq   = pb.SearchOnecNodeAnnotationListReq
	SearchOnecNodeAnnotationListResp  = pb.SearchOnecNodeAnnotationListResp
	SearchOnecNodeLabelListReq        = pb.SearchOnecNodeLabelListReq
	SearchOnecNodeLabelListResp       = pb.SearchOnecNodeLabelListResp
	SearchOnecNodeReq                 = pb.SearchOnecNodeReq
	SearchOnecNodeResp                = pb.SearchOnecNodeResp
	SearchOnecNodeTaintListReq        = pb.SearchOnecNodeTaintListReq
	SearchOnecNodeTaintListResp       = pb.SearchOnecNodeTaintListResp
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

	OnecProjectQuotaService interface {
		// -----------------------项目与集群的对应关系表，记录资源配额和使用情况-----------------------
		AddOnecProjectQuota(ctx context.Context, in *AddOnecProjectQuotaReq, opts ...grpc.CallOption) (*AddOnecProjectQuotaResp, error)
		UpdateOnecProjectQuota(ctx context.Context, in *UpdateOnecProjectQuotaReq, opts ...grpc.CallOption) (*UpdateOnecProjectQuotaResp, error)
		DelOnecProjectQuota(ctx context.Context, in *DelOnecProjectQuotaReq, opts ...grpc.CallOption) (*DelOnecProjectQuotaResp, error)
		GetOnecProjectQuotaById(ctx context.Context, in *GetOnecProjectQuotaByIdReq, opts ...grpc.CallOption) (*GetOnecProjectQuotaByIdResp, error)
		SearchOnecProjectQuota(ctx context.Context, in *SearchOnecProjectQuotaReq, opts ...grpc.CallOption) (*SearchOnecProjectQuotaResp, error)
	}

	defaultOnecProjectQuotaService struct {
		cli zrpc.Client
	}
)

func NewOnecProjectQuotaService(cli zrpc.Client) OnecProjectQuotaService {
	return &defaultOnecProjectQuotaService{
		cli: cli,
	}
}

// -----------------------项目与集群的对应关系表，记录资源配额和使用情况-----------------------
func (m *defaultOnecProjectQuotaService) AddOnecProjectQuota(ctx context.Context, in *AddOnecProjectQuotaReq, opts ...grpc.CallOption) (*AddOnecProjectQuotaResp, error) {
	client := pb.NewOnecProjectQuotaServiceClient(m.cli.Conn())
	return client.AddOnecProjectQuota(ctx, in, opts...)
}

func (m *defaultOnecProjectQuotaService) UpdateOnecProjectQuota(ctx context.Context, in *UpdateOnecProjectQuotaReq, opts ...grpc.CallOption) (*UpdateOnecProjectQuotaResp, error) {
	client := pb.NewOnecProjectQuotaServiceClient(m.cli.Conn())
	return client.UpdateOnecProjectQuota(ctx, in, opts...)
}

func (m *defaultOnecProjectQuotaService) DelOnecProjectQuota(ctx context.Context, in *DelOnecProjectQuotaReq, opts ...grpc.CallOption) (*DelOnecProjectQuotaResp, error) {
	client := pb.NewOnecProjectQuotaServiceClient(m.cli.Conn())
	return client.DelOnecProjectQuota(ctx, in, opts...)
}

func (m *defaultOnecProjectQuotaService) GetOnecProjectQuotaById(ctx context.Context, in *GetOnecProjectQuotaByIdReq, opts ...grpc.CallOption) (*GetOnecProjectQuotaByIdResp, error) {
	client := pb.NewOnecProjectQuotaServiceClient(m.cli.Conn())
	return client.GetOnecProjectQuotaById(ctx, in, opts...)
}

func (m *defaultOnecProjectQuotaService) SearchOnecProjectQuota(ctx context.Context, in *SearchOnecProjectQuotaReq, opts ...grpc.CallOption) (*SearchOnecProjectQuotaResp, error) {
	client := pb.NewOnecProjectQuotaServiceClient(m.cli.Conn())
	return client.SearchOnecProjectQuota(ctx, in, opts...)
}
