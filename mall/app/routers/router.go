package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"mall/app/handler/v1/cart"
	"mall/app/handler/v1/coupon"
	"mall/app/handler/v1/order"

	"mall/app/conf"
	"mall/app/handler"
	"mall/app/handler/v1"
	"mall/app/handler/v1/goods"
	"mall/app/handler/v1/user"
	mw "mall/app/middleware"
	_ "mall/docs"
	"mall/pkg/app"
	"mall/pkg/net/http/middleware"
)

//NewRouter Load loads the middlewares, routes, handlers.
func NewRouter(c *conf.AppConfig) *gin.Engine {
	g := gin.New()
	// 使用中间件
	g.Use(middleware.NoCache)
	g.Use(middleware.Cors)
	g.Use(middleware.Secure)
	//g.Use(middleware.HandleErrors)

	g.NoRoute(app.RouteNotFound)
	g.NoMethod(app.RouteNotFound)

	// HealthCheck 健康检查路由 使用grpc健康检查
	//router.GET("/health", api.HealthCheck)
	// metrics router 可以在 prometheus 中进行监控
	// 通过 grafana 可视化查看 prometheus 的监控数据，使用插件6671查看
	g.GET("/metrics", gin.WrapH(promhttp.Handler()))
	// HealthCheck 健康检查路由
	g.GET("/health", handler.HealthCheck)
	// 静态资源，主要是图片
	//g.Static("/static", "./static")

	// 返回404，仅在debug环境下开启，线上关闭
	if c.Debug {
		g.Use(gin.Logger())
		// swagger api docs
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		// pprof router 性能分析路由
		// 默认关闭，开发环境下可以打开
		// 访问方式: HOST/debug/pprof
		// 通过 HOST/debug/pprof/profile 生成profile
		// 查看分析图 go tool pprof -http=:5000 profile
		// see: https://github.com/gin-contrib/pprof
		pprof.Register(g)
	} else {
		// disable swagger docs for release  env=release
		g.GET("/swagger/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "env"))
	}
	apiV1 := g.Group("/v1")
	apiV1.Use(
		middleware.RequestID(),
		middleware.Logging(),
		middleware.Prom(middleware.WithNamespace("mall")),
		//middleware.MaxLimiter(c.MaxLimit),
		//middleware.IPRateLimiter(c.IPLimit, c.IPLimitExpr),
	)

	apiV1.GET("/home", v1.Home)
	apiV1.GET("/area", v1.Area)
	apiV1.GET("/pay_list", v1.PayList)
	apiV1.POST("/login", user.Login)
	apiV1.POST("/reg", user.Register)
	apiV1.GET("/home_setting", v1.HomeSetting)
	apiV1.GET("/notice", v1.Notice)
	apiV1.GET("/hot_keyword", v1.HotKeyword)

	version1(apiV1)

	return g
}

func version1(v1 *gin.RouterGroup) {
	// 用户模块
	g := v1.Group("/goods")
	g.Use()
	{
		g.GET("/category", goods.CategoryAll)
		g.POST("/list", goods.List)
		g.GET("/detail", goods.Detail)
		g.GET("/sku", goods.Sku)
	}

	u := v1.Group("/user")
	u.Use(mw.JWT())
	{
		u.POST("/edit", user.Update)
		u.GET("/logout", user.Logout)
	}

	ct := v1.Group("/cart")
	ct.Use(mw.JWT())
	{
		ct.POST("/edit_num", cart.EditNum)
		ct.POST("/edit", cart.Edit)
		ct.POST("/add", cart.Add)
		ct.POST("/del", cart.Del)
		ct.GET("/empty", cart.Empty)
		ct.GET("/list", cart.List)
	}

	cp := v1.Group("/coupon")
	cp.Use(mw.JWT())
	{
		cp.GET("/list", coupon.List)
		cp.GET("/my", coupon.My)
		cp.GET("/draw", coupon.Draw)
	}

	addr := v1.Group("/address")
	addr.Use(mw.JWT())
	{
		addr.GET("/list", user.AddressList)
		addr.POST("/add", user.AddressAdd)
		addr.POST("/edit", user.AddressEdit)
		addr.GET("/del", user.AddressDel)
	}

	od := v1.Group("/order")
	od.Use(mw.JWT())
	{
		od.POST("/submit", order.Submit)
		od.POST("/goods_submit", order.GoodsSubmit)
		od.GET("/detail", order.Detail)
		od.GET("/list", order.List)
		od.POST("/refund", order.Refund)
		od.POST("/receipt", order.Receipt)
		od.POST("/notify", order.Notify)
		od.POST("/cancel", order.Cancel)
		od.POST("/comment", order.Comment)
	}
}
