package role

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/client/sysroleservice"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysRoleLogic {
	return &SearchSysRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSysRoleLogic) SearchSysRole(req *types.SearchSysRoleRequest) (resp *types.SearchSysRoleResponse, err error) {
	// 调用 RPC 接口获取数据
	res, err := l.svcCtx.SysRoleRpc.SearchSysRole(l.ctx, &sysroleservice.SearchSysRoleReq{
		Page:        req.Page,
		PageSize:    req.PageSize,
		RoleName:    req.RoleName,
		OrderStr:    req.OrderStr,
		IsAsc:       req.IsAsc,
		Description: req.Description,
		CreateBy:    req.CreateBy,
		UpdateBy:    req.UpdateBy,
	})
	if err != nil {
		// 错误日志，明确上下文
		logx.WithContext(l.ctx).Errorf("查询角色失败: 请求=%+v, 错误=%v", req, err)
		return nil, err
	}

	// 初始化响应对象，指定容量避免动态扩容
	resp = &types.SearchSysRoleResponse{
		Items: make([]types.GetSysRoleByIdResponse, 0, len(res.Data)),
		Total: res.Total,
	}

	// 手动映射数据
	for _, v := range res.Data {
		temp := types.GetSysRoleByIdResponse{
			Id:          v.Id,
			RoleName:    v.RoleName,
			CreateTime:  v.CreateTime,
			UpdateTime:  v.UpdateTime,
			Description: v.Description,
			CreateBy:    v.CreateBy,
			UpdateBy:    v.UpdateBy,
		}
		resp.Items = append(resp.Items, temp)
	}

	// 成功日志，打印关键数据
	logx.WithContext(l.ctx).Infof("查询角色成功: 请求=%+v, 总数=%d", req, res.Total)
	return resp, nil
}
