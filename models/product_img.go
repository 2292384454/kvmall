//Package models ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 19:56:52
 * @LastEditors: congz
 * @LastEditTime: 2020-07-22 10:51:04
 */
package models

import (
	"github.com/jinzhu/gorm"
)

// ProductImg 商品图片模型
type ProductImg struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}
