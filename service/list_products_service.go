//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:52:23
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 20:32:43
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// ListProductsService 商品列表服务
type ListProductsService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
	CategoryID uint `form:"category_id" json:"category_id"`
}

// List 商品列表
func (service *ListProductsService) List() serializer.Response {
	var products []models.Product

	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}
	// 如果类别 ID 为 0 ，查询所有商品
	// 否则根据类别 ID 查询该类别的商品
	if service.CategoryID == 0 {
		// 1. 查询商品总数
		if err := models.DB.Model(models.Product{}).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 2. 查询所有商品
		if err := models.DB.Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		// 1. 查询该类商品总数
		if err := models.DB.Model(models.Product{}).Where("category_id=?", service.CategoryID).Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		// 2. 查询该类所有商品
		if err := models.DB.Where("category_id=?", service.CategoryID).Limit(service.Limit).Offset(service.Start).Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	// 根据查询结果构造序列化返回
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
