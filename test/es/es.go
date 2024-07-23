package main

import "fmt"

type Person struct {
	id    int
	Name  string // 姓名属性
	Age   int    // 年龄属性
	Email string // 邮箱属性
}

func main() {
	p := Person{
		id:    5,
		Name:  "哈哈哈哈",
		Age:   22,
		Email: "Hao@example.com",
	}

	fmt.Println(p)
	//TODO 添加
	//err := es.ElasticInit("pos", strconv.Itoa(p.id), &p)
	//if err != nil {
	//	fmt.Println("添加失败")
	//} else {
	//	fmt.Println("添加成功")
	//}

	///TODO 删除
	//err := es.ElasticDeletes("pos", "1")
	//if err != nil {
	//	fmt.Println("es删除失败")
	//	return
	//}
	//fmt.Println("es删除成功")

	//TODO 修改
	//err := es.ElasticUpdate("pos", strconv.Itoa(p.id), p)
	//if err != nil {
	//	fmt.Println("es修改失败")
	//	return
	//} else {
	//	fmt.Println("es修改成功")
	//
	//}

	//TODO 搜索
	//name := "李"
	//query := Es.NewMultiMatchQuery(name, "Name")
	//res, err := es.ElasticSingle("pos", query)
	//if err != nil {
	//	fmt.Println("es搜索失败")
	//	return
	//} else {
	//	fmt.Println("es搜索成功")
	//}
	//
	//var list []map[string]interface{}
	//for _, i2 := range res {
	//	list = append(list, i2)
	//}
	//fmt.Println(list)
	//fmt.Println("111111111111111")

	//TODO SearchHighlight搜索+高亮
	//name := "明"
	//query := elastic.NewMultiMatchQuery(name, "Name")
	//high := elastic.NewHighlight()
	//text := elastic.NewHighlighterField("Name")
	//text.PreTags("<span style:=`color:red`>").PostTags("</span>")
	//high.Fields(text)
	//res, err := es.SearchHighlight("pos", query, high)
	//if err != nil {
	//	fmt.Println("es搜索失败")
	//} else {
	//	fmt.Println("es搜索成功")
	//}
	//var list []map[string]interface{}
	//for _, re := range res {
	//	list = append(list, re)
	//}
	//fmt.Println(list)

	//TODO  搜索 + 高亮 + 分页
	//res, err := es.EsSearch("明", 1, 2)
	//if err != nil {
	//	fmt.Println("es搜索失败")
	//} else {
	//	fmt.Println("es搜索成功")
	//}
	//var list []map[string]interface{}
	//for _, re := range res {
	//	list = append(list, re)
	//}
	//fmt.Println(list)

	//TODO  分类搜索
	//classList, err := es.ElasticClassify("pos", "Email.keyword", "yes")
	//if err != nil {
	//	// 处理错误
	//	fmt.Println("Error during classification:", err)
	//	return
	//}
	//// 打印分类结果
	//for _, item := range classList {
	//	fmt.Printf("Type: %s, Count: %d\n", item["types"], item["count"])
	//}
	//TODO  条件搜索
	// IntervalSearch 执行区间查询，返回符合条件的文档内容
	//results, err := es.IntervalSearch("pos", "Age", 5.0)
	//if err != nil {
	//	fmt.Println("Error during search:", err)
	//	return
	//}
	//
	//// 打印查询结果
	//for _, result := range results {
	//	fmt.Println(result)
	//}
	//TODO 搜索求和
	// EsSearchSum 执行搜索并对指定字段进行求和
	//re, err := es.EsSearchSum("pos", "Age")
	//if err != nil {
	//	fmt.Println("求和失败")
	//} else {
	//	fmt.Println("求和成功")
	//}
	//fmt.Println("和是", re)

	//TODO EsSearchSum求平均值
	//res, err := es.ElasticValue("pos", "Age")
	//if err != nil {
	//	fmt.Println("求平均值失败")
	//	return
	//} else {
	//	fmt.Println("求平均值成功")
	//}
	//fmt.Println("平均值", res)
	// EsSearchWithSort TODO 搜索并排序
	//res, err := es.EsSearchWithSort("pos", "Age")
	//if err != nil {
	//	fmt.Println("排序失败")
	//	return
	//} else {
	//	fmt.Println("排序成功")
	//}
	//for _, re := range res {
	//	fmt.Println(re)
	//}

	//TODO 创建IK
	//err := es.CreateIkIndex("demo")
	//if err != nil {
	//	fmt.Println("创建失败")
	//	return
	//} else {
	//	fmt.Println("创建成功")
	//}
	////添加数据
	//err = es.ElasticInit("demo", strconv.Itoa(p.id), &p)
	//if err != nil {
	//	fmt.Println("添加失败")
	//} else {
	//	fmt.Println("添加成功")
	//}

	////TODO  Ik的搜索
	//res, err := es.ElasticSearchIk("demo", "明")
	//if err != nil {
	//	fmt.Println("Ik搜索失败")
	//} else {
	//	fmt.Println("Ik搜索成功")
	//}
	//for _, re := range res {
	//	fmt.Println(re)
	//}

}
