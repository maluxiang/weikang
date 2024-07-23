package start

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"weikang/global"
)

func InitMinioClient() (*minio.Client, error) {
	var err error

	// 初始化 Minio 客户端
	global.MinioClient, err = minio.New(global.NacosConfig.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.NacosConfig.Minio.AccessKeyId, global.NacosConfig.Minio.SecretAccessKey, ""),
		Secure: false,
	})
	//fmt.Println(global.NacosConfig.Minio)
	if err != nil {
		zap.S().Info("Minio 连接失败:", err)
		return nil, err
	} //else {
	//zap.S().Info("Minio 连接成功")
	//}

	return global.MinioClient, nil
}
