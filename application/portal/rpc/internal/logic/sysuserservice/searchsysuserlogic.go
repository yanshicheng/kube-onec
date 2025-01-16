package sysuserservicelogic

import (
	"context"
	"errors"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/code"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	utils2 "github.com/yanshicheng/kube-onec/pkg/utils"

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
	var queryParts []string
	var params []interface{}

	// 动态拼接条件
	if in.Account != "" {
		queryParts = append(queryParts, "`account` LIKE ? AND ")
		params = append(params, "%"+in.Account+"%")
	}
	if in.Email != "" {
		queryParts = append(queryParts, "`email` LIKE ? AND ")
		params = append(params, "%"+in.Email+"%")
	}
	if in.Mobile != "" {
		queryParts = append(queryParts, "`mobile` LIKE ? AND ")
		params = append(params, "%"+in.Mobile+"%")
	}
	if in.WorkNumber != "" {
		queryParts = append(queryParts, "`work_number` LIKE ? AND ")
		params = append(params, "%"+in.WorkNumber+"%")
	}
	if in.UserName != "" {
		queryParts = append(queryParts, "`user_name` LIKE ? AND ")
		params = append(params, "%"+in.UserName+"%")
	}

	if in.OrganizationId != 0 {
		queryParts = append(queryParts, "`organization_id` = ? AND ")
		params = append(params, in.OrganizationId)
	}

	if in.PositionId != 0 {
		queryParts = append(queryParts, "`position_id` = ? AND ")
		params = append(params, in.PositionId)
	}

	if in.HireDate != 0 {
		queryParts = append(queryParts, "`hire_date` = ? AND ")
		params = append(params, utils2.ConvertTimestampToFormattedTime(in.HireDate, "2006-01-02"))
	}
	queryParts = append(queryParts, "`is_disabled` = ? AND ")
	params = append(params, in.IsDisabled)
	queryParts = append(queryParts, "`is_leave` = ? AND ")
	params = append(params, in.IsLeave)

	if in.StartLastLoginTime != 0 && in.EndLastLoginTime != 0 {
		queryParts = append(queryParts, "`last_login_time` >= ? AND `last_login_time` <= ? ")
		params = append(params, utils2.ConvertTimestampToFormattedTime(in.StartLastLoginTime),
			utils2.ConvertTimestampToFormattedTime(in.EndLastLoginTime))
	}

	// 去掉最后一个 " AND "，避免 SQL 语法错误
	query := utils2.RemoveQueryADN(queryParts)

	matchedUsers, total, err := l.svcCtx.SysUser.Search(l.ctx, in.OrderStr, in.IsAsc, in.Page, in.PageSize, query, params...)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("查询用户为空:%v ,sql: %v", err, query)
			return &pb.SearchSysUserResp{
				Data:  make([]*pb.SysUser, 0),
				Total: 0,
			}, nil
		}
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
