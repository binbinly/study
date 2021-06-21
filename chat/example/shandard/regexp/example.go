package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main()  {
	exampleMatch()

	exampleQuoteMeta()

	exampleCompile()

	exampleCompilePOSTIX()

	exampleSubMatch()
}

func exampleMatch()  {

	data := "Hello World[test]"
	patterm := `\d*`

	ok, err := regexp.Match(patterm, []byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)

	ok, err = regexp.MatchReader(patterm, strings.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)

	ok, err = regexp.MatchString(patterm, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ok)
}

func exampleQuoteMeta()  {

	str := regexp.QuoteMeta(`Escaping symbols like: .+*?()|[]{}^$`)
	fmt.Println(str)

}

func exampleCompile()  {

	data := "Hello foo!"

	r, err := regexp.Compile(`foo.?`)
	if err != nil {
		log.Fatal(err)
	}

	r = regexp.MustCompile(`foo(.?)`)
	fmt.Println(r.MatchString(data))

	fmt.Println(r.Match([]byte(data)))

	fmt.Println(r.MatchReader(strings.NewReader(data)))

	fmt.Println("正则表达式：", r.String())

	fmt.Printf("%s\n", r.Find([]byte(data)))

	fmt.Println(r.FindIndex([]byte(data)))

	fmt.Println(r.FindString(data))

	fmt.Println(r.FindStringIndex(data))

	fmt.Printf("%q\n", r.FindAll([]byte(data), -1))

	fmt.Println(r.FindAllIndex([]byte(data), -1))

	fmt.Println(r.FindAllString(data, -1))

	fmt.Println(r.FindAllStringIndex(data, -1))

	fmt.Println(r.FindReaderIndex(strings.NewReader(data)))

	fmt.Println(r.Split(data, -1))

	fmt.Println(r.LiteralPrefix())

	fmt.Printf("%s\n", r.ReplaceAll([]byte(data), []byte("World!")))
	fmt.Printf("%s\n", r.ReplaceAllString(data, "World!"))

	fmt.Printf("%s\n", r.ReplaceAllLiteral([]byte(data), []byte("World!")))
	fmt.Printf("%s\n", r.ReplaceAllLiteralString(data, "World!"))

	fmt.Printf("%s\n", r.ReplaceAllFunc([]byte(data), func(b []byte) []byte { return []byte("World!") }))
	fmt.Printf("%s\n", r.ReplaceAllStringFunc(data, func(s string) string { return "World!" }))

	r.Longest()
}

func exampleCompilePOSTIX()  {

	data := "Hello World!"

	r, err := regexp.CompilePOSIX(`foo.?`)
	if err != nil {
		log.Fatal(err)
	}

	r = regexp.MustCompilePOSIX(`foo.?`)

	fmt.Println(r.MatchString(data))
}

func exampleSubMatch()  {

	data := "seafood fool"

	r := regexp.MustCompile(`foo(.?)`)

	fmt.Println(r.NumSubexp())

	fmt.Println(r.SubexpNames())

	fmt.Printf("%s\n", r.FindSubmatch([]byte(data)))

	fmt.Println(r.FindSubmatchIndex([]byte(data)))

	fmt.Println(r.FindStringSubmatch(data))

	fmt.Println(r.FindStringSubmatchIndex(data))

	fmt.Printf("%q\n", r.FindAllSubmatch([]byte(data), -1))

	fmt.Println(r.FindAllSubmatchIndex([]byte(data), -1))

	fmt.Println(r.FindAllStringSubmatch(data, -1))

	fmt.Println(r.FindAllStringSubmatchIndex(data, -1))

	fmt.Println(r.FindReaderSubmatchIndex(strings.NewReader(data)))
}

