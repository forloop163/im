package controllers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"im-project/server/web/services"
)

type AuthController struct{}

func (a *AuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/", "Login")
}

func (a AuthController) getMd5String2(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func (a AuthController) Login(ctx iris.Context) {
	type userForm struct {
		Name     string
		Password string
	}
	var userFormInfo userForm
	err := ctx.ReadJSON(&userFormInfo)
	if err != nil {
		golog.Infof("用户登录错误, %v", err)
		ctx.JSON(iris.Map{
			"code":    400,
			"message": "登录数据有误",
			"data":    "",
		})
		return
	}
	pwd := a.getMd5String2([]byte(userFormInfo.Password + "im-project"))
	row := services.UserFindByNamePassword(userFormInfo.Name, pwd)
	marshal, _ := json.Marshal(row)
	ctx.JSON(iris.Map{
		"code": 200,
		"message": "success",
		"data": marshal,
	})
}

func (a AuthController) Register(ctx iris.Context) {
	type form struct {
		Name     string `form:"name" json:"id"`
		Password string `form:"password" json:"password"`
		Repassed string `form:"repassed" json:"repassed"`
	}

	var formInfo form
	err := ctx.ReadJSON(&formInfo)
	if err != nil {
		golog.Infof("用户注册错误, %v", err)
		ctx.JSON(iris.Map{
			"code":    400,
			"message": "注册数据有误",
			"data":    "",
		})
		return
	}
	marshal, _ := json.Marshal(formInfo)
	ctx.JSON(iris.Map{
		"code":    200,
		"message": "注册数据有误",
		"data":    marshal,
	})
}

func (a AuthController) ReSetPassword() {}
