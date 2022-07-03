package controllers

import (
	"github.com/kataras/iris/v12"
	"kvmall/serializer"
	"kvmall/service"
)

type ProductController struct {
	Ctx iris.Context //上下文对象
}

func NewProductController() *ProductController {
	return new(ProductController)
}

func (r *ProductController) Get() (result serializer.Response) {
	categoryID, _ := r.Ctx.URLParamInt("category_id")
	start, _ := r.Ctx.URLParamInt("start")
	limit, _ := r.Ctx.URLParamInt("limit")
	listProductService := service.ListProductsService{Limit: limit, Start: start, CategoryID: uint(categoryID)}
	return listProductService.List()
}

func (r *ProductController) GetBy(id string) (result serializer.Response) {
	showProductService := service.ShowProductService{}
	return showProductService.Show(id)
}
