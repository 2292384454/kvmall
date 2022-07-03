//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 11:03:04
 * @LastEditors: congz
 * @LastEditTime: 2020-07-17 17:54:59
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// DeleteFavoriteService 删除收藏的服务
type DeleteFavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Delete 删除收藏
func (service *DeleteFavoriteService) Delete() serializer.Response {
	var favorite models.Favorite
	code := e.SUCCESS

	err := models.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}

	err = models.DB.Delete(&favorite).Error
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
