//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:08:41
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:48:26
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// ShowImgsService 商品图片详情的服务
type ShowImgsService struct {
}

// Show 商品图片
func (service *ShowImgsService) Show(id string) serializer.Response {
	var imgs []models.ProductImg
	code := e.SUCCESS

	err := models.DB.Where("product_id=?", id).Find(&imgs).Error
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
		Data: serializer.BuildImgs(imgs),
	}
}
