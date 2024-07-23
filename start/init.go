package start

func Init() {
	Zap()
	Viper()
	Nacos()
	Mysql()
	Mongo()
	Elastic()
	InitMinioClient()
}
