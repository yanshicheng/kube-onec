package sysdictitemservicelogic

import (
	"context"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/pkg/utils"
	"strings"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

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
	var queryStr strings.Builder
	var params []interface{}
	if in.DictCode != "" {
		queryStr.WriteString("dict_id = ? AND")
		params = append(params, in.DictCode)
	}
	if in.ItemText != "" {
		queryStr.WriteString("item_text LIKE ? AND")
		params = append(params, "%"+in.ItemText+"%")
	}
	if in.ItemCode != "" {
		queryStr.WriteString("item_value LIKE ? AND")
		params = append(params, "%"+in.ItemCode+"%")
	}
	if in.Description != "" {
		queryStr.WriteString("description LIKE ? AND")
		params = append(params, "%"+in.Description+"%")
	}
	if in.CreatedBy != "" {
		queryStr.WriteString("create_by LIKE ? AND")
		params = append(params, "%"+in.CreatedBy+"%")
	}
	if in.UpdatedBy != "" {
		queryStr.WriteString("update_by LIKE ? AND")
		params = append(params, "%"+in.UpdatedBy+"%")
	}

	// 去掉最后一个 " AND "，避免 SQL 语法错误
	query := utils.RemoveQueryADN(queryStr)
	res, total, err := l.svcCtx.SysDictItem.Search(
		l.ctx,
		in.OrderStr, // 使用请求中的 orderStr
		in.IsAsc,    // 使用请求中的 isAsc
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
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
