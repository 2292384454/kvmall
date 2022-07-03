//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 15:48:10
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 11:02:17
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// ShowCartsService 订单详情的服务
type ShowCartsService struct {
}

// Show 订单
func (service *ShowCartsService) Show(id string) serializer.Response {
	var carts []models.Cart
	code := e.SUCCESS

	err := models.DB.Where("user_id=?", id).Find(&carts).Error
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
		Data:   serializer.BuildCarts(carts),
	}
}
