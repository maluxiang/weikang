package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	"log"
	"weikang/global"
)

// esv7多条件搜索
func EsSearchWhere(index string, begintime, endtime string, name string, status, page, size int) ([]byte, error) {
	var boolQuery *elastic.BoolQuery

	// 构建基本的BoolQuery，包括非可选的查询条件
	boolQuery = elastic.NewBoolQuery().
		Filter(
			elastic.NewRangeQuery("created_at").Gte(begintime).Lte(endtime),
			elastic.NewTermQuery("status", status),
			elastic.NewBoolQuery().MustNot(elastic.NewExistsQuery("deleted_at")),
		)

	// 如果name不为空，则添加MatchQuery到BoolQuery中
	if name != "" {
		boolQuery.Must(elastic.NewMatchQuery("name", name))
	}

	offset := (page - 1) * size

	// 使用设置好的BoolQuery进行搜索
	search := global.EsClient.Search().
		Index(index).
		Query(boolQuery).
		From(offset). // 设置起始位置
		Size(size).   // 设置每页大小
		Pretty(true)
	result, err := search.Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("search::::::::", err.Error())
	}
	var map1 []map[string]interface{}
	for _, hit := range result.Hits.Hits {
		var map2 map[string]interface{}
		json.Unmarshal(hit.Source, &map2)
		map1 = append(map1, map2)
	}
	marshal, err := json.Marshal(map1)
	fmt.Print("搜索返回值", string(marshal))
	if err != nil {
		return nil, fmt.Errorf("json解析：：；", err.Error())
	}
	return marshal, nil
}

// ElasticInit es添加
// 调用时实例在test中的elastic
func ElasticInit(Index string, Id string, data interface{}) error {
	_, err := global.EsClient.Index().Index(Index).Id(Id).BodyJson(data).Do(context.Background())
	return err
}

// ES删除
// 调用时实例在test中的elastic
func ElasticDeletes(index string, Id string) (err error) {
	_, err = global.EsClient.Delete().Index(index).Id(Id).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// 修改
// 调用时实例在test中的elastic
func ElasticUpdate(index string, id string, doc interface{}) (err error) {
	_, err = global.EsClient.Update().Index(index).Id(id).Doc(doc).Do(context.Background())
	return err
}

// ElasticSingle 搜索
func ElasticSingle(index string, query elastic.Query) ([]map[string]interface{}, error) {
	result, err := global.EsClient.Search().Index(index).Query(query).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var list []map[string]interface{}
	for _, hit := range result.Hits.Hits {
		tmp := make(map[string]interface{})
		err := json.Unmarshal(hit.Source, &tmp)
		if err != nil {
			fmt.Println("解析失败")
		}
		list = append(list, tmp)
	}
	return list, err
}

// SearchHighlight 搜索+高亮
func SearchHighlight(index string, query elastic.Query, highlight *elastic.Highlight) ([]map[string]interface{}, error) {
	result, err := global.EsClient.Search().Index(index).Query(query).Highlight(highlight).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var list []map[string]interface{}
	for _, hit := range result.Hits.Hits {
		tmp := make(map[string]interface{})
		err := json.Unmarshal(hit.Source, &tmp)
		if err != nil {
			fmt.Println("解析失败")
		}
		for key, val := range hit.Highlight {
			tmp[key] = val[0]
		}
		list = append(list, tmp)
	}

	return list, err
}

// EsSearch TODO 搜索+高亮+分页
func EsSearch(keyword string, page int64, pageSize int64) ([]map[string]interface{}, error) {
	query := elastic.NewMultiMatchQuery(keyword).Fuzziness("AUTO")
	result, err := global.EsClient.Search().Query(query).Highlight(elastic.NewHighlight().Fields(elastic.NewHighlighterField("*"))).From(int((page - 1) * pageSize)).Size(int(pageSize)).Do(context.Background())
	if err != nil {
		return nil, err
	}
	i := result.Hits.TotalHits.Value / pageSize
	if result.Hits.TotalHits.Value%pageSize != 0 {
		i++
	}
	if page < 1 {
		page = 1
	}
	if page > i {
		page = i
	}
	var results []map[string]interface{}
	for _, hit := range result.Hits.Hits {
		results = append(results, map[string]interface{}{
			"Source":    string(hit.Source),
			"Highlight": hit.Highlight,
		})
	}
	return results, nil
}

// 分类
func ElasticClassify(index string, data string, name string) ([]map[string]interface{}, error) {
	//要分类的字段
	//Name随便取一个名称
	polymerization := elastic.NewTermsAggregation().Field(data) //要搜索的字段
	res, err := global.EsClient.Search().Index(index).Aggregation(name, polymerization).Do(context.Background())
	if err != nil {
		return nil, err
	}
	result, found := res.Aggregations.Terms(name)

	if !found {
		return nil, err
	}
	var list []map[string]interface{}
	for _, v := range result.Buckets {
		types := map[string]interface{}{
			"types": v.Key.(string),
			"count": v.DocCount,
		}
		list = append(list, types)

	}
	return list, err
}

// IntervalSearch 执行区间查询，返回符合条件的文档内容
func IntervalSearch(index string, field string, start float64) ([]string, error) {
	// 创建范围查询条件
	rangeQuery := elastic.NewRangeQuery(field).Gt(start)
	// 执行搜索并限制返回结果数量为100
	result, err := global.EsClient.Search().Index(index).Query(rangeQuery).Size(100).Do(context.Background())
	if err != nil {
		return nil, err
	}
	var results []string
	// 处理搜索结果
	for _, hit := range result.Hits.Hits {
		results = append(results, string(hit.Source))
	}
	return results, nil
}

// EsSearchSum 执行搜索并对指定字段进行求和
func EsSearchSum(index string, field string) (float64, error) {
	// 创建求和聚合查询
	sumQuery := elastic.NewSumAggregation().Field(field) //要搜索的字段
	// 执行搜索并聚合求和
	result, err := global.EsClient.Search().Index(index).Query(elastic.NewMatchAllQuery()).Aggregation(field, sumQuery).Do(context.Background())
	if err != nil {
		return 0, err
	}
	// 解析聚合结果
	res, ok := result.Aggregations.Sum(field)
	if !ok {
		return 0, nil
	}
	return *res.Value, nil
}

// 求平均值
func ElasticValue(index string, field string) (float64, error) {
	// 创建求和聚合查询
	sumQuery := elastic.NewAvgAggregation().Field(field)
	// 执行搜索并聚合求和
	result, err := global.EsClient.Search().Index(index).Query(elastic.NewMatchAllQuery()).Aggregation(field, sumQuery).Do(context.Background())
	if err != nil {
		return 0, err
	}
	// 解析聚合结果
	res, ok := result.Aggregations.Avg(field)
	if !ok {
		return 0, nil
	}
	return *res.Value, nil
}

// EsSearchWithSort TODO 搜索并排序
func EsSearchWithSort(index string, field string) ([]string, error) {
	searchResult, err := global.EsClient.Search().
		Index(index).
		Query(elastic.NewMatchAllQuery()). // 查询所有文档
		Sort(field, false).                // 根据指定字段进行降序排序
		Do(context.Background())
	if err != nil {
		log.Fatalf("Error executing search: %s", err)
	}
	// 返回搜索结果
	var results []string
	for _, hit := range searchResult.Hits.Hits {
		results = append(results, string(hit.Source))
	}
	return results, nil
}

// EsSearch 使用指定的关键字在给定的Elasticsearch索引中执行搜索操作，并返回搜索结果及其高亮信息。
func ElasticSearchIk(index string, keyword string) ([]interface{}, error) {
	// 创建基于关键字的多字段匹配查询，并指定分析器为IK分词器
	query := elastic.NewMultiMatchQuery(keyword, "Title", "Description").Analyzer("ik_max_word")

	// 执行搜索请求
	result, err := global.EsClient.Search().
		Index(index).
		Query(query).
		Highlight(elastic.NewHighlight().Field("*")).
		Sort("Stock", false).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	// 处理搜索结果
	var res []interface{}
	for _, hit := range result.Hits.Hits {
		// 将搜索结果和高亮信息保存到一个map中，并添加到结果切片中
		res = append(res, map[string]interface{}{
			"Source":    string(hit.Source), // 将hit.Source转换为字符串格式
			"Highlight": hit.Highlight,      // 高亮信息
		})
	}
	return res, nil
}

// CreateIkIndex TODO 创建IK分词器索引
func CreateIkIndex(indexName string) error {
	_, err := global.EsClient.CreateIndex(indexName).BodyString(
		`{
 "settings": {
    "index": {
      "analysis": {
        "analyzer": {
          "default": {
            "type": "ik_max_word"
          }
        }
      }
    }
  }
}`).Do(context.Background())
	if err != nil {
		log.Println("创建索引失败")
	}
	return nil
}
