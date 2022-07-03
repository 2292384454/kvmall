//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:21:05
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:55:15
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// ListCarouselsService 轮播图列表服务
type ListCarouselsService struct {
}

// List 轮播图列表
func (service *ListCarouselsService) List() serializer.Response {
	var carousels []models.Carousel
	code := e.SUCCESS

	if err := models.DB.Find(&carousels).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildCarousels(carousels),
	}
}
