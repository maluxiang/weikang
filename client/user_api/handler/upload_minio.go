package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"log"
	"net/http"
	"weikang/global"
	"weikang/models"
	"weikang/start"
)

func UploadFile(c *gin.Context) {
	// 解析表单数据
	err := c.Request.ParseMultipartForm(10 << 20) // 允许上传的最大文件大小为 10 MB
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("解析表单数据失败: %s", err.Error()))
		return
	}

	// 获取上传的文件
	file, handler, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("获取上传的文件失败: %s", err.Error()))
		return
	}
	defer file.Close()

	// 初始化 MinIO 客户端
	minioClient, err := start.InitMinioClient()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("初始化 MinIO 客户端失败: %s", err.Error()))
		return
	}

	// 创建一个叫 mybucket 的存储桶。
	bucketName := "mybucket"
	location := "beijing"

	err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(context.Background(), bucketName)
		if err == nil && exists {
			log.Printf("存储桶 %s 已经存在", bucketName)
		} else {
			log.Printf("查询存储桶状态异常: %s", err.Error())
			c.String(http.StatusInternalServerError, fmt.Sprintf("查询存储桶状态异常: %s", err.Error()))
			return
		}
	}
	log.Printf("创建存储桶 %s 成功", bucketName)

	// 上传文件到 MinIO 桶
	objectName := handler.Filename
	contentType := handler.Header.Get("Content-Type")
	putObjectOptions := minio.PutObjectOptions{ContentType: contentType}
	_, err = minioClient.PutObject(context.Background(), bucketName, objectName, file, handler.Size, putObjectOptions)
	if err != nil {
		log.Printf("上传文件到 MinIO 失败: %s", err.Error())
		c.String(http.StatusInternalServerError, fmt.Sprintf("上传文件到 MinIO 失败: %s", err.Error()))
		return
	}
	Up := models.UploadFile{
		FileName: bucketName,
		FilePath: objectName,
		FileType: contentType,
		FileSize: handler.Size,
	}

	err2 := global.DB.Create(&Up).Error
	if err2 != nil {
		log.Printf("文件信息保存到数据库失败: %s", err.Error())
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    500,
			"message": "文件上传失败",
			"Data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "文件上传成功",
		"Data":    objectName,
	})
}
