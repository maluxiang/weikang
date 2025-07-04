package start

import (
	"github.com/olivere/elastic/v7"
	"weikang/global"
)

func Elastic() {
	appends := global.NacosConfig.Es.Url
	client, err := elastic.NewClient(elastic.SetURL(appends), elastic.SetSniff(false))
	if err != nil {
		return
	}
	global.EsClient = client
}
