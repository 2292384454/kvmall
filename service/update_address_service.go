//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-20 14:40:45
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 10:15:32
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// UpdateAddressService 收货地址修改的服务
type UpdateAddressService struct {
	ID      uint   `form:"id" json:"id"`
	UserID  uint   `form:"user_id" json:"user_id"`
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

// Update 修改购物车信息
func (service *UpdateAddressService) Update() serializer.Response {
	address := models.Address{
		UserID:  service.UserID,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	address.ID = service.ID
	code := e.SUCCESS
	err := models.DB.Save(&address).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	var addresses []models.Address
	err = models.DB.Where("user_id=?", service.UserID).Order("created_at desc").Find(&addresses).Error
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
		Data:   serializer.BuildAddresses(addresses),
	}
}
