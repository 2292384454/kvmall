//Package models ...
/*
 * @Descripttion:
 * @Author: congz
 * @Date: 2020-06-10 11:11:17
 * @LastEditors: congz
 * @LastEditTime: 2020-07-23 15:23:09
 */
package models

import (
	"kvmall/cache"
	"strconv"

	"github.com/jinzhu/gorm"
)

// Product 商品模型
type Product struct {
	gorm.Model
	Name          string `json:"name"`
	CategoryID    int    `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info" gorm:"size:1000"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
}

// View 获取点击数
func (product *Product) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ProductViewKey(product.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

// AddView 视频游览
func (product *Product) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.ProductViewKey(product.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}

// AddElecRank 增加家电排行点击数
func (product *Product) AddElecRank() {
	// 增加家电排行点击数
	cache.RedisClient.ZIncrBy(cache.ElectricalRank, 1, strconv.Itoa(int(product.ID)))
}

// AddAcceRank 增加配件排行点击数
func (product *Product) AddAcceRank() {
	// 增加配件排行点击数
	cache.RedisClient.ZIncrBy(cache.AccessoryRank, 1, strconv.Itoa(int(product.ID)))
}
