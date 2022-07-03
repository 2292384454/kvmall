//Package models ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-12 22:09:35
 * @LastEditors: congz
 * @LastEditTime: 2020-09-24 13:39:01
 */
package models

import (
	"github.com/jinzhu/gorm"
)

// Category 分类模型
type Category struct {
	gorm.Model
	CategoryID   uint   `json:"category_id"`
	CategoryName string `json:"category_name"`
	Num          uint   `json:"-"`
}
