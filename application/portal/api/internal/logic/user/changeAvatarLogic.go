package user

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"io"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangeAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangeAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangeAvatarLogic {
	return &ChangeAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangeAvatarLogic) ChangeAvatar(r *http.Request, req *types.ChangeAvatarRequest) (resp string, err error) {
	// 获取头像文件
	avatar, handler, err := r.FormFile("icon")
	if err != nil {
		l.Logger.Errorf("获取文件失败: %v", err)
		return "", err
	}
	defer func(avatar multipart.File) {
		err := avatar.Close()
		if err != nil {
			l.Logger.Errorf("关闭文件失败: %v", err)
			return
		}
	}(avatar)
	l.Logger.Infof("上传的文件名: %s", handler.Filename)
	// 读取文件内容
	imageBytes, err := readFile(avatar, handler.Size)
	if err != nil {
		l.Logger.Errorf("读取文件失败: %v", err)
		return "", err
	}

	// 并发处理上传请求
	res, err := l.uploadImageAsync(imageBytes, handler.Filename, "users")
	if err != nil {
		l.Logger.Errorf("上传图片失败: %v", err)
		return "", err
	}
	l.Logger.Infof("图片上传成功: %v", res)

	_, err = l.svcCtx.SysUserRpc.UpdateIcon(l.ctx, &pb.UpdateIconReq{
		Icon: res.ImageUri,
		Id:   req.Id,
	})
	if err != nil {
		l.Logger.Errorf("更新用户头像失败: %v", err)
		return "", err
	}
	return
}

// 读取文件内容的函数，避免重复代码
func readFile(image io.Reader, size int64) ([]byte, error) {
	imageBytes := make([]byte, size)
	_, err := image.Read(imageBytes)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}
	return imageBytes, nil
}

// 异步上传图片的函数
func (l *ChangeAvatarLogic) uploadImageAsync(imageData []byte, fileName string, project string) (*pb.UploadImageResponse, error) {
	// 使用 goroutine 并发上传，处理完成后返回结果
	var wg sync.WaitGroup
	var res *pb.UploadImageResponse
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		res, err = l.svcCtx.StoreRpc.UploadImage(l.ctx, &pb.UploadImageRequest{
			ImageData: imageData,
			FileName:  fileName,
			Project:   project, // 可以根据实际项目传递动态参数
		})
	}()
	wg.Wait()

	if err != nil {
		return nil, fmt.Errorf("上传图片失败: %w", err)
	}

	return res, nil
}
