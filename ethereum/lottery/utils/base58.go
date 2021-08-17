package utils

import (
	"bytes"
	"fmt"
	"math/big"
)

//base58(区块链)：去掉6个容易混淆的，去掉0，大写的O、大写的I、小写的L、/、+
var b58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

func encode(src string) string {
	// he : 104 101 --> 104*256 + 101 = 26725

	// 26725 / 58 =   15 16 17

	// 1.ascii码对应的值
	srcByte := []byte(src)

	//转十进制
	i := big.NewInt(0).SetBytes(srcByte)
	fmt.Println('i', i)

	var modSlice []byte
	//遍历取余
	for i.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0)
		i58 := big.NewInt(58)
		//取余
		i.DivMod(i, i58, mod)
		// 将余数添加到数组中
		modSlice = append(modSlice, b58[mod.Int64()])
	}
	fmt.Println("mod slice", modSlice)
	// 把0使用字节'1'代替
	for _, s := range srcByte {
		if s != 0 {
			break
		}
		modSlice = append(modSlice, byte('1'))
	}
	// 反转byte数组
	retModSlice := ReverseByteArr2(modSlice)
	fmt.Println("ret", retModSlice)
	return string(retModSlice)
}

func decode(src string) string {
	srcByte := []byte(src)

	// 这里得到的是十进制
	ret := big.NewInt(0)
	for _, b := range srcByte {
		i := bytes.IndexByte(b58, b)
		//乘
		ret.Mul(ret, big.NewInt(58))
		//加
		ret.Add(ret, big.NewInt(int64(i)))
	}
	return string(ret.Bytes())
}

// byte数组进行反转方式2
func ReverseByteArr2(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}
