package main

import (
	"github.com/manyodream/todolist-ddd/conf"
	"github.com/manyodream/todolist-ddd/infrastructure/container"
	"github.com/manyodream/todolist-ddd/infrastructure/persistence/dbs"
	"github.com/manyodream/todolist-ddd/interfaces/initilaize"
)

func loadingInfra() {
	conf.InitConfig()
	dbs.MySQLInit()
	container.LoadingDomain()
}

func main() {
	// 加载配置文件
	loadingInfra()

	// 加载路由
	r := initilaize.NewRouter()

	// 启动服务
	_ = r.Run(conf.Todolist.Server.Port)
}
