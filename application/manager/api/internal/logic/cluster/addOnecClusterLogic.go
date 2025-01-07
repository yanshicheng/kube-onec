package cluster

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOnecClusterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOnecClusterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOnecClusterLogic {
	return &AddOnecClusterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOnecClusterLogic) AddOnecCluster(req *types.AddOnecClusterRequest) (resp string, err error) {
	// 调用 RPC 方法，填充 AddOnecClusterReq
	account, ok := l.ctx.Value("account").(string)
	if !ok {
		return "", err
	}
	_, err = l.svcCtx.ClusterRpc.AddOnecCluster(l.ctx, &pb.AddOnecClusterReq{
		Name:         req.Name,         // 集群名称，从前端请求中获取
		SkipInsecure: req.SkipInsecure, // 是否跳过不安全连接
		Host:         req.Host,         // 集群主机地址
		Token:        req.Token,        // 访问集群的令牌
		ConnCode:     req.ConnCode,     // 连接类型（枚举值，需要转换）
		EnvCode:      req.EnvCode,      // 集群环境标签
		Location:     req.Location,     // 集群所在地址
		NodeLbIp:     req.NodeLbIp,     // Node 负载均衡 IP
		Description:  req.Description,  // 集群描述信息
		CreatedBy:     account,          // 记录创建人
		UpdatedBy:     account,          // 记录更新人
	})
	if err != nil {
		// 返回错误
		l.Logger.Errorf("添加集群失败: %v", err)
		return "", err
	}

	return "添加集群成功!", nil
}
