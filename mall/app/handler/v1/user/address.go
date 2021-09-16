package user

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"mall/app/service"
	"mall/pkg/app"
	"mall/pkg/errno"
	"mall/pkg/log"
)

// AddressList 用户收货地址
// @Summary 用户收货地址
// @Description 用户收货地址
// @Tags 收货地址
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @success 0 {object} app.Response{data=[]model.UserAddress} "调用成功结构"
// @Router /address/list [get]
func AddressList(c *gin.Context) {
	list, err := service.Svc.UserAddressList(c.Request.Context(), app.GetUserID(c))
	if err != nil {
		log.Warnf("[v1.user] address list err: %v", err)
		app.Error(c, errno.ErrEmpty)
		return
	}
	app.Success(c, list)
}

// AddressAdd 添加收货地址
// @Summary 添加收货地址
// @Description 添加收货地址
// @Tags 收货地址
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body AddressAddParams true "sku"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /address/add [post]
func AddressAdd(c *gin.Context) {
	var req AddressAddParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	id, err := service.Svc.UserAddressAdd(c.Request.Context(), app.GetUserID(c), req.Name, req.Phone, req.Province, req.City, req.County, req.Detail, req.AreaCode, req.IsDefault)
	if err != nil {
		log.Warnf("[v1.user] add address err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.Success(c, id)
}

// AddressEdit 修改收货地址
// @Summary 修改收货地址
// @Description 修改收货地址
// @Tags 收货地址
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param req body AddressEditParams true "sku"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /address/edit [post]
func AddressEdit(c *gin.Context) {
	var req AddressEditParams
	v := app.BindJSON(c, &req)
	if !v {
		app.Error(c, errno.ErrBind)
		return
	}

	addr := map[string]interface{}{
		"name":       req.Name,
		"phone":      req.Phone,
		"province":   req.Province,
		"city":       req.City,
		"county":     req.County,
		"detail":     req.Detail,
		"is_default": req.IsDefault,
	}
	err := service.Svc.UserAddressEdit(c.Request.Context(), req.ID, app.GetUserID(c), addr)
	if err != nil {
		log.Warnf("[v1.user] edit address err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}

// AddressDel 删除收货地址
// @Summary 删除收货地址
// @Description 删除收货地址
// @Tags 收货地址
// @Accept json
// @Produce json
// @Param Token header string true "用户令牌"
// @Param id body int true "id"
// @Success 0 {string} json "{"code":0,"msg":"OK","data":{}}"
// @Router /address/del [post]
func AddressDel(c *gin.Context) {
	id := cast.ToInt(c.Query("id"))
	if id == 0 {
		app.Error(c, errno.ErrInvalidParam)
		return
	}

	err := service.Svc.DelUserAddress(c.Request.Context(), id, app.GetUserID(c))
	if err != nil {
		log.Warnf("[v1.cart] del err: %v", err)
		app.Error(c, errno.InternalServerError)
		return
	}
	app.SuccessNil(c)
}