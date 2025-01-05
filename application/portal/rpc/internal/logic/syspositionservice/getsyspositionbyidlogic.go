package syspositionservicelogic

import (
	"context"
	"errors"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/model"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"github.com/yanshicheng/kube-onec/common/handler/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSysPositionByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSysPositionByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSysPositionByIdLogic {
	return &GetSysPositionByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSysPositionByIdLogic) GetSysPositionById(in *pb.GetSysPositionByIdReq) (*pb.GetSysPositionByIdResp, error) {
	// 查询职位信息
	resp, err := l.svcCtx.SysPosition.FindOne(l.ctx, in.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			l.Logger.Infof("职位未找到: 职位ID=%d", in.Id)
			return nil, errorx.DatabaseQueryErr
		}
		l.Logger.Errorf("查询职位信息失败: 职位ID=%d, 错误=%v", in.Id, err)
		return nil, errorx.DatabaseQueryErr
	}

	// 转换时间字段
	// 将查询结果拷贝到 pb.SysPosition
	data := &pb.SysPosition{}
	data.Id = resp.Id
	data.Name = resp.Name
	data.UpdateTime = resp.UpdateTime.Unix()
	data.CreateTime = resp.CreateTime.Unix()

	// 返回结果
	l.Logger.Infof("成功获取职位信息: 职位ID=%d", in.Id)
	return &pb.GetSysPositionByIdResp{
		Data: data,
	}, nil
}
