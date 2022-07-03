package main

import (
	"github.com/kataras/iris/v12"
	"kvmall/conf"
	"kvmall/route"
	"os"
)

func main() {
	//jsons.LoadJson()

	// 从配置文件读取配置
	conf.Init()

	app := newApp()
	route.InitRouter(app)
	err := app.Run(iris.Addr(":"+os.Getenv("PORT")), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}

func newApp() *iris.Application {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	app.AllowMethods(iris.MethodOptions)
	return app
}
