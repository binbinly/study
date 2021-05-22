package utils

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	Env        = "API_ENV"
	EnvDev     = "dev"
	EnvTest    = "test"
	EnvProduct = "prod"
)

var CurEnv = EnvDev

func init() {
	CurEnv = strings.ToLower(os.Getenv(Env))
	CurEnv = strings.TrimSpace(CurEnv)

	if len(CurEnv) == 0 {
		CurEnv = EnvDev
	}
}

//RootDir 运行根目录
func RootDir() (rootPath string) {
	exePath := os.Args[0]
	rootPath = filepath.Dir(exePath)
	return rootPath
}

func IsProd() bool {
	return CurEnv == EnvProduct
}

func IsTest() bool {
	return CurEnv == EnvTest
}

func IsDev() bool {
	return CurEnv == EnvDev
}

func GetEnv() string {
	return CurEnv
}