//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 20:39:40
 * @LastEditors: congz
 * @LastEditTime: 2020-07-20 20:59:57
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// DeleteAddressService 购物车删除的服务
type DeleteAddressService struct {
	AddressID uint `form:"address_id" json:"address_id"`
}

// Delete 删除购物车
func (service *DeleteAddressService) Delete() serializer.Response {
	var address models.Address
	code := e.SUCCESS

	err := models.DB.Where("id=?", service.AddressID).Find(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = models.DB.Delete(&address).Error
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
	}
}
