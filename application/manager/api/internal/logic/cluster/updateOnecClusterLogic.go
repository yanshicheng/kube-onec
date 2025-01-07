package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecclusterservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOnecClusterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOnecClusterLogic {
	return &UpdateOnecClusterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOnecClusterLogic) UpdateOnecCluster(req *types.UpdateOnecClusterRequest) (resp string, err error) {
	// 调用 RPC 方法，将请求参数映射到 RPC 请求结构
	account, ok := l.ctx.Value("account").(string)
	if !ok {
		return "", err
	}
	_, err = l.svcCtx.ClusterRpc.UpdateOnecCluster(l.ctx, &onecclusterservice.UpdateOnecClusterReq{
		Id:           req.Id,           // 自增主键
		Name:         req.Name,         // 集群名称
		SkipInsecure: req.SkipInsecure, // 是否跳过不安全连接（0：否，1：是）
		Host:         req.Host,         // 集群主机地址
		Token:        req.Token,        // 访问集群的令牌
		ConnCode:     req.ConnCode,     // 连接类型（枚举类型，需要转换）
		EnvCode:      req.EnvCode,      // 集群环境标签
		Location:     req.Location,     // 集群所在地址
		NodeLbIp:     req.NodeLbIp,     // Node 负载均衡 IP
		Description:  req.Description,  // 集群描述信息
		UpdatedBy:     account,          // 更新人
	})
	if err != nil {
		// 错误日志记录
		l.Logger.Errorf("更新集群信息失败: %v", err)
		return "", err
	}

	// 返回成功消息
	return "集群信息更新成功", nil
}
