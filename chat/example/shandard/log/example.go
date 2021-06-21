package main

import (
	"log"
	"os"
)

const LogFile = "example/testdata/testLog"

func main()  {
	file, err := os.Create(LogFile)
	if err != nil {
		log.Fatal(err)
	}

	l := log.New(file, "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)

	l.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)

	l.SetPrefix("")

	l.SetOutput(file)

	l.Output(2, "hello world")

	l.Prefix()

	l.Flags()

	l.Println("ok")
}