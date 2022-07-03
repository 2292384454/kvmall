//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:11:04
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:52:26
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// CreateCarouselService 轮播图创建的服务
type CreateCarouselService struct {
	ImgPath string `form:"img_path" json:"img_path"`
}

// Create 创建商品
func (service *CreateCarouselService) Create() serializer.Response {
	carousel := models.Carousel{
		ImgPath: service.ImgPath,
	}
	code := e.SUCCESS

	err := models.DB.Create(&carousel).Error
	if err != nil {
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
		Data:   serializer.BuildCarousel(carousel),
	}
}
