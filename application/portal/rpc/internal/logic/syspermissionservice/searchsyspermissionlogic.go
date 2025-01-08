package syspermissionservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/yanshicheng/kube-onec/utils"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysPermissionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysPermissionLogic {
	return &SearchSysPermissionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSysPermissionLogic) SearchSysPermission(in *pb.SearchSysPermissionReq) (*pb.SearchSysPermissionResp, error) {
	// 搜索
	// 构建 SQL 查询条件
	// 构建动态 SQL 查询条件
	var queryStr strings.Builder
	var params []interface{}

	// 动态拼接条件
	if in.Name != "" {
		queryStr.WriteString("name LIKE ? AND ")
		params = append(params, "%"+in.Name+"%")
	}
	if in.Uri != "" {
		queryStr.WriteString("uri LIKE ? AND ")
		params = append(params, "%"+in.Uri+"%")
	}
	if in.Action != "" {
		queryStr.WriteString("action LIKE ? AND ")
		params = append(params, "%"+in.Action+"%")
	}

	// 去掉最后一个 " AND "，避免 SQL 语法错误
	query := utils.RemoveQueryADN(queryStr)
	permissions, total, err := l.svcCtx.SysPermission.Search(l.ctx, in.OrderStr, in.IsAsc, in.Page, in.PageSize, query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("查询权限失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}
	var permissionList []*pb.SysPermission
	for _, permission := range permissions {
		permissionList = append(permissionList, &pb.SysPermission{
			Id:        permission.Id,
			ParentId:  permission.ParentId,
			Name:      permission.Name,
			Uri:       permission.Uri,
			Action:    permission.Action,
			Level:     permission.Level,
			CreatedAt: permission.CreatedAt.Unix(),
			UpdatedAt: permission.UpdatedAt.Unix(),
		})

	}
	return &pb.SearchSysPermissionResp{
		Data:  permissionList,
		Total: total,
	}, nil
}
