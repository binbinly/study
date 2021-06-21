package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
)

func main()  {
	example()
	example2()
}

type Hello struct {
	XMLName xml.Name `xml:"xml_name"`
	Id int `xml:"id"`
	FullName string `xml:"full_name"`
	FirstName string `xml:"first_name"`
	LastName string `xml:"last_name"`
	Sex int `xml:"sex"`
	Comment string `xml:"comment"`
}

func example()  {

	hello := Hello{
		XMLName:   xml.Name{},
		Id:        12,
		FullName:  "zc",
		FirstName: "z",
		LastName:  "c",
		Sex:       1,
		Comment:   "test notes",
	}

	b, err := xml.Marshal(hello)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))

	origin := Hello{}
	if err = xml.Unmarshal(b, &origin); err != nil {
		log.Fatal(err)
	}
	fmt.Println(origin)

	bi, err := xml.MarshalIndent(hello, "\r", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bi))
}

func example2()  {

	var buf bytes.Buffer

	hello := Hello{
		XMLName:   xml.Name{},
		Id:        12,
		FullName:  "zc",
		FirstName: "z",
		LastName:  "c",
		Sex:       1,
		Comment:   "test notes",
	}

	e := xml.NewEncoder(&buf)

	e.Indent("\r", "  ")

	err := e.EncodeElement(hello, xml.StartElement{
		Name: xml.Name{
			Space: "test",
			Local: "newXml",
		},
		Attr: nil,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf.Bytes()))
}