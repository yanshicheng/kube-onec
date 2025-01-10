package sysroleservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/pkg/utils"
	"strings"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysRoleLogic {
	return &SearchSysRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *SearchSysRoleLogic) SearchSysRole(in *pb.SearchSysRoleReq) (*pb.SearchSysRoleResp, error) {
	// 构建动态 SQL 查询条件
	var queryStr strings.Builder
	var params []interface{}

	// 动态拼接条件
	if in.RoleName != "" {
		queryStr.WriteString("role_name LIKE ? AND")
		params = append(params, "%"+in.RoleName+"%")
	}
	if in.Description != "" {
		queryStr.WriteString("description LIKE ? AND ")
		params = append(params, "%"+in.Description+"%")
	}
	if in.RoleCode != "" {
		queryStr.WriteString("role_code LIKE ? AND ")
		params = append(params, "%"+in.RoleCode+"%")
	}
	if in.CreatedBy != "" {
		queryStr.WriteString("create_by LIKE ? AND ")
		params = append(params, "%"+in.CreatedBy+"%")
	}
	if in.UpdatedBy != "" {
		queryStr.WriteString("update_by LIKE ? AND ")
		params = append(params, "%"+in.UpdatedBy+"%")
	}

	// 去掉最后一个 " AND "，避免 SQL 语法错误
	query := utils.RemoveQueryADN(queryStr)

	// 调用查询
	roles, total, err := l.svcCtx.SysRole.Search(l.ctx, in.OrderStr, in.IsAsc, in.Page, in.PageSize, query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("查询角色失败: %v", err)
			return nil, errorx.DatabaseNotFound
		}
		l.Logger.Errorf("查询角色失败: %v", err)
		return nil, errorx.DatabaseQueryErr
	}

	// 将 created_at 和 update_time 转换为时间戳
	var data []*pb.SysRole
	for _, role := range roles {
		data = append(data, &pb.SysRole{
			Id:          role.Id,
			RoleName:    role.RoleName,
			RoleCode:    role.RoleCode,
			Description: role.Description,
			CreatedAt:   role.CreatedAt.Unix(),
			UpdatedAt:   role.UpdatedAt.Unix(),
			CreatedBy:   role.CreatedBy,
			UpdatedBy:   role.UpdatedBy,
		})
	}

	return &pb.SearchSysRoleResp{
		Data:  data,
		Total: total,
	}, nil
}
