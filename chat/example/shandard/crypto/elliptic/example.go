package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// 实现了几条覆盖素数有限域的标准椭圆曲线
func main()  {
	//返回一个实现P-224曲线
	elliptic.P224()

	elliptic.P256()

	elliptic.P384()

	curve := elliptic.P521()

	priv, x, y, err := elliptic.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(priv))
	fmt.Println(x)
	fmt.Println(y)

	d := elliptic.Marshal(curve, x, y)

	elliptic.Unmarshal(curve, d)
}