package controllers

import (
	"github.com/kataras/iris/v12"
	"kvmall/serializer"
	"kvmall/service"
)

type AcceRankingController struct {
	Ctx iris.Context //上下文对象
}

func NewAcceRankingController() *AcceRankingController {
	return new(AcceRankingController)
}

func (r *AcceRankingController) Get() (result serializer.Response) {
	listAcceRankingService := service.ListAcceRankingService{}
	return listAcceRankingService.List()
}
