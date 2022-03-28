package web

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"im-project/server/web/controllers"
)

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/special")).Handle(new(controllers.SpecialController))
	return app
}

func RunWebServer()  {
	app := newApp()
	app.Run(iris.Addr(":9527"))
}