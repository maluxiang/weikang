package global

import (
	"github.com/minio/minio-go/v7"
	"github.com/olivere/elastic/v7"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	LocalConfig     LocalConf
	NacosConfig     NacosConf
	DB              *gorm.DB
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
		Dsn string `json:"Dsn"`
	} `json:"Mysql"`
	Redis struct {
		Addr     string `json:"Addr"`
		Password string `json:"Password"`
	} `json:"Redis"`
	Es struct {
		Url string `json:"Url"`
	} `json:"Es"`
	Mongo struct {
		ApplyURI string `json:"ApplyURI"`
	} `json:"Mongo"`
	RabbitMQ struct {
		Url string `json:"Url"`
	} `json:"RabbitMQ"`
	Consul struct {
		RpcName       string `json:"RpcName"`
		RpcPortString string `json:"RpcPortString"`
		RpcPortInt    int    `json:"RpcPortInt"`
		RpcAddress    string `json:"RpcAddress"`
		Address       string `json:"Address"`
	} `json:"Consul"`
	ChatGpt struct {
		HostUrl   string `json:"hostUrl"`
		Appid     string `json:"appid"`
		ApiSecret string `json:"apiSecret"`
		ApiKey    string `json:"apiKey"`
	} `json:"ChatGpt"`
	Emqx struct {
		Agreement string `json:"agreement"`
		Host      string `json:"host"`
		Post      int    `json:"post"`
	} `json:"Emqx"`
	Minio struct {
		Endpoint        string `json:"endpoint"`
		AccessKeyId     string `json:"accessKeyId"`
		SecretAccessKey string `json:"secretAccessKey"`
	} `json:"Minio"`
}
