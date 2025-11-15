package main

import (
	"github.com/manyodream/todolist-ddd/conf"
	"github.com/manyodream/todolist-ddd/interfaces/initilaize"
)

func main() {
	// 加载配置文件
	conf.InitConfig()

	// 加载配置文件

	// 加载infrastructure

	// 加载路由
	r := initilaize.NewRouter()

	// 启动服务
	_ = r.Run(conf.Todolist.Server.Port)
}
