package web

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"im-project/server/web/controllers"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Validator = validator.New()
	app.Logger().SetLevel("debug")

	mvc.New(app.Party("/auth")).Handle(new(controllers.AuthController))
	return app
}

func RunWebServer() {
	app := newApp()
	app.Run(iris.Addr(":9527"))
}
