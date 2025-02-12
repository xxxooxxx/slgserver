package loginserver

import (
	"slgserver/db"
	"slgserver/net"
	"slgserver/server/loginserver/controller"
)


var MyRouter = &net.Router{}

func Init() {
	db.TestDB()
	//全部初始化完才注册路由，防止服务器还没启动就绪收到请求
	initRouter()
}

func initRouter() {
	controller.DefaultAccount.InitRouter(MyRouter)
}

