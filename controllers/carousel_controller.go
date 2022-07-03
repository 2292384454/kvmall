package controllers

import (
	"kvmall/serializer"
	"kvmall/service"
)

type CarouselController struct {
}

func NewCarouselController() *CarouselController {
	return new(CarouselController)
}

func (r *CarouselController) Get() (result serializer.Response) {
	listCarouselsService := service.ListCarouselsService{}
	return listCarouselsService.List()
}
