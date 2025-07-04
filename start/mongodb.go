package start

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"weikang/global"
)

// Mongo mongodb连接和MongoDB连接池
func Mongo() {
	var err error
	// 设置 MongoDB 连接选项
	clientOptions := options.Client().ApplyURI(global.NacosConfig.Mongo.ApplyURI)
	clientOptions.SetMaxPoolSize(20) // 设置最大连接池大小
	clientOptions.SetMinPoolSize(5)  // 设置最小连接池大小
	clientOptions.SetConnectTimeout(10 * time.Second)
	// 连接到 MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	global.MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("连接 MongoDB 失败：%v", err)
	}
	// 验证连接
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = global.MongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("验证 MongoDB 连接失败：%v", err)
	}

	global.MongoCollection = global.MongoClient.Database("2210a").Collection("users")
}
