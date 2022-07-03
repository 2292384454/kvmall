//Package models ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 14:03:57
 * @LastEditors: congz
 * @LastEditTime: 2020-08-12 21:02:15
 */
package models

import (
	"github.com/jinzhu/gorm"
)

// Carousel 轮播图模型
type Carousel struct {
	gorm.Model
	ImgPath   string `json:"img_path"`
	ProductID uint   `json:"product_id"`
}
