package onecprojectapplicationservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/shared"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/manager/rpc/pb"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOnecProjectApplicationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOnecProjectApplicationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOnecProjectApplicationLogic {
	return &SearchOnecProjectApplicationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOnecProjectApplicationLogic) SearchOnecProjectApplication(in *pb.SearchOnecProjectApplicationReq) (*pb.SearchOnecProjectApplicationResp, error) {
	var queryParts []string
	var params []interface{}
	if in.ProjectId == 0 {
		l.Logger.Errorf("`projectId` is empty")
		return nil, code.ProjectIdEmptyErr
	}
	queryParts = append(queryParts, "`project_id` = ? AND")
	params = append(params, in.ProjectId)
	if in.ClusterUuid != "" {
		queryParts = append(queryParts, "`cluster_uuid` = ? AND")
		params = append(params, in.ClusterUuid)
	}
	if in.Status != "" {
		queryParts = append(queryParts, "`status` = ? AND")
		params = append(params, in.Status)
	}
	if in.Uuid != "" {
		queryParts = append(queryParts, "`uuid` = ? AND")
		params = append(params, in.Uuid)
	}
	if in.Name != "" {
		queryParts = append(queryParts, "`name` LIKE ? AND")
		params = append(params, "%"+in.Name+"%")
	}
	if in.Identifier != "" {
		queryParts = append(queryParts, "`identifier` LIKE ? AND")
		params = append(params, "%"+in.Identifier+"%")
	}
	query := utils.RemoveQueryADN(queryParts)
	res, total, err := l.svcCtx.ProjectApplicationModel.Search(l.ctx,
		in.OrderStr,
		in.IsAsc,
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("查询项目应用为空:%v ,sql: %v", err, query)
			return &pb.SearchOnecProjectApplicationResp{
				Data:  make([]*pb.OnecProjectApplication, 0),
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("AnnotationsResourceModel err: %v", err)
		return nil, code.SearchNodeAnnotationDBErr
	}
	data := make([]*pb.OnecProjectApplication, len(res))
	for i, v := range res {
		data[i] = shared.ConvertModelToPbApplication(v)
	}
	return &pb.SearchOnecProjectApplicationResp{
		Data:  data,
		Total: total,
	}, nil
}
