package csv

import (
	"os"
	"testing"
)

func Test_readCsv(t *testing.T) {
	file, _ := os.Open("test.csv")
	readCsv(file)
}

func Test_writeCsv(t *testing.T) {
	file, _ := os.OpenFile("test.csv", os.O_WRONLY|os.O_CREATE, os.ModePerm)
	writeCsv(file)
}
