package controllers

import (
	"im-project/models"
	"im-project/server/web/services"

	"github.com/kataras/iris/v12/mvc"
)

type SpecialController struct{}

func (s *SpecialController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/{id:uint64}", "GetBy")
	b.Handle("GET", "/detail/{id:uint64}", "GetByDetail")
}

func (s *SpecialController) GetBy(id uint64) models.Special {
	return services.SpecialGetById(id)
}

func (s *SpecialController) GetByDetail(id uint64) models.Special {
	return services.SpecialWithDetailById(id)
}
