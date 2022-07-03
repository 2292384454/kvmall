//Package service ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 20:04:25
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 18:16:19
 */
package service

import (
	"kvmall/models"
	"kvmall/pkg/e"
	"kvmall/pkg/logging"
	"kvmall/serializer"
)

// UpdateNoticeService 公告更新的服务
type UpdateNoticeService struct {
	NoticeID uint   `form:"notice_id" json:"notice_id"`
	Text     string `form:"text" json:"text"`
}

// Update 公告更新的服务
func (service *UpdateNoticeService) Update() serializer.Response {
	var notice models.Notice
	code := e.SUCCESS
	if err := models.DB.First(&notice, service.NoticeID).Error; err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	notice.Text = service.Text
	if err := models.DB.Save(&notice).Error; err != nil {
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
