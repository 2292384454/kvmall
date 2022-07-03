//Package models ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-21 23:05:55
 * @LastEditors: congz
 * @LastEditTime: 2020-07-21 23:08:52
 */
package models

import (
	"github.com/jinzhu/gorm"
)

// ProductInfoImg 商品图片模型
type ProductInfoImg struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}
