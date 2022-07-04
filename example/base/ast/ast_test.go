package ast

import "testing"

func Test_genConstComment(t *testing.T) {
	genConstComment("example/code.go", "example/code_msg.go")
}
