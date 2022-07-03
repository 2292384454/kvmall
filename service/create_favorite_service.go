//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 09:03:27
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:45:17
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// CreateFavoriteService 收藏创建的服务
type CreateFavoriteService struct {
	UserID    uint `form:"user_id" json:"user_id"`
	ProductID uint `form:"product_id" json:"product_id"`
}

// Create 创建收藏夹
func (service *CreateFavoriteService) Create() serializer.Response {
	var favorite models.Favorite
	code := e.SUCCESS
	models.DB.Where("user_id=? AND product_id=?", service.UserID, service.ProductID).Find(&favorite)
	if favorite == (models.Favorite{}) {
		favorite = models.Favorite{
			UserID:    service.UserID,
			ProductID: service.ProductID,
		}
		if err := models.DB.Create(&favorite).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		code = e.ERROR_EXIST_FAVORITE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
