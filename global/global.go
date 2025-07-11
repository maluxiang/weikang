package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/minio/minio-go/v7"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	LocalConfig     LocalConf
	NacosConfig     NacosConf
	DB              *gorm.DB
	Rds             *redis.Client
	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection
	EsClient        *elastic.Client
	MinioClient     *minio.Client
)

type LocalConf struct {
	Nacos struct {
		SpaceId string
		Address string
		Port    int
		DataId  string
		Group   string
	}
}
type NacosConf struct {
	Mysql struct {
		User     string
		Password string
		Host     string
		Port     int
		Database string
	}
	Redis struct {
		Host     string
		Password string
		Db       int
	}
	Es struct {
		Url string
	}
	Mongo struct {
		ApplyURI string
	}
	RabbitMQ struct {
		Url string
	}
	Consul struct {
		RpcName       string
		RpcPortString string
		RpcPortInt    int
		RpcAddress    string
		Address       string
	}
	ChatGpt struct {
		HostUrl   string
		Appid     string
		ApiSecret string
		ApiKey    string
	}
	Minio struct {
		Endpoint        string
		AccessKeyId     string
		SecretAccessKey string
	}
}
