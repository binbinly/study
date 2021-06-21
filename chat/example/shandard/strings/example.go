package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main()  {

	strDemo()

	s := reverse("Hello, World!")
	fmt.Println(s)
}

func strDemo()  {

	var a, b string

	strings.NewReader("Hello World")

	r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	fmt.Println(r.Replace("This is <b>HTML</b>!"))

	strings.Compare(a, b)

	strings.EqualFold(a, b)

	strings.HasPrefix(a, b)

	strings.HasSuffix(a, b)

	var c string

	strings.TrimSuffix(a, b)

	strings.TrimPrefix(a, b)

	strings.TrimSpace(a)

	strings.Trim(a, " ")

	strings.TrimFunc(a, unicode.IsLetter)

	strings.TrimLeft(a, "0123456789")

	strings.TrimLeftFunc(a, unicode.IsLetter)

	strings.TrimRight(a, "0123456789")

	strings.TrimRightFunc(a, unicode.IsLetter)

	strings.Fields(" foo base base    ")

	// 根据指定方法分割成切片
	strings.FieldsFunc("  foo1;bar2,baz3...", func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) // 以 不是字符或者数字 进行分割
	})

	strings.Contains(a, b)

	strings.ContainsAny("i like seafood.", "fÄo!")

	strings.ContainsRune("I like seafood.", 'f')

	strings.Count(a, b)

	strings.Index(a, b)

	strings.IndexByte(a, byte('k'))

	strings.IndexFunc("Hello,世界", func(r rune) bool {
		return unicode.Is(unicode.Han, r)
	})

	strings.IndexAny(a, "abc")

	strings.IndexRune("chicken", 'k')

	strings.LastIndex(a, b)

	strings.LastIndexByte(a, byte('k'))

	strings.LastIndexFunc(a, unicode.IsLetter)

	s := []string{a, b}
	strings.Join(s, ",")

	strings.Repeat(a, 2)

	strings.Replace(a, b, c, -1)

	strings.ReplaceAll(a, b, c)

	strings.Split(a, b)

	strings.SplitN(a, b, 2)

	strings.SplitAfter(a, b)

	strings.SplitAfterN(a, b, 2)

	strings.Title(a)

	strings.ToTitle(a)

	strings.ToTitleSpecial(unicode.SpecialCase{}, a)

	strings.ToUpper(a)

	strings.ToUpperSpecial(unicode.SpecialCase{}, a)

	strings.ToLower(a)

	strings.ToLowerSpecial(unicode.SpecialCase{}, a)

	strings.Map(func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return r
		} else {
			return 'a'
		}
	}, a)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r) - 1; i<len(r)/2; i, j = i+1, j-1 {
		r[i],r[j] = r[j], r[i]
	}
	return string(r)

}
