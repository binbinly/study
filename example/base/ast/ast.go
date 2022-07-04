package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

// tpl 生成代码需要用到模板
const tpl = `
package {{.pkg}}

// messages get msg from const comment
var messages = map[{{.type}}]string{
	{{range $key, $value := .comments}}
	{{$key}}: "{{$value}}",{{end}}
}

// GetErrMsg get error msg
func GetErrMsg(code {{.type}}) string {
	if msg, ok := messages[code]; ok {
		return msg
	}
	return ""
}
`

// 根据常量注释，生成map映射关系
func genConstComment(file, outFile string)  {
	// 保存注释信息
	var comments = make(map[interface{}]string)
	fileSet := token.NewFileSet()
	f, err := parser.ParseFile(fileSet, file, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	var (
		cType interface{}
		once sync.Once
	)

	// 从ast获取注释
	nodes := ast.NewCommentMap(fileSet, f, f.Comments)
	for node := range nodes {
		if spec, ok := node.(*ast.ValueSpec); ok && len(spec.Names) == 1 {
			ident := spec.Names[0]
			once.Do(func() {
				if spec.Type != nil {
					if t, ok := spec.Type.(*ast.Ident); ok {
						cType = t
					}
				} else {
					for _, v := range spec.Values {
						if t, ok := v.(*ast.BasicLit); ok {
							cType = strings.ToLower(fmt.Sprintf("%v", t.Kind))
						}
					}
				}
			})

			if ident.Obj.Kind == ast.Con && spec.Comment != nil {
				comments[ident] = getComment(spec.Comment)
			}
		}
	}

	bs, err := gen(cType, comments)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(outFile, bs, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

// getComment 获取注释信息，来自AST标准库的summary方法
func getComment(group *ast.CommentGroup) string {
	var buf bytes.Buffer
	for _, comment := range group.List {
		text := strings.TrimSpace(strings.TrimLeft(comment.Text, "/"))
		buf.WriteString(text)
	}

	bs := buf.Bytes()
	for i, b := range bs {
		switch b {
		case '\t', '\n', '\r':
			bs[i] = ' '
		}
	}

	return string(bs)
}

func gen(cType interface{}, comments map[interface{}]string) ([]byte, error) {
	var buf = bytes.NewBufferString("")

	data := map[string]interface{}{
		"pkg": "example",
		"type": cType,
		"comments":comments,
	}

	t, err := template.New("").Parse(tpl)
	if err != nil {
		return nil, errors.Wrap(err, "template init err")
	}

	err = t.Execute(buf, data)
	if err != nil {
		return nil, errors.Wrap(err, "template data err")
	}
	
	return format.Source(buf.Bytes())
}