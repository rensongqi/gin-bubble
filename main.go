package main

import (
	"gin-bubble/dao"
	"gin-bubble/routers"
)

func main() {
	// 创建数据库
	// SQL：CREATE DATABASE bubble;
	// 连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// 模型绑定
	dao.InitModel()
	// 注册路由
	r := routers.SetupRouter()
	r.Run(":8080")
}
