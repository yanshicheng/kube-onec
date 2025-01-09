package onec_project

import (
	"context"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/client/onecprojectservice"

	"github.com/yanshicheng/kube-onec/application/manager/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecProjectLogic {
	return &SearchOnecProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchOnecProjectLogic) SearchOnecProject(req *types.SearchOnecProjectRequest) (resp *types.SearchOnecProjectResponse, err error) {
	res, err := l.svcCtx.ProjectRpc.SearchOnecProject(l.ctx, &onecprojectservice.SearchOnecProjectReq{
		Page:       req.Page,
		PageSize:   req.PageSize,
		OrderStr:   req.OrderStr,
		IsAsc:      req.IsAsc,
		Identifier: req.Identifier,
		Name:       req.Name,
		CreatedBy:  req.CreatedBy,
		UpdatedBy:  req.UpdatedBy,
	})
	if err != nil {
		// 错误日志，明确上下文
		l.Logger.Errorf("查询项目信息失败: 请求=%+v, 错误=%v", req, err)
		return nil, err
	}
	data := make([]types.OnecProject, len(res.Data))
	for i, v := range res.Data {
		data[i] = types.OnecProject{
			Id:          v.Id,
			Identifier:  v.Identifier,
			Name:        v.Name,
			CreatedBy:   v.CreatedBy,
			CreatedAt:   v.CreatedAt,
			UpdatedBy:   v.UpdatedBy,
			UpdatedAt:   v.UpdatedAt,
			Description: v.Description,
		}
	}
	return &types.SearchOnecProjectResponse{
		Items: data,
		Total: res.Total,
	}, nil
}
