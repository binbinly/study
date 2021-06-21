package main

import (
	"fmt"
	"log"
	"net/url"
)

// url包解析URL并实现了查询的逸码
func main()  {
	example()

	exampleURL()

	exampleQuery()

	exampleUser()
}

func example()  {
	s := "https://www.baidu.com?test=1&name=测试"

	se := url.QueryEscape(s)
	fmt.Println("urlEncode:", se)

	sd, err := url.QueryUnescape(se)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("urlDecode:", sd)
}

func exampleURL()  {
	u, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme:", u.Scheme)

	fmt.Println("Host:", u.Host)

	fmt.Println("Opaque:", u.Opaque)

	fmt.Println("User:", u.User)

	fmt.Println("Path:", u.Path)

	fmt.Println("RawQuery", u.RawQuery)

	fmt.Println("Fragment:", u.Fragment)

	u.Scheme = "https"
	u.Host = "google.com"

	fmt.Println("IsAbs:", u.IsAbs())

	fmt.Println("Hostname:", u.Hostname())

	fmt.Println("Port", u.Port())

	fmt.Println("String:", u.String())

	fmt.Println("EscapePath", u.EscapedPath())

	fmt.Println("RequestURI", u.RequestURI())

	text, err := u.MarshalBinary()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MarshalBinary:%s\n", string(text))

	if err = u.UnmarshalBinary(text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("UnmarshalBinary:ok")

	q := u.Query()
	q.Set("q", "golang")
	q.Add("say", "hello")

	fmt.Println(q.Get("say"))

	fmt.Println(q.Encode())

	q.Del("say")

	fmt.Println(u)

	// 解析一个相对路径的url
	uri, err := url.Parse("../../..//search?q=test")
	if err != nil {
		log.Fatal(err)
	}
	uriNew := u.ResolveReference(uri)
	fmt.Println(uriNew.String())
}

func exampleQuery()  {
	m, err := url.ParseQuery("x=1&y=2&y=3;z")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m)

	params := url.Values{}
	params.Set("say", "hello")
	params.Set("test", "123")
	params.Set("q", "Go")
	fmt.Println(params.Encode())
}

func exampleUser()  {
	u := url.User("gopher")

	fmt.Println(u.String())

	fmt.Println(u.Username())

	fmt.Println(u.Password())

	up := url.UserPassword("gopher", "123123")
	fmt.Println(up.String())
	fmt.Println(up.Password())

}