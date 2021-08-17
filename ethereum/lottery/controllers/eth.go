package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"lottery/utils"
	"math/big"
)

//客户端连接
var connects = make(map[*websocket.Conn]bool)

var accountsChan = make(chan map[string]interface{})

// 升级为ws
var upgrader = websocket.Upgrader{}

var contract *utils.ContractConnect

type EthController struct {
	beego.Controller
}

func (e *EthController) Index() {
	total, err := contract.GetPeriod()
	if err != nil {
		fmt.Println("err", err)
	}
	e.Data["total"] = total
	e.TplName = "index.html"
}

// 获取投注信息
func (e *EthController) PostTouZhu() {
	username := e.GetString("username")
	pwd := e.GetString("pwd")
	num1 := e.GetString("num1")
	num2 := e.GetString("num2")
	num3 := e.GetString("num3")
	num4 := e.GetString("num4")
	num5 := e.GetString("num5")

	nums := num1 + " " + num2 + " " + num3 + " " + num4 + " " + num5
	fmt.Println(nums, username, pwd)
	// 提交到以太坊合约
	//tx, err := contract.Bet(strings.ToLower(username), pwd, nums)
	//if err == nil {
	//	ret := map[string]interface{}{
	//		"code": 200,
	//		"msg":  "投注成功",
	//		"data": tx.Hash(),
	//	}
	//	e.Data["json"] = ret
	//	e.ServeJSON()
	//}

	ret := map[string]interface{}{
		"code": 400,
		"msg":  "投注失败",
	}
	e.Data["json"] = ret
	e.ServeJSON()
}

// 实时获取投注账户及投注号码,返回json
func (e *EthController) GetBetAccounts() {
	ws, err := upgrader.Upgrade(e.Ctx.ResponseWriter, e.Ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
	}
	// 放到clients中
	connects[ws] = true
}

// 查询页面
func (e *EthController) Search() {
	e.TplName = "search.html"
}

// 获取查询数据
func (e *EthController) GetAcountInfo() {

	addr := e.GetString("addr")

	// 把addr传到智能合约，返回改用户的账户余额和投注号码，
	// 替换下面的balance、nums
	balance, err := contract.GetBalance(addr)

	if err != nil {
		balance = 0
	}

	_, nums, err := contract.GetBetAccounts(big.NewInt(1))
	if err != nil {
		fmt.Println("err", err)
	}

	ret_map := map[string]interface{}{}
	if addr == "" {
		ret_map = map[string]interface{}{
			"addr":    addr,
			"balance": 0,
			"nums":    [][]int{},
		}
	} else {
		ret_map = map[string]interface{}{
			"addr":    addr,
			"balance": balance,
			"nums":    nums,
		}
	}

	e.Data["account_info"] = ret_map
	e.TplName = "search.html"
}

// 开奖页面
func (e *EthController) KaiJiang() {
	total, err := contract.GetPeriod()
	if err != nil {
		total = big.NewInt(0)
	}
	e.Data["total"] = total.Int64()
	e.TplName = "kaijiang.html"

}

// 开奖功能
func (e *EthController) DoKaiJiang() {
	err := contract.Draw(nil)

	if err != nil {
		e.Data["ret_num"] = []*big.Int{}
		e.Data["admin"] = "ee7ec7e6b303601ff62e947664710c552c3f7942"
		e.Data["total"] = 0
		e.Data["money"] = 0
	} else {
		e.Data["ret_num"] = []*big.Int{}
		e.Data["admin"] = "ee7ec7e6b303601ff62e947664710c552c3f7942"
		e.Data["total"] = 0
		e.Data["money"] = 0
	}
	e.TplName = "kaijiang.html"
}

// 智能合约页面
func (e *EthController) Contract() {
	e.TplName = "contract.html"
}
