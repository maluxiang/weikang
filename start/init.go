package start

func Init() {
	Zap()
	Viper()
	Nacos()
	Mysql()
	InitRedis()
	Mongo()
	Elastic()
	InitMinioClient()
}
