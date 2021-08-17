package controllers

import "github.com/astaxie/beego/toolbox"

func InitTask()  {
	task := toolbox.NewTask("GetAccountsAndWtChan", "0/3 * * * * *", GetAccountsAndWtChan)

	toolbox.AddTask("GetAccountsAndWtChan", task)
}

// 读取投注信息并写入通道
func GetAccountsAndWtChan() error {
	//获取投注信息

	return nil
}