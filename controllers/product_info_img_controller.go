package controllers

import (
	"github.com/kataras/iris/v12"
	"kvmall/serializer"
	"kvmall/service"
)

type ProductInfoImgController struct {
	Ctx iris.Context //上下文对象
}

func NewProductInfoImgController() *ProductInfoImgController {
	return new(ProductInfoImgController)
}

func (r *ProductInfoImgController) GetBy(id string) (result serializer.Response) {
	showInfoImgsService := service.ShowInfoImgsService{}
	return showInfoImgsService.Show(id)
}
