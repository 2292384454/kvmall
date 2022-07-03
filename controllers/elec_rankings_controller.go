package controllers

import (
	"github.com/kataras/iris/v12"
	"kvmall/serializer"
	"kvmall/service"
)

type ElecRankingController struct {
	Ctx iris.Context //上下文对象
}

func NewElecRankingController() *ElecRankingController {
	return new(ElecRankingController)
}

func (r *ElecRankingController) Get() (result serializer.Response) {
	listElecRankingService := service.ListElecRankingService{}
	return listElecRankingService.List()
}
