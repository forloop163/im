package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"im-project/models"
	"im-project/server/web/services"
)

type AuthController struct{}

func (a *AuthController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("POST", "/login", "Login")
	b.Handle("POST", "/register", "Register")
	b.Handle("GET", "/reset", "ReSetPassword")
}

func (a *AuthController) getMd5String2(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

func (a *AuthController) Login(ctx iris.Context) {
	type form struct {
		Name     string `form:"name" validate:"required"`
		Password string `form:"password" validate:"required"`
	}

	var userForm form
	if err := ctx.ReadForm(&userForm); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(iris.Map{
				"code":    iris.StatusBadRequest,
				"message": err.Error(),
				"data":    "",
			})
			return
		}
	}

	pwd := a.getMd5String2([]byte(userForm.Password + "im-project"))
	db, user := services.UserFindByNamePassword(userForm.Name, pwd)
	if db.RowsAffected == 0 {
		ctx.JSON(iris.Map{
			"code":    iris.StatusNotFound,
			"message": "用户不存在",
			"data":    "",
		})
		return
	}
	ctx.JSON(iris.Map{
		"code":    200,
		"message": "success",
		"data":    user,
	})
}

func (a *AuthController) Register(ctx iris.Context) {
	type form struct {
		Name     string `form:"name" validate:"required"`
		Password string `form:"password" validate:"required"`
		Repassed string `form:"repassed" validate:"required"`
		Email    string `form:"email" validate:"required"`
		Mobile   string `from:"mobile" validate:"required"`
	}

	var regInfo form
	err := ctx.ReadForm(&regInfo)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			ctx.JSON(iris.Map{
				"code":    iris.StatusBadRequest,
				"message": err.Error(),
				"data":    "",
			})
			return
		}
		golog.Infof("用户注册错误, %v", err)
		return
	}

	var user models.User
	user.Name = regInfo.Name
	user.Email = regInfo.Email
	user.Password = a.getMd5String2([]byte(regInfo.Password))
	userId := services.UserCreate(&user)
	ctx.JSON(iris.Map{
		"code":    iris.StatusOK,
		"message": "success",
		"data":    userId,
	})
}

func (a *AuthController) ReSetPassword(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"code":    200,
		"message": "注册数据有误",
	})
	return
}
