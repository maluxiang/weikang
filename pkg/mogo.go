package pkg

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"weikang/global"
)

// InsertDocument 插入单个文档
// Set选着自己要添加的集合
// 调用示例在文件mongodb中的test.go文件中
func InsertDocument(set string, document interface{}) error {
	_, err := global.MongoCollection.Database().Collection(set).InsertOne(context.Background(), document)
	if err != nil {
		fmt.Println("插入文档时出错：", err)
		return err
	}
	fmt.Println("文档插入成功！")
	return nil
}

// InsertStudents 将多个文档插入 MongoDB 集合中。
// Set选着自己要添加的集合
// 调用示例在文件mongodb中的test.go文件中
func InsertStudents(set string, documents []interface{}) error {
	// 插入多条文档
	if _, err := global.MongoCollection.Database().Collection(set).InsertMany(context.Background(), documents); err != nil {
		return err
	}
	fmt.Println("文档插入成功！")
	return nil
}

// DelDocumentation 删除文档
// Set选着自己要添加的集合
// 调用示例在文件mongodb中的test.go文件中
func DelDocumentation(set string, data interface{}) error {
	_, err := global.MongoCollection.Database().Collection(set).DeleteOne(context.Background(), data)
	if err != nil {
		fmt.Println("删除文档失败！")
	}
	fmt.Println("删除文档成功！")
	return err
}

// UpdateDocumentation 更新单条单条
// Set选着自己要添加的集合
// 更新文档
// 调用示例在文件mongodb中的test.go文件中
func UpdateDocumentation(set string, filter interface{}, update interface{}) error {
	collection := global.MongoCollection.Database().Collection(set)
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		fmt.Println("更新文档失败：", err)
		return err
	}
	if result.MatchedCount == 0 {
		fmt.Println("没有找到匹配的文档")
		return fmt.Errorf("没有找到匹配的文档")
	}
	fmt.Println("文档更新成功！")
	return nil
}

// UpdatesDocumentation 更新多条文档
// Set选着自己要添加的集合
// 更新文档
// 调用示例在文件mongodb中的test.go文件中
func UpdatesDocumentation(set string, update interface{}) error {
	_, err := global.MongoCollection.Database().Collection(set).UpdateMany(context.Background(), bson.M{}, update)
	if err != nil {
		fmt.Println("更新文档失败！")
	}
	return err
}

// QueryID select * from stu where id = ?
// 查询文档
func QueryID(set string, data interface{}) error {
	res := global.MongoCollection.Database().Collection(set).FindOne(context.Background(), bson.M{})
	if res.Err() != nil {
		return res.Err()
	}

	// 解码结果到传入的参数中
	if err := res.Decode(data); err != nil {
		return err
	}
	fmt.Println("查询成功!")
	return nil
}

// QueryAllDocuments 所有查询，类似select name age from stu where id = ?
// Set选着自己要添加的集合
// 调用示例在文件mongodb中的test.go文件中
// 查询整个集合的所有文档
func QueryAllDocuments(set string, data interface{}) error {
	// 查询文档
	res, err := global.MongoCollection.Database().Collection(set).Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}
	defer res.Close(context.Background())

	// 解码结果到结构体切片
	if err := res.All(context.Background(), data); err != nil {
		return err
	}
	fmt.Println("Document Find successfully!")
	fmt.Printf("Documents found: %+v\n", data)
	return nil
}

// QueryMongoDB 示例：在其他地方使用全局的 MongoDB 客户端执行查询操作
func QueryMongoDB() {
	// 使用全局的 MongoDB 客户端进行查询操作
	cursor, err := global.MongoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		zap.S().Error("查询MongoDB失败", zap.Error(err))
		return
	}
	defer cursor.Close(context.Background())

	// 处理查询结果
	for cursor.Next(context.Background()) {
		var result bson.M
		err := cursor.Decode(&result)
		if err != nil {
			zap.S().Error("解码结果失败", zap.Error(err))
			continue
		}
		// 处理 result
	}
}
