package main

import "math/big"

// big包实现了大数字的多精度计算
func main()  {
	// 创建一个值未x的big.int
	i := big.NewInt(10)

	i.Int64()

	i.Uint64()

	i.Bytes()

	_ = i.String()

	i.BitLen()

	i.Bits()

	i.Bit(1)

	i.SetInt64(11)

	i.SetUint64(11)

	i.SetBytes([]byte("11"))

	i.SetString("11", big.MaxBase)
}

