package main

import (
	"fmt"
	"unicode"
)

func main()  {
	// 判断示例
	exampleIs()
	// 对应示例
	exampleSimpleFold()
	// 转换示例
	exampleTo()
}

func exampleIs()  {

	const mixed = "\b5Ὂg̀9! ℃ᾭG"
	for _, c := range mixed {
		fmt.Printf("For %q:\n", c)

		if unicode.IsControl(c) {
			fmt.Println("\tis control rune")
		}

		if unicode.IsDigit(c) {
			fmt.Println("\nis digit rune")
		}

		if unicode.IsGraphic(c) {
			fmt.Println("\tis graphic rune")
		}

		if unicode.IsLetter(c) {
			fmt.Println("\tis letter rune")
		}

		if unicode.IsLower(c) {
			fmt.Println("\tis lower case rune")
		}

		if unicode.IsUpper(c) {
			fmt.Println("\tis upper case rune")
		}

		if unicode.IsMark(c) {
			fmt.Println("\tis mark rune")
		}

		if unicode.IsNumber(c) {
			fmt.Println("\tis number rune")
		}

		if unicode.IsPrint(c) {
			fmt.Println("\tis printable rune")
		}

		if unicode.IsPunct(c) {
			fmt.Println("\tis punct rune")
		}

		if unicode.IsSpace(c) {
			fmt.Println("\tis space rune")
		}

		if unicode.IsSymbol(c) {
			fmt.Println("\tis symbol rune")
		}

		if unicode.IsTitle(c) {
			fmt.Println("\tis title case rune")
		}
	}

}

func exampleSimpleFold()  {

	fmt.Printf("%#U\n", unicode.SimpleFold('A'))
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A'))
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))

}

func exampleTo()  {

	const lcG = 'g'

	fmt.Printf("%#U\n", unicode.To(unicode.UpperCase, lcG))
	fmt.Println(unicode.ToUpper(lcG))

	fmt.Printf("%#U\n", unicode.To(unicode.LowerCase, lcG))
	fmt.Println(unicode.ToLower(lcG))

	fmt.Printf("%#U\n", unicode.To(unicode.TitleCase, lcG))
	fmt.Println(unicode.ToTitle(lcG))

}