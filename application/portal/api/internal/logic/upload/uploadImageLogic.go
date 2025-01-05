package upload

import (
	"context"
	"fmt"
	"github.com/yanshicheng/kube-onec/application/portal/rpc/pb"
	"io"
	"net/http"
	"sync"

	"github.com/yanshicheng/kube-onec/application/portal/api/internal/svc"
	"github.com/yanshicheng/kube-onec/application/portal/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(r *http.Request, req *types.UploadImageRequest) (resp *types.UploadImageResponse, err error) {
	// 获取图片文件
	image, handler, err := r.FormFile("image")
	if err != nil {
		l.Logger.Errorf("获取文件失败: %v", err)
		return nil, err
	}
	defer image.Close()

	l.Logger.Infof("上传的文件名: %s", handler.Filename)
	// 读取文件内容
	imageBytes, err := readFile(image, handler.Size)
	if err != nil {
		l.Logger.Errorf("读取文件失败: %v", err)
		return nil, err
	}

	// 并发处理上传请求
	res, err := l.uploadImageAsync(imageBytes, handler.Filename, req.Project)
	if err != nil {
		l.Logger.Errorf("上传图片失败: %v", err)
		return nil, err
	}

	l.Logger.Infof("图片上传成功: %v", res)
	return &types.UploadImageResponse{
		ImageUri: res.ImageUri,
		ImageUrl: res.ImageUrl,
	}, nil
}

// 读取文件内容的函数，避免重复代码
func readFile(image io.Reader, size int64) ([]byte, error) {
	imagebytes := make([]byte, size)
	_, err := image.Read(imagebytes)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}
	return imagebytes, nil
}

// 异步上传图片的函数
func (l *UploadImageLogic) uploadImageAsync(imageData []byte, fileName string, project string) (*pb.UploadImageResponse, error) {
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
