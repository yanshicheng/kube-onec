package sysuserservicelogic

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/utils"
	"strings"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSysUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSysUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSysUserLogic {
	return &SearchSysUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchSysUserLogic) SearchSysUser(in *pb.SearchSysUserReq) (*pb.SearchSysUserResp, error) {
	// 构建动态 SQL 查询条件
	var queryStr strings.Builder
	var params []interface{}

	// 动态拼接条件
	if in.Account != "" {
		queryStr.WriteString("account LIKE ? AND ")
		params = append(params, "%"+in.Account+"%")
	}
	if in.Email != "" {
		queryStr.WriteString("email LIKE ? AND ")
		params = append(params, "%"+in.Email+"%")
	}
	if in.Mobile != "" {
		queryStr.WriteString("mobile LIKE ? AND ")
		params = append(params, "%"+in.Mobile+"%")
	}
	if in.WorkNumber != "" {
		queryStr.WriteString("work_number LIKE ? AND ")
		params = append(params, "%"+in.WorkNumber+"%")
	}
	if in.UserName != "" {
		queryStr.WriteString("user_name LIKE ? AND ")
		params = append(params, "%"+in.UserName+"%")
	}

	if in.OrganizationId != 0 {
		queryStr.WriteString("organization_id = ? AND ")
		params = append(params, in.OrganizationId)
	}

	if in.PositionId != 0 {
		queryStr.WriteString("position_id = ? AND ")
		params = append(params, in.PositionId)
	}

	if in.HireDate != 0 {

		queryStr.WriteString("hire_date = ? AND ")
		params = append(params, utils.ConvertTimestampToFormattedTime(in.HireDate, "2006-01-02"))
	}
	queryStr.WriteString("is_disabled = ? AND ")
	params = append(params, in.IsDisabled)
	queryStr.WriteString("is_leave = ? AND ")
	params = append(params, in.IsLeave)

	if in.StartLastLoginTime != 0 && in.EndLastLoginTime != 0 {
		queryStr.WriteString("last_login_time >= ? AND last_login_time <= ? ")
		params = append(params, utils.ConvertTimestampToFormattedTime(in.StartLastLoginTime),
			utils.ConvertTimestampToFormattedTime(in.EndLastLoginTime))
	}

	// 去掉最后一个 " AND "，避免 SQL 语法错误
	query := queryStr.String()
	if len(query) > 0 {
		query = query[:len(query)-5] // 去掉 " AND "
	}

	matchedUsers, total, err := l.svcCtx.SysUser.Search(l.ctx, in.OrderStr, in.IsAsc, in.Page, in.PageSize, query, params...)
	if err != nil {
		l.Logger.Errorf("查询用户列表出错: %v", err)
		return nil, code.FindSysUserListErr
	}
	data := make([]*pb.SysUser, len(matchedUsers))
	for i, user := range matchedUsers {

		// 获取 icon 绝对地址
		httpPattern := "http://"
		if l.svcCtx.Config.StorageConf.UseTLS {
			httpPattern = "https://"
		}
		absIcon := fmt.Sprintf("%s%s/%s%s", httpPattern, l.svcCtx.Config.StorageConf.Endpoints[0], l.svcCtx.Config.StorageConf.BucketName, user.Icon)

		data[i] = &pb.SysUser{
			Id:              user.Id,
			UserName:        user.UserName,
			Account:         user.Account,
			Icon:            absIcon,
			Mobile:          user.Mobile,
			Email:           user.Email,
			WorkNumber:      user.WorkNumber,
			HireDate:        user.HireDate.Unix(),
			IsResetPassword: user.IsResetPassword,
			IsDisabled:      user.IsDisabled,
			IsLeave:         user.IsLeave,
			PositionId:      user.PositionId,
			OrganizationId:  user.OrganizationId,
			LastLoginTime:   user.LastLoginTime.Unix(),
			CreatedAt:       user.CreatedAt.Unix(),
			UpdatedAt:       user.UpdatedAt.Unix(),
		}
	}
	return &pb.SearchSysUserResp{
		Data:  data,
		Total: total,
	}, nil
}
