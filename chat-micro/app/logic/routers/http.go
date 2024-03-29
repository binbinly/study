package routers

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"chat-micro/app/logic/handler/http"
	"chat-micro/app/logic/handler/http/v1/apply"
	"chat-micro/app/logic/handler/http/v1/chat"
	"chat-micro/app/logic/handler/http/v1/collect"
	"chat-micro/app/logic/handler/http/v1/emoticon"
	"chat-micro/app/logic/handler/http/v1/friend"
	"chat-micro/app/logic/handler/http/v1/group"
	"chat-micro/app/logic/handler/http/v1/moment"
	"chat-micro/app/logic/handler/http/v1/upload"
	"chat-micro/app/logic/handler/http/v1/user"
	mw "chat-micro/app/logic/routers/middleware"
	_ "chat-micro/docs"
	"chat-micro/pkg/app"
	"chat-micro/pkg/middleware"
)

//NewRouter Load loads the middlewares, routes, handlers.
func NewRouter(debug bool) *gin.Engine {
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
	g.GET("/health", http.HealthCheck)
	// 静态资源，主要是图片
	//g.Static("/static", "./static")

	// 返回404，仅在debug环境下开启，线上关闭
	if debug {
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
	v1 := g.Group("/v1")
	v1.Use(
		middleware.RequestID(),
		middleware.Logging(),
		middleware.Trace(),
		middleware.Prom(middleware.WithNamespace("chat")),
		//middleware.MaxLimiter(c.MaxLimit),
		//middleware.IPRateLimiter(c.IPLimit, c.IPLimitExpr),
	)

	// 认证相关路由
	v1.POST("/reg", user.Register)
	v1.POST("/login", user.Login)
	v1.POST("/login_phone", user.PhoneLogin)
	v1.POST("/send_code", user.SendCode)

	chatV1(v1)

	return g
}

func chatV1(v1 *gin.RouterGroup) {
	up := v1.Group("/upload")
	up.Use(mw.JWT())
	{
		up.POST("/url", upload.SignUrl)
	}
	// 用户模块
	u := v1.Group("/user")
	u.Use(mw.JWT())
	{
		u.POST("/edit", user.Update)
		u.GET("/profile", user.Profile)
		u.GET("/tag", user.Tag)
		u.GET("/logout", user.Logout)
		u.POST("/search", user.Search)
		u.POST("/report", user.Report)
	}

	// 申请模块
	a := v1.Group("/apply")
	a.Use(mw.JWT())
	{
		a.POST("/friend", apply.Friend)
		a.POST("/handle", apply.Handle)
		a.GET("/list", apply.List)
		a.GET("/count", apply.Count)
	}

	// 好友模块
	f := v1.Group("/friend")
	f.Use(mw.JWT())
	{
		f.GET("/info", friend.Info)
		f.GET("/list", friend.List)
		f.POST("/black", friend.Black)
		f.POST("/star", friend.Star)
		f.POST("/auth", friend.Auth)
		f.POST("/remark", friend.Remark)
		f.POST("/destroy", friend.Destroy)
		f.GET("/tag_list", friend.TagList)
	}

	// 聊天模块
	c := v1.Group("/chat")
	c.Use(mw.JWT())
	{
		c.POST("/detail", chat.Detail)
		c.POST("/send", chat.Send)
		c.POST("/recall", chat.Recall)
	}

	// 群组模块
	gr := v1.Group("/group")
	gr.Use(mw.JWT())
	{
		gr.POST("/create", group.Create)
		gr.POST("/edit", group.Update)
		gr.POST("/nickname", group.UpdateNickname)
		gr.GET("/list", group.List)
		gr.GET("/info", group.Info)
		gr.GET("/user", group.User)
		gr.GET("/quit", group.Quit)
		gr.GET("/join", group.Join)
		gr.POST("/kickoff", group.KickOff)
		gr.POST("/invite", group.Invite)
	}

	coll := v1.Group("/collect")
	coll.Use(mw.JWT())
	{
		coll.POST("/create", collect.Create)
		coll.GET("/list", collect.List)
		coll.POST("/destroy", collect.Destroy)
	}

	mom := v1.Group("/moment")
	mom.Use(mw.JWT())
	{
		mom.POST("/create", moment.Create)
		mom.GET("/list", moment.List)
		mom.GET("/timeline", moment.Timeline)
		mom.POST("/like", moment.Like)
		mom.POST("/comment", moment.Comment)
	}

	emo := v1.Group("/emoticon")
	emo.Use(mw.JWT())
	{
		emo.GET("/list", emoticon.List)
		emo.GET("/cat", emoticon.Cat)
	}
}
