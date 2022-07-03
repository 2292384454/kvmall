package controllers

import (
	"kvmall/serializer"
	"kvmall/service"
)

type CategoryController struct {
}

func NewCategoryController() *CategoryController {
	return new(CategoryController)
}

func (r *CategoryController) Get() (result serializer.Response) {
	listCategoriesService := service.ListCategoriesService{}
	return listCategoriesService.List()
}
