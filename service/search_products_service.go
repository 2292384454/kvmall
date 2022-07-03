//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-14 10:47:54
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 11:02:07
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// SearchProductsService 搜索商品的服务
type SearchProductsService struct {
	Search string `form:"search" json:"search"`
}

// Show 搜索商品
func (service *SearchProductsService) Show() serializer.Response {
	products := []models.Product{}
	code := e.SUCCESS

	err := models.DB.Where("name LIKE ?", "%"+service.Search+"%").Find(&products).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products1 := []models.Product{}
	err = models.DB.Where("info LIKE ?", "%"+service.Search+"%").Find(&products1).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products = append(products, products1...)
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProducts(products),
	}
}
