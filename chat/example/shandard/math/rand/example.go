package main

import (
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"
)

// rand包实现了伪随机数生成器
func main()  {

	example()

	exampleRand()
	exampleShuffle()

}

func example()  {
	rand.Seed(time.Now().UnixNano())

	fmt.Println(rand.Int())

	fmt.Println(rand.Intn(100))

	fmt.Println(rand.Int31())

	fmt.Println(rand.Int31n(100))

	fmt.Println(rand.Int63())

	fmt.Println(rand.Int63n(100))

	fmt.Println(rand.Uint32())

	fmt.Println(rand.Uint64())

	fmt.Println(rand.Float32())

	fmt.Println(rand.Float64())

	fmt.Println(rand.NormFloat64())

	fmt.Println(rand.ExpFloat64())

	fmt.Println(rand.Perm(13))
}

func exampleRand()  {

	source := rand.NewSource(99)

	r := rand.New(source)

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()

	show := func(name string, v1, v2, v3 interface{}) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}

	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())
	show("Int31", r.Int31(), r.Int31(), r.Int31())
	show("Int63", r.Int63(), r.Int63(), r.Int63())
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())
	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
}

func exampleShuffle()  {
	numbers := []byte("12345")
	letters := []byte("ABCDR")

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		letters[i], letters[j] = letters[j], letters[i]
	})
	for i := range numbers {
		fmt.Printf("%c: %c\n", letters[i], numbers[i])
	}
}