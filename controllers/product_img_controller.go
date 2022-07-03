package controllers

import (
	"github.com/kataras/iris/v12"
	"kvmall/serializer"
	"kvmall/service"
)

type ProductImgController struct {
	Ctx iris.Context //上下文对象
}

func NewProductImgController() *ProductImgController {
	return new(ProductImgController)
}

func (r *ProductImgController) GetBy(id string) (result serializer.Response) {
	showImgsService := service.ShowImgsService{}
	return showImgsService.Show(id)
}
