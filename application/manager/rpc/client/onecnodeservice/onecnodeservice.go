// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: manager.proto

package onecnodeservice

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
	GetOnecProjectQuotaReq            = pb.GetOnecProjectQuotaReq
	GetOnecProjectQuotaResp           = pb.GetOnecProjectQuotaResp
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
	SearchOnecProjectReq              = pb.SearchOnecProjectReq
	SearchOnecProjectResp             = pb.SearchOnecProjectResp
	SyncOnecClusterReq                = pb.SyncOnecClusterReq
	SyncOnecClusterResp               = pb.SyncOnecClusterResp
	SyncOnecNodeReq                   = pb.SyncOnecNodeReq
	SyncOnecNodeResp                  = pb.SyncOnecNodeResp
	SyncOnecProjectReq                = pb.SyncOnecProjectReq
	SyncOnecProjectResp               = pb.SyncOnecProjectResp
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

	OnecNodeService interface {
		// -----------------------节点表，用于管理各集群中的节点信息-----------------------
		DelOnecNode(ctx context.Context, in *DelOnecNodeReq, opts ...grpc.CallOption) (*DelOnecNodeResp, error)
		GetOnecNodeById(ctx context.Context, in *GetOnecNodeByIdReq, opts ...grpc.CallOption) (*GetOnecNodeByIdResp, error)
		SearchOnecNode(ctx context.Context, in *SearchOnecNodeReq, opts ...grpc.CallOption) (*SearchOnecNodeResp, error)
		// 节点添加标签
		AddOnecNodeLabel(ctx context.Context, in *AddOnecNodeLabelReq, opts ...grpc.CallOption) (*AddOnecNodeLabelResp, error)
		// 节点删除标签
		DelOnecNodeLabel(ctx context.Context, in *DelOnecNodeLabelReq, opts ...grpc.CallOption) (*DelOnecNodeLabelResp, error)
		// 节点添加注解
		AddOnecNodeAnnotation(ctx context.Context, in *AddOnecNodeAnnotationReq, opts ...grpc.CallOption) (*AddOnecNodeAnnotationResp, error)
		// 节点删除注解
		DelOnecNodeAnnotation(ctx context.Context, in *DelOnecNodeAnnotationReq, opts ...grpc.CallOption) (*DelOnecNodeAnnotationResp, error)
		// 禁止调度
		ForbidScheduled(ctx context.Context, in *ForbidScheduledReq, opts ...grpc.CallOption) (*ForbidScheduledResp, error)
		// 取消禁止调度
		EnableScheduledNode(ctx context.Context, in *EnableScheduledNodeReq, opts ...grpc.CallOption) (*EnableScheduledNodeResp, error)
		// 添加污点
		AddOnecNodeTaint(ctx context.Context, in *AddOnecNodeTaintReq, opts ...grpc.CallOption) (*AddOnecNodeTaintResp, error)
		// 删除污点
		DelOnecNodeTaint(ctx context.Context, in *DelOnecNodeTaintReq, opts ...grpc.CallOption) (*DelOnecNodeTaintResp, error)
		// 同步节点信息
		SyncOnecNode(ctx context.Context, in *SyncOnecNodeReq, opts ...grpc.CallOption) (*SyncOnecNodeResp, error)
		// 驱逐节点pod
		EvictNodePod(ctx context.Context, in *EvictNodePodReq, opts ...grpc.CallOption) (*EvictNodePodResp, error)
		SearchOnecNodeLabelList(ctx context.Context, in *SearchOnecNodeLabelListReq, opts ...grpc.CallOption) (*SearchOnecNodeLabelListResp, error)
		SearchOnecNodeAnnotationList(ctx context.Context, in *SearchOnecNodeAnnotationListReq, opts ...grpc.CallOption) (*SearchOnecNodeAnnotationListResp, error)
		SearchOnecNodeTaintList(ctx context.Context, in *SearchOnecNodeTaintListReq, opts ...grpc.CallOption) (*SearchOnecNodeTaintListResp, error)
	}

	defaultOnecNodeService struct {
		cli zrpc.Client
	}
)

func NewOnecNodeService(cli zrpc.Client) OnecNodeService {
	return &defaultOnecNodeService{
		cli: cli,
	}
}

// -----------------------节点表，用于管理各集群中的节点信息-----------------------
func (m *defaultOnecNodeService) DelOnecNode(ctx context.Context, in *DelOnecNodeReq, opts ...grpc.CallOption) (*DelOnecNodeResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.DelOnecNode(ctx, in, opts...)
}

func (m *defaultOnecNodeService) GetOnecNodeById(ctx context.Context, in *GetOnecNodeByIdReq, opts ...grpc.CallOption) (*GetOnecNodeByIdResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.GetOnecNodeById(ctx, in, opts...)
}

func (m *defaultOnecNodeService) SearchOnecNode(ctx context.Context, in *SearchOnecNodeReq, opts ...grpc.CallOption) (*SearchOnecNodeResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.SearchOnecNode(ctx, in, opts...)
}

// 节点添加标签
func (m *defaultOnecNodeService) AddOnecNodeLabel(ctx context.Context, in *AddOnecNodeLabelReq, opts ...grpc.CallOption) (*AddOnecNodeLabelResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.AddOnecNodeLabel(ctx, in, opts...)
}

// 节点删除标签
func (m *defaultOnecNodeService) DelOnecNodeLabel(ctx context.Context, in *DelOnecNodeLabelReq, opts ...grpc.CallOption) (*DelOnecNodeLabelResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.DelOnecNodeLabel(ctx, in, opts...)
}

// 节点添加注解
func (m *defaultOnecNodeService) AddOnecNodeAnnotation(ctx context.Context, in *AddOnecNodeAnnotationReq, opts ...grpc.CallOption) (*AddOnecNodeAnnotationResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.AddOnecNodeAnnotation(ctx, in, opts...)
}

// 节点删除注解
func (m *defaultOnecNodeService) DelOnecNodeAnnotation(ctx context.Context, in *DelOnecNodeAnnotationReq, opts ...grpc.CallOption) (*DelOnecNodeAnnotationResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.DelOnecNodeAnnotation(ctx, in, opts...)
}

// 禁止调度
func (m *defaultOnecNodeService) ForbidScheduled(ctx context.Context, in *ForbidScheduledReq, opts ...grpc.CallOption) (*ForbidScheduledResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.ForbidScheduled(ctx, in, opts...)
}

// 取消禁止调度
func (m *defaultOnecNodeService) EnableScheduledNode(ctx context.Context, in *EnableScheduledNodeReq, opts ...grpc.CallOption) (*EnableScheduledNodeResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.EnableScheduledNode(ctx, in, opts...)
}

// 添加污点
func (m *defaultOnecNodeService) AddOnecNodeTaint(ctx context.Context, in *AddOnecNodeTaintReq, opts ...grpc.CallOption) (*AddOnecNodeTaintResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.AddOnecNodeTaint(ctx, in, opts...)
}

// 删除污点
func (m *defaultOnecNodeService) DelOnecNodeTaint(ctx context.Context, in *DelOnecNodeTaintReq, opts ...grpc.CallOption) (*DelOnecNodeTaintResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.DelOnecNodeTaint(ctx, in, opts...)
}

// 同步节点信息
func (m *defaultOnecNodeService) SyncOnecNode(ctx context.Context, in *SyncOnecNodeReq, opts ...grpc.CallOption) (*SyncOnecNodeResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.SyncOnecNode(ctx, in, opts...)
}

// 驱逐节点pod
func (m *defaultOnecNodeService) EvictNodePod(ctx context.Context, in *EvictNodePodReq, opts ...grpc.CallOption) (*EvictNodePodResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.EvictNodePod(ctx, in, opts...)
}

func (m *defaultOnecNodeService) SearchOnecNodeLabelList(ctx context.Context, in *SearchOnecNodeLabelListReq, opts ...grpc.CallOption) (*SearchOnecNodeLabelListResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.SearchOnecNodeLabelList(ctx, in, opts...)
}

func (m *defaultOnecNodeService) SearchOnecNodeAnnotationList(ctx context.Context, in *SearchOnecNodeAnnotationListReq, opts ...grpc.CallOption) (*SearchOnecNodeAnnotationListResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.SearchOnecNodeAnnotationList(ctx, in, opts...)
}

func (m *defaultOnecNodeService) SearchOnecNodeTaintList(ctx context.Context, in *SearchOnecNodeTaintListReq, opts ...grpc.CallOption) (*SearchOnecNodeTaintListResp, error) {
	client := pb.NewOnecNodeServiceClient(m.cli.Conn())
	return client.SearchOnecNodeTaintList(ctx, in, opts...)
}
