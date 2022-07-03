package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"kvmall/controllers"
)

func InitRouter(app *iris.Application) {
	bathUrl := "/api/v1"
	mvc.New(app.Party(bathUrl + "/products")).Handle(controllers.NewProductController())
	mvc.New(app.Party(bathUrl + "/carousels")).Handle(controllers.NewCarouselController())
	mvc.New(app.Party(bathUrl + "/categories")).Handle(controllers.NewCategoryController())
	mvc.New(app.Party(bathUrl + "/imgs")).Handle(controllers.NewProductImgController())
	mvc.New(app.Party(bathUrl + "/info-imgs")).Handle(controllers.NewProductInfoImgController())
	mvc.New(app.Party(bathUrl + "/param-imgs")).Handle(controllers.NewProductParamImgController())
	mvc.New(app.Party(bathUrl + "/elec-rankings")).Handle(controllers.NewElecRankingController())
	mvc.New(app.Party(bathUrl + "/acce-rankings")).Handle(controllers.NewAcceRankingController())
}
