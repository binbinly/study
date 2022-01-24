package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"unicode/utf8"
)

//see: https://www.topgoer.cn/docs/golang/chapter21-1
func main()  {
	dynamicMultiArray()
	dynamicMultiArray2()
	mapIn()
	stringEdit()
	stringRune()
	ascii()
	stringUtf8()
	stringLen()
	stringNorm()
	rangeMap()
	switchChar()
	increment()
	bit()
	encode()
	goroutine()
	waitGroup()
	receiver()
	channelClose()
	channelNil()
	channelNil2()
}

// 数组使用值拷贝传参
func array() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) // [7 2 3]
	}(x)
	fmt.Println(x) // [1 2 3]    // 并不是你以为的 [7 2 3]
}

// 会修改 slice 的底层 array，从而修改 slice
func slice() {
	x := []int{1, 2, 3}
	func(arr []int) {
		arr[0] = 7
		fmt.Println(x)    // [7 2 3]
	}(x)
	fmt.Println(x)    // [7 2 3]
}

// 使用各自独立的 6 个 slice 来创建 [2][3] 的动态多维数组
func dynamicMultiArray() {
	x := 2
	y := 4

	table := make([][]int, x)
	for i  := range table {
		table[i] = make([]int, y)
	}
}

//1.使用“共享底层数组”的切片
//创建一个存放原始数据的容器 slice
//创建其他的 slice
//切割原始 slice 来初始化其他的 slice
func dynamicMultiArray2() {
	h, w := 2, 4
	raw := make([]int, h*w)

	for i := range raw {
		raw[i] = i
	}

	// 初始化原始 slice
	fmt.Println(raw, &raw[4])    // [0 1 2 3 4 5 6 7] 0xc420012120

	table := make([][]int, h)
	for i := range table {

		// 等间距切割原始 slice，创建动态多维数组 table
		// 0: raw[0*4: 0*4 + 4]
		// 1: raw[1*4: 1*4 + 4]
		table[i] = raw[i*w : i*w + w]
	}

	fmt.Println(table, &table[1][0])    // [[0 1 2 3] [4 5 6 7]] 0xc420012120
}

// 检查 key 是否存在可以用 map 直接访问，检查返回的第二个参数即可：
func mapIn() {
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	if _, ok := x["two"]; !ok {
		fmt.Println("key two is no entry")
	}
}

// 尝试使用索引遍历字符串，来更新字符串中的个别字符，是不允许的。
//string 类型的值是只读的二进制 byte slice，如果真要修改字符串中的字符，将 string 转为 []byte 修改后，再转为 string 即可：
func stringEdit() {
	x := "text"
	xBytes := []byte(x)
	xBytes[0] = 'T'    // 注意此时的 T 是 rune 类型
	x = string(xBytes)
	fmt.Println(x)    // Text
}

//注意： 上边的示例并不是更新字符串的正确姿势，因为一个 UTF8 编码的字符可能会占多个字节，比如汉字就需要 3~4个字节来存储，此时更新其中的一个字节是错误的。
//更新字串的正确姿势：将 string 转为 rune slice（此时 1 个 rune 可能占多个 byte），直接更新 rune 中的字符
func stringRune() {
	x := "text"
	xRunes := []rune(x)
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x)    // 我ext
}

//对字符串用索引访问返回的不是字符，而是一个 byte 值。
func ascii() {
	x := "ascii"
	fmt.Println(x[0])        // 97
	fmt.Printf("%T\n", x[0])// uint8
}

//string 的值不必是 UTF8 文本，可以包含任意的值。只有字符串是文字字面值时才是 UTF8 文本，字串可以通过转义来包含其他数据。
//判断字符串是否是 UTF8 文本，可使用 “unicode/utf8” 包中的 ValidString() 函数：
func stringUtf8() {
	str1 := "ABC"
	fmt.Println(utf8.ValidString(str1))    // true

	str2 := "A\xfeC"
	fmt.Println(utf8.ValidString(str2))    // false

	str3 := "A\\xfeC"
	fmt.Println(utf8.ValidString(str3))    // true    // 把转义字符转义成字面值
}

//Go 的内建函数 len() 返回的是字符串的 byte 数量，而不是像 Python 中那样是计算 Unicode 字符数。
//如果要得到字符串的字符数，可使用 “unicode/utf8” 包中的 RuneCountInString(str string) (n int)
func stringLen() {
	char := "♥"
	fmt.Println(utf8.RuneCountInString(char))    // 1
	char = "é"
	fmt.Println(len(char))    // 3
	fmt.Println(utf8.RuneCountInString(char))    // 2
	fmt.Println("cafe\u0301")    // café    // 法文的 cafe，实际上是两个 rune 的组合
}

//range 得到的索引是字符值（Unicode point / rune）第一个字节的位置，与其他编程语言不同，这个索引并不直接是字符在字符串中的位置。
//注意一个字符可能占多个 rune，比如法文单词 café 中的 é。操作特殊字符可使用norm 包。
//for range 迭代会尝试将 string 翻译为 UTF8 文本，对任何无效的码点都直接使用 0XFFFD rune（�）UNicode 替代字符来表示。如果 string 中有任何非 UTF8 的数据，应将 string 保存为 byte slice 再进行操作。
func stringNorm() {
	data := "A\xfe\x02\xff\x04"

	for _, v := range []byte(data) {
		fmt.Printf("%#x ", v)    // 0x41 0xfe 0x2 0xff 0x4    // 正确
	}
}

//如果你希望以特定的顺序（如按 key 排序）来迭代 map，要注意每次迭代都可能产生不一样的结果。
//Go 的运行时是有意打乱迭代顺序的，所以你得到的迭代结果可能不一致。但也并不总会打乱，得到连续相同的 5 个迭代结果也是可能的
func rangeMap() {
	m := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func switchChar() {
	isSpace := func(char byte) bool {
		switch char {
		case ' ', '\t':
			return true
		}
		return false
	}
	fmt.Println(isSpace('\t'))    // true
	fmt.Println(isSpace(' '))    // true
}

// 很多编程语言都自带前置后置的 ++、– 运算。但 Go 特立独行，去掉了前置操作，同时 ++、— 只作为运算符而非表达式。
func increment() {
	data := []int{1, 2, 3}
	i := 0
	i++
	fmt.Println(data[i])    // 2
}

//很多编程语言使用 ~ 作为一元按位取反（NOT）操作符，Go 重用 ^ XOR 操作符来按位取反：
func bit() {
	var d uint8 = 2
	fmt.Printf("%08b\n", d)        // 00000010
	fmt.Printf("%08b\n", ^d)    // 11111101

	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n", a)
	fmt.Printf("%08b [B]\n", b)

	fmt.Printf("%08b (NOT B)\n", ^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)

	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))
}

//以小写字母开头的字段成员是无法被外部直接访问的，所以 struct 在进行 json、xml、gob 等格式的 encode 操作时，这些私有字段会被忽略，导出时得到零值：
func encode() {
	type MyData struct {
		One int
		two string
	}
	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)    // main.MyData{One:1, two:"two"}

	encoded, _ := json.Marshal(in)
	fmt.Println(string(encoded))    // {"One":1}    // 私有字段 two 被忽略了

	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out)     // main.MyData{One:1, two:""}
}

//程序默认不等所有 goroutine 都执行完才退出，这点需要特别注意：
func goroutine() {
	workerCount := 2
	for i := 0; i < workerCount; i++ {
		go doIt(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("all done!")
}

func doIt(workerID int) {
	fmt.Printf("[%v] is running\n", workerID)
	time.Sleep(3 * time.Second)        // 模拟 goroutine 正在执行
	fmt.Printf("[%v] is done\n", workerID)
}

//常用解决办法：使用 “WaitGroup” 变量，它会让主程序等待所有 goroutine 执行完毕再退出。
//如果你的 goroutine 要做消息的循环处理等耗时操作，可以向它们发送一条 kill 消息来关闭它们。或直接关闭一个它们都等待接收数据的 channel：
// 等待所有 goroutine 执行完毕
// 使用传址方式为 WaitGroup 变量传参
// 使用 channel 关闭 goroutine
func waitGroup() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	ch := make(chan interface{})

	workerCount := 2
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doIt1(i, ch, done, &wg)    // wg 传指针，doIt() 内部会改变 wg 的值
	}

	for i := 0; i < workerCount; i++ {    // 向 ch 中发送数据，关闭 goroutine
		ch <- i
	}

	close(done)
	wg.Wait()
	close(ch)
	fmt.Println("all done!")
}

func doIt1(workerID int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerID)
	defer wg.Done()
	for {
		select {
		case m := <-ch:
			fmt.Printf("[%v] m => %v\n", workerID, m)
		case <-done:
			fmt.Printf("[%v] is done\n", workerID)
			return
		}
	}
}

//只有在数据被 receiver 处理时，sender 才会阻塞。因运行环境而异，在 sender 发送完数据后，receiver 的 goroutine 可能没有足够的时间处理下一个数据。如：
func receiver() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("Processed:", m)
			time.Sleep(1 * time.Second)    // 模拟需要长时间运行的操作
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" // 不会被接收处理
}

//从已关闭的 channel 接收数据是安全的：
//接收状态值 ok 是 false 时表明 channel 中已没有数据可以接收了。
//类似的，从有缓冲的 channel 中接收数据，缓存的数据获取完再没有数据可取时，状态值也是 false
//向已关闭的 channel 中发送数据会造成 panic：
func channelClose() {
	ch := make(chan int)
	done := make(chan struct{})

	for i := 0; i < 3; i++ {
		go func(idx int) {
			select {
			case ch <- (idx + 1) * 2:
				fmt.Println(idx, "Send result")
			case <-done:
				fmt.Println(idx, "Exiting")
			}
		}(i)
	}

	fmt.Println("Result: ", <-ch)
	close(done)
	time.Sleep(3 * time.Second)
}

//在一个值为 nil 的 channel 上发送和接收数据将永久阻塞：
func channelNil() {
	var ch chan int // 未初始化，值为 nil
	for i := 0; i < 3; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}

	//fmt.Println("Result: ", <-ch)
	time.Sleep(2 * time.Second)
}

func channelNil2() {
	inCh := make(chan int)
	outCh := make(chan int)

	go func() {
		var in <-chan int = inCh
		var out chan<- int
		var val int

		for {
			select {
			case out <- val:
				println("--------")
				out = nil
				in = inCh
			case val = <-in:
				println("++++++++++")
				out = outCh
				in = nil
			}
		}
	}()

	go func() {
		for r := range outCh {
			fmt.Println("Result: ", r)
		}
	}()

	time.Sleep(0)
	inCh <- 1
	inCh <- 2
	time.Sleep(3 * time.Second)
}