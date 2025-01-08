package onecprojectservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/utils"
	"strings"

	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"

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
	var queryStr strings.Builder
	var params []interface{}
	if in.Name != "" {
		queryStr.WriteString(" and name like ?")
		params = append(params, "%"+in.Name+"%")
	}
	if in.Identifier != "" {
		queryStr.WriteString(" and identifier like ?")
		params = append(params, "%"+in.Identifier+"%")
	}
	if in.CreatedBy != "" {
		queryStr.WriteString(" and created_by like ?")
		params = append(params, "%"+in.CreatedBy+"%")
	}
	if in.UpdatedBy != "" {
		queryStr.WriteString(" and updated_by like ?")
		params = append(params, "%"+in.UpdatedBy+"%")
	}
	query := utils.RemoveQueryADN(queryStr)

	res, total, err := l.svcCtx.ProjectModel.Search(l.ctx,
		in.OrderStr,
		in.IsAsc,
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return &pb.SearchOnecProjectResp{}, nil
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
