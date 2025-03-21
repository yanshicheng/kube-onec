package sysdictitemservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/pkg/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysDictItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysDictItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysDictItemLogic {
	return &SearchSysDictItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSysDictItemLogic) SearchSysDictItem(in *pb.SearchSysDictItemReq) (*pb.SearchSysDictItemResp, error) {

	// 构建动态 SQL 查询条件
	var queryParts []string
	var params []interface{}
	if in.DictCode != "" {
		queryParts = append(queryParts, "`dict_id` = ? AND")
		params = append(params, in.DictCode)
	}
	if in.ItemText != "" {
		queryParts = append(queryParts, "`item_text` LIKE ? AND")
		params = append(params, "%"+in.ItemText+"%")
	}
	if in.ItemCode != "" {
		queryParts = append(queryParts, "`item_value` LIKE ? AND")
		params = append(params, "%"+in.ItemCode+"%")
	}
	if in.Description != "" {
		queryParts = append(queryParts, "`description` LIKE ? AND")
		params = append(params, "%"+in.Description+"%")
	}
	if in.CreatedBy != "" {
		queryParts = append(queryParts, "`create_by` LIKE ? AND")
		params = append(params, "%"+in.CreatedBy+"%")
	}
	if in.UpdatedBy != "" {
		queryParts = append(queryParts, "`update_by` LIKE ? AND")
		params = append(params, "%"+in.UpdatedBy+"%")
	}

	// 去掉最后一个 " AND "，避免 SQL 语法错误
	query := utils.RemoveQueryADN(queryParts)
	res, total, err := l.svcCtx.SysDictItem.Search(
		l.ctx,
		in.OrderStr, // 使用请求中的 orderStr
		in.IsAsc,    // 使用请求中的 isAsc
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("查询字典项为空: %v, sql: %v", err, query)
			return &pb.SearchSysDictItemResp{
				Data:  make([]*pb.SysDictItem, 0),
				Total: 0,
			}, nil
		}
		l.Logger.Errorf("查询字典项失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	var items []*pb.SysDictItem
	for _, item := range res {
		items = append(items, &pb.SysDictItem{
			Id:          item.Id,
			DictCode:    item.DictCode,
			ItemText:    item.ItemText,
			ItemCode:    item.ItemCode,
			Description: item.Description,
			UpdatedBy:   item.UpdatedBy,
			CreatedBy:   item.CreatedBy,
			CreatedAt:   item.CreatedAt.Unix(),
			UpdatedAt:   item.UpdatedAt.Unix(),
		})
	}

	return &pb.SearchSysDictItemResp{
		Data:  items,
		Total: total,
	}, nil
}
