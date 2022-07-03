package controllers

import (
	"github.com/kataras/iris/v12"
	"kvmall/serializer"
	"kvmall/service"
)

type ProductParamImgController struct {
	Ctx iris.Context //上下文对象
}

func NewProductParamImgController() *ProductParamImgController {
	return new(ProductParamImgController)
}

func (r *ProductParamImgController) GetBy(id string) (result serializer.Response) {
	showParamImgsService := service.ShowParamImgsService{}
	return showParamImgsService.Show(id)
}
