package sysuserservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"

	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateIconLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateIconLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateIconLogic {
	return &UpdateIconLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateIconLogic) UpdateIcon(in *pb.UpdateIconReq) (*pb.UpdateIconResp, error) {
	// 查询用户
	user, err := l.svcCtx.SysUser.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Errorf("用户不存在: %v", err)
			return nil, errorx.DatabaseNotFound
		}
		// 查询过程中的其他错误
		l.Logger.Errorf("查询用户信息失败，用户ID=%d，错误信息：%v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}
	if in.Icon != "" {
		user.Icon = in.Icon
		if err := l.svcCtx.SysUser.Update(l.ctx, user); err != nil {
			l.Logger.Errorf("更新用户信息失败，用户ID=%d，错误信息：%v", in.Id, err)
			return nil, errorx.DatabaseUpdateErr
		}
	}
	l.Logger.Infof("更新用户头像成功，用户ID=%d", in.Id)
	return &pb.UpdateIconResp{}, nil
}
