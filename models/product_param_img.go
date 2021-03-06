//Package models ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-07-21 23:09:41
 * @LastEditors: congz
 * @LastEditTime: 2020-07-21 23:10:23
 */
package models

import (
	"github.com/jinzhu/gorm"
)

// ProductParamImg 商品图片模型
type ProductParamImg struct {
	gorm.Model
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}
