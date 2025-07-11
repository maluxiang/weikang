package pkg

import (
	"bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"mime"
	"path/filepath"
	"weikang/global"
)

func UploadImageToMinio(fileBytes []byte, fileName string) (string, error) {
	bucketName := "mybucket"
	ext := filepath.Ext(fileName) // 例如 ".jpg"
	contentType := mime.TypeByExtension(ext)
	if contentType == "" {
		contentType = "application/octet-stream" // 默认类型
	}

	// 确保桶存在
	ctx := context.Background()
	exists, err := global.MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return "", err
	}
	if !exists {
		err = global.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", err
		}
	}
	// 上传
	_, err = global.MinioClient.PutObject(ctx, bucketName, fileName, bytes.NewReader(fileBytes), int64(len(fileBytes)), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", err
	}
	// 拼接图片访问URL（假设MinIO配置了静态文件服务或反向代理）
	url := fmt.Sprintf("http://14.103.235.214:9000/%s/%s", bucketName, fileName)
	return url, nil
}
