package onecprojectservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecProjectLogic {
	return &SearchOnecProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecProjectLogic) SearchOnecProject(in *pb.SearchOnecProjectReq) (*pb.SearchOnecProjectResp, error) {
	var queryParts []string
	var params []interface{}
	if in.Name != "" {
		queryParts = append(queryParts, "name LIKE ? AND")
		params = append(params, "%"+in.Name+"%")
	}
	if in.Identifier != "" {
		queryParts = append(queryParts, "identifier LIKE ? AND")
		params = append(params, "%"+in.Identifier+"%")
	}
	if in.CreatedBy != "" {
		queryParts = append(queryParts, "created_by like ? AND")
		params = append(params, "%"+in.CreatedBy+"%")
	}
	if in.UpdatedBy != "" {
		queryParts = append(queryParts, "updated_by like ? AND")
		params = append(params, "%"+in.UpdatedBy+"%")
	}
	query := utils.RemoveQueryADN(queryParts)

	res, total, err := l.svcCtx.ProjectModel.Search(l.ctx,
		in.OrderStr,
		in.IsAsc,
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("查询项目为空:%v ,sql: %v", err, query)
			return &pb.SearchOnecProjectResp{
				Data:  []*pb.OnecProject{},
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("查询项目失败: %v", err)
		return nil, code.GetProjectErr
	}
	data := make([]*pb.OnecProject, 0, len(res))
	for _, v := range res {
		data = append(data, &pb.OnecProject{
			Id:          v.Id,
			Name:        v.Name,
			Identifier:  v.Identifier,
			Description: v.Description,
			CreatedAt:   v.CreatedAt.Unix(),
			UpdatedAt:   v.UpdatedAt.Unix(),
			CreatedBy:   v.CreatedBy,
			UpdatedBy:   v.UpdatedBy,
		})
	}
	l.Logger.Infof("查询项目成功: %v", data)
	return &pb.SearchOnecProjectResp{
		Data:  data,
		Total: total,
	}, nil
}
