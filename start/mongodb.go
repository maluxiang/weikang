package start

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"weikang/global"
)

// Mongo mongodb连接和MongoDB连接池
func Mongo() {
	clientOptions := options.Client().ApplyURI(global.NacosConfig.Mongo.ApplyURI)
	clientOptions.SetMaxPoolSize(20) // 设置最大连接池大小
	clientOptions.SetMinPoolSize(5)  // 设置最小连接池大小

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		zap.S().Fatal("MongoDB连接失败", zap.Error(err))
		return
	}

	// 确保连接池连接的可用性
	err = client.Ping(context.Background(), nil)
	if err != nil {
		zap.S().Fatal("MongoDB连接失败", zap.Error(err))
		return
	}

	global.MongoCollection = client.Database("2110a").Collection("day")
}
