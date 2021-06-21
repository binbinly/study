package main

import "unicode/utf8"

func main()  {

	var b = []byte("Hello World")

	utf8.Valid(b)

	utf8.ValidRune('H')

	utf8.ValidString(string(b))

	utf8.EncodeRune(b, 'H')

	utf8.DecodeRune(b)

	utf8.DecodeRuneInString(string(b))

	utf8.DecodeLastRune(b)

	utf8.DecodeLastRuneInString(string(b))

	utf8.FullRune(b)

	utf8.FullRuneInString(string(b))

	utf8.RuneCount(b)

	utf8.RuneCountInString(string(b))

	utf8.RuneLen('世')

	utf8.RuneStart('世')

}
