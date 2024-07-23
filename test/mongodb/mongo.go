package main

import (
	"fmt"
	"weikang/start"
)

type day struct {
	ID   int    `bson:"_id"`
	Age  int    `bson:"age"`
	Name string `bson:"name"`
	Addr string `bson:"address"`
}

func main() {
	start.Mongo()
	//// 创建 Student 实例
	stus := []interface{}{
		day{
			ID:   2,
			Age:  18,
			Name: "alice",
			Addr: "beijing",
		},
		day{
			ID:   3,
			Age:  19,
			Name: "bob",
			Addr: "shanghai",
		},
		day{
			ID:   4,
			Age:  20,
			Name: "charlie",
			Addr: "guangzhou",
		},
	}
	fmt.Println(stus)
	fmt.Println("111111111111111111")
	// 插入多条文档
	//err = mongodb.InsertStudents("2111", stus)

	//删除
	//file := bson.M{"_id": 2}
	//err := mongodb.DelDocumentation(file)

	// 更新文档
	//filter := bson.M{"_id": 3} //过滤器，匹配 _id 为 1 的文档
	//$set 用来更新文档中的某些字段
	//update := bson.M{"$set": bson.M{"name": "好好哈"}} // 更新操作，将 name 字段设置为 好好哈
	//err := mongodb.UpdateDocumentation("day", filter, update)

	//更新多条文档
	//filter := bson.M{"age": 88} /// 过滤条件，匹配 age 为 88 的文档
	//update := bson.M{"$set": bson.M{"name": "浩浩荡荡"}}
	//err := mongodb.UpdateDocumentation("day", filter, update)

	////查询文档
	//var s day
	//err := mongodb.QueryID("day", &s) // 传递指针
	//if err != nil {
	//	fmt.Println("查询失败：", err)
	//	return
	//}
	//fmt.Println("成功", s)

	////所有查询
	//// 定义存储结果的结构体切片
	//var stu []day
	//
	//// 查询多个文档并将结果解码到结构体切片中
	//err := mongodb.QueryAllDocuments("day", &stu)
	//if err != nil {
	//	fmt.Println("查询失败：", err)
	//	return
	//}
	//// 打印查询结果
	//for _, stu := range stus {
	//	fmt.Printf("查询结果：%+v\n", stu)
	//}

}
