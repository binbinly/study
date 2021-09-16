package app

import (
	"os"
	"path/filepath"
)

const (
	//EnvDev 开发环境
	EnvDev  = "dev"
	//EnvTest 测试环境
	EnvTest = "test"
	//EnvProd 生成环境
	EnvProd = "prod"
)

var (
	curEnv = EnvDev	//当前环境
	appName = "app"	//应用名称
)

//Init 初始化
func Init(env, name string) {
	curEnv = env
	appName = name
}

//RootDir 运行根目录
func RootDir() (rootPath string) {
	exePath := os.Args[0]
	rootPath = filepath.Dir(exePath)
	return rootPath
}

//IsProd 是否生成环境
func IsProd() bool {
	return curEnv == EnvProd
}

//IsTest 是否测试环境
func IsTest() bool {
	return curEnv == EnvTest
}

//IsDev 是否为开发环境
func IsDev() bool {
	return curEnv == EnvDev
}

//GetEnv 获取应用环境
func GetEnv() string {
	return curEnv
}

//GetAppName 后去应用名
func GetAppName() string {
	return appName
}