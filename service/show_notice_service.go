//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 23:12:42
 * @LastEditors: congz
 * @LastEditTime: 2020-08-09 10:18:54
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// ShowNoticeService 公告详情服务
type ShowNoticeService struct {
}

// Show 公告详情服务
func (service *ShowNoticeService) Show() serializer.Response {
	var notice models.Notice
	code := e.SUCCESS
	if err := models.DB.First(&notice, 1).Error; err != nil {
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
		Data:   serializer.BuildNotice(notice),
	}
}
