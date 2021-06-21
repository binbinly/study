package main

import (
	"fmt"
	"log"
	"strconv"
)

func main()  {
	exampleBool()

	exampleFloat()

	exampleInt()

	exampleUint()

	exampleQuote()
}

func exampleBool()  {

	s := strconv.FormatBool(true)
	fmt.Printf("%T, %v\n", s, s)

	bl, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", bl, bl)

	b := []byte("bool:")
	b = strconv.AppendBool(b, true)
	fmt.Println(string(b))

}

func exampleFloat()  {

	v := 3.1415926535

	s32 := strconv.FormatFloat(v, 'f', -1, 32)
	fmt.Printf("%T, %v\n", s32, s32)

	s64 := strconv.FormatFloat(v, 'f', -1, 64)
	fmt.Printf("%T, %v\n", s64, s64)

	if s, err := strconv.ParseFloat(s32, 32); err== nil {
		s1 := float32(s)
		fmt.Printf("%T, %v\n", s1, s1)
	}

	if s, err := strconv.ParseFloat(s64, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

	b32 := []byte("float32:")
	b32 = strconv.AppendFloat(b32, v, 'f', -1, 32)
	fmt.Println(string(b32))

	b64 := []byte("float64:")
	b64 = strconv.AppendFloat(b64, v, 'f', -1, 64)
	fmt.Println(string(b64))
}

func exampleInt()  {

	v := 12

	s := strconv.Itoa(v)
	fmt.Printf("%T, %v\n", s, s)

	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", i, i)

	s = strconv.FormatInt(int64(v), 10)
	fmt.Printf("%T, %v\n", s, s)

	i64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", i64, i64)

	b10 := []byte("int (base 10)")
	b10 = strconv.AppendInt(b10, -42, 10)
	fmt.Println(string(b10))

	b16 := []byte("int (base 16):")
	b16 = strconv.AppendInt(b16, -42, 16)
	fmt.Println(string(b16))

}

func exampleUint()  {

	v := uint64(42)
	s := strconv.FormatUint(v, 10)
	fmt.Printf("%T, %v\n", s, s)

	u, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T, %v\n", u, u)

	b10 := []byte("uint (base 10):")
	b10 = strconv.AppendUint(b10, 42, 10)
	fmt.Println(string(b10))

}

func exampleQuote()  {

	v := "Fran & Freddie's Diner o"
	vr := 'o'

	ok := strconv.CanBackquote(v)
	fmt.Println(ok)

	s := strconv.Quote(v)
	fmt.Println(s)

	sa := strconv.QuoteToASCII(v)
	fmt.Println(sa)

	rok := strconv.IsPrint(vr)
	fmt.Println(rok)

	sr := strconv.QuoteRune(vr)
	fmt.Println(sr)

	sra := strconv.QuoteRuneToASCII(vr)
	fmt.Println(sra)

	b := strconv.AppendQuote([]byte("quota:"), `"Fran & Freddie's Diner"`)
	fmt.Println(string(b))

	ba := strconv.AppendQuoteToASCII([]byte("quota (ascii):"), v)
	fmt.Println(string(ba))

	br := strconv.AppendQuoteRune([]byte("rune:"), vr)
	fmt.Println(string(br))

	bra := strconv.AppendQuoteRuneToASCII([]byte("rune (ascii):"), vr)
	fmt.Println(string(bra))

}
