package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"mall/app/constvar"
	"mall/app/handler"
	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// Home 首页数据
// @Summary 首页数据
// @Description 首页数据
// @Tags 公共
// @Accept json
// @Produce json
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{"id":1,"name":"a","list":[]}}"
// @Router /home [get]
func Home(c *gin.Context) {
	list, err := service.Svc.HomeData(c.Request.Context())
	if err != nil {
		log.Warnf("[v1.comm] home err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// HomeSetting 首页分类数据
// @Summary 首页分类数据
// @Description 首页分类数据
// @Tags 公共
// @Accept json
// @Produce json
// @Param cid body int true "分类id"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":["type":1,"data":{}]}"
// @Router /home_setting [get]
func HomeSetting(c *gin.Context) {
	//商品分类
	cid := cast.ToInt(c.Query("cid"))
	list, err := service.Svc.HomeCatData(c.Request.Context(), cid)
	if err != nil {
		log.Warnf("[v1.comm] HomeSetting err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// Notice 公告
// @Summary 公告
// @Description 公告
// @Tags 公共
// @Accept json
// @Produce json
// @success 0 {object} app.Response{data=[]model.AppNoticeModel} "调用成功结构"
// @Router /notice [get]
func Notice(c *gin.Context) {
	list, err := service.Svc.NoticeList(c.Request.Context(), handler.GetPageOffset(c), constvar.DefaultLimit)
	if err != nil {
		log.Warnf("[v1.comm] notice err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// HotKeyword 热词
// @Summary 热词
// @Description 热词
// @Tags 公共
// @Accept json
// @Produce json
// @success 0 {object} app.Response{data=[]string} "调用成功结构"
// @Router /hot_keyword [get]
func HotKeyword(c *gin.Context) {
	list, err := service.Svc.SearchHotWord(c.Request.Context())
	if err != nil {
		log.Warnf("[v1.comm] hot keyword err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// Area 三级省市县
// @Summary 三级省市县
// @Description 三级省市县
// @Tags 公共
// @Accept json
// @Produce json
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{"province":{},"city":{},"county":{}}}"
// @Router /area [get]
func Area(c *gin.Context) {
	list, err := service.Svc.AreaList(c.Request.Context())
	if err != nil {
		log.Warnf("[v1.comm] area err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// PayList 支付列表
// @Summary 支付列表
// @Description 支付列表
// @Tags 公共
// @Accept json
// @Produce json
// @success 0 {object} app.Response{data=[]model.ConfigPayList} "调用成功结构"
// @Router /pay_list [get]
func PayList(c *gin.Context) {
	list, err := service.Svc.PayList(c.Request.Context())
	if err != nil {
		log.Warnf("[v1.comm] pay list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}