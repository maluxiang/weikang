package start

import (
	"github.com/olivere/elastic/v7"
	"weikang/global"
)

func Elastic() {
	client, err := elastic.NewClient(elastic.SetURL(global.NacosConfig.Es.Url), elastic.SetSniff(false))
	if err != nil {
		return
	}
	global.EsClient = client
}
