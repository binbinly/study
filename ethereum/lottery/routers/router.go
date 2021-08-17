package routers

import (
	"github.com/astaxie/beego"
	"lottery/controllers"
)

func init() {
	beego.Router("/", &controllers.EthController{}, "get:Index")
	beego.Router("/get_accounts", &controllers.EthController{}, "get:GetTouZhuAccounts")
	beego.Router("/post_touzhu", &controllers.EthController{}, "post:PostTouZhu")
	beego.Router("/get_acountInfo", &controllers.EthController{}, "get:GetAcountInfo")
	beego.Router("/search", &controllers.EthController{}, "get:Search")
	beego.Router("/kai_jiang", &controllers.EthController{}, "get:KaiJiang")
	beego.Router("/do_kai_jiang", &controllers.EthController{}, "post:DoKaiJiang")
	beego.Router("/contract", &controllers.EthController{}, "get:Contract")
}
