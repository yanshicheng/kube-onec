package sysdictservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/utils"
	"strings"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysDictLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysDictLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysDictLogic {
	return &SearchSysDictLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSysDictLogic) SearchSysDict(in *pb.SearchSysDictReq) (*pb.SearchSysDictResp, error) {
	// 去掉最后一个 " AND "，避免 SQL 语法错误
	var queryStr strings.Builder
	var params []interface{}
	if in.DictCode != "" {
		queryStr.WriteString("dict_code = ? AND")
		params = append(params, "%"+in.DictCode+"%")
	}
	if in.DictName != "" {
		queryStr.WriteString("dict_name = ? AND")
		params = append(params, "%"+in.DictName+"%")
	}
	if in.Description != "" {
		queryStr.WriteString("description = ? AND")
		params = append(params, "%"+in.Description+"%")
	}
	if in.CreatedBy != "" {
		queryStr.WriteString("create_by = ? AND")
		params = append(params, in.CreatedBy)
	}
	if in.UpdatedBy != "" {
		queryStr.WriteString("update_by = ? AND")
		params = append(params, in.UpdatedBy)
	}
	query := utils.RemoveQueryADN(queryStr)
	res, total, err := l.svcCtx.SysDict.Search(
		l.ctx,
		in.OrderStr, // 使用请求中的 orderStr
		in.IsAsc,    // 使用请求中的 isAsc
		in.Page, in.PageSize,
		query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("查询字典失败: %v", err)
			return &pb.SearchSysDictResp{}, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("查询字典失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	var data []*pb.SysDict
	for _, v := range res {
		data = append(data, &pb.SysDict{
			Id:          v.Id,
			DictName:    v.DictName,
			DictCode:    v.DictCode,
			Description: v.Description,
			UpdatedBy:   v.UpdatedBy,
			UpdatedAt:   v.UpdatedAt.Unix(),
			CreatedBy:   v.CreatedBy,
			CreatedAt:   v.CreatedAt.Unix(),
		})
	}
	return &pb.SearchSysDictResp{
		Data:  data,
		Total: total,
	}, nil
}
