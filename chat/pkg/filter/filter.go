package tools

import (
	"github.com/importcjj/sensitive"
)
// 敏感词过滤
var filter *sensitive.Filter

//加载敏感词库
func LoadFilter(path string) error {
	filter = sensitive.New()
	err := filter.LoadWordDict(path)
	if err != nil {
		return err
	}
	filter.UpdateNoisePattern(`x`)
	filter.UpdateNoisePattern(` `)
	return nil
}

//替换文本
func Replace(content string) string {
	return filter.Replace(content, '*')
}
