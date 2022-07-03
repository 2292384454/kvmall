package jsons

import (
	"encoding/json"
	"io/ioutil"
	"kvmall/models"
	"time"
)

func LoadJson() {
	// 连接数据库
	models.Database("kevinhwang:huangkaiyan..QQ@tcp(www.kevinhwang.cn:3306)/cmall?charset=utf8&parseTime=True&loc=Local")

	jsonToStruct := map[string]string{
		"jsons/products.json":         "products",
		"jsons/categories.json":       "categories",
		"jsons/carousels.json":        "carousels",
		"jsons/productimgs.json":      "productimgs",
		"jsons/productinfoimgs.json":  "productinfoimgs",
		"jsons/productparamimgs.json": "productparamimgs",
	}

	for jsonPath, stru := range jsonToStruct {
		jsonBytes, err1 := ioutil.ReadFile(jsonPath)
		if err1 != nil {
			panic(err1)
		}
		switch stru {
		case "products":
			var prods []models.Product
			err1 = json.Unmarshal(jsonBytes, &prods)
			if err1 != nil {
				panic(err1)
			}
			for _, prod := range prods {
				prod.CreatedAt = time.Now()
				prod.UpdatedAt = time.Now()
				err2 := models.DB.Create(prod).Error
				if err2 != nil {
					panic(err2)
				}
			}
		case "categories":
			var categories []models.Category
			err1 = json.Unmarshal(jsonBytes, &categories)
			if err1 != nil {
				panic(err1)
			}
			for _, cate := range categories {
				cate.CreatedAt = time.Now()
				cate.UpdatedAt = time.Now()
				err2 := models.DB.Create(cate).Error
				if err2 != nil {
					panic(err2)
				}
			}
		case "carousels":
			var carousels []models.Carousel
			err1 = json.Unmarshal(jsonBytes, &carousels)
			if err1 != nil {
				panic(err1)
			}
			for _, caro := range carousels {
				caro.CreatedAt = time.Now()
				caro.UpdatedAt = time.Now()
				err2 := models.DB.Create(caro).Error
				if err2 != nil {
					panic(err2)
				}
			}
		case "productimgs":
			var productImgs []models.ProductImg
			err1 = json.Unmarshal(jsonBytes, &productImgs)
			if err1 != nil {
				panic(err1)
			}
			for _, pi := range productImgs {
				pi.CreatedAt = time.Now()
				pi.UpdatedAt = time.Now()
				err2 := models.DB.Create(pi).Error
				if err2 != nil {
					panic(err2)
				}
			}
		case "productinfoimgs":
			var productInfoImgs []models.ProductInfoImg
			err1 = json.Unmarshal(jsonBytes, &productInfoImgs)
			if err1 != nil {
				panic(err1)
			}
			for _, pii := range productInfoImgs {
				pii.CreatedAt = time.Now()
				pii.UpdatedAt = time.Now()
				err2 := models.DB.Create(pii).Error
				if err2 != nil {
					panic(err2)
				}
			}
		case "productparamimgs":
			var productParamImgs []models.ProductParamImg
			err1 = json.Unmarshal(jsonBytes, &productParamImgs)
			if err1 != nil {
				panic(err1)
			}
			for _, ppi := range productParamImgs {
				ppi.CreatedAt = time.Now()
				ppi.UpdatedAt = time.Now()
				err2 := models.DB.Create(ppi).Error
				if err2 != nil {
					panic(err2)
				}
			}
		}
	}
}
