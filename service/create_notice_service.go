//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:04:25
 * @LastEditors: congz
 * @LastEditTime: 2020-08-04 11:09:13
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// CreateNoticeService 公告创建的服务
type CreateNoticeService struct {
	Text string `form:"text" json:"text"`
}

// Create 公告创建的服务
func (service *CreateNoticeService) Create() serializer.Response {
	notice := models.Notice{
		Text: service.Text,
	}
	code := e.SUCCESS
	err := models.DB.Create(&notice).Error
	if err != nil {
		logging.Info(err)
		code := e.ERROR_DATABASE
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
