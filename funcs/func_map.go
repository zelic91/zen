package funcs

import (
	"html/template"
	"strings"
)

func FuncMap() template.FuncMap {
	return map[string]interface{}{
		"pluralize":   pluralize,
		"singularize": singularize,
		"loop":        loop,
	}
}

func pluralize(input string) string {
	if strings.HasSuffix(input, "s") {
		return input
	}
	return input + "s"
}

func singularize(input string) string {
	if !strings.HasSuffix(input, "s") {
		return input
	}
	return strings.TrimSuffix(input, "s")
}

func loop(start int64, length int64) []int64 {
	ret := []int64{}
	for i := start; int64(len(ret)) < length; i++ {
		ret = append(ret, i)
	}
	return ret
}
