package funcs

import (
	"html/template"
	"strings"
)

func FuncMap() template.FuncMap {
	return map[string]interface{}{
		"pluralize":   Pluralize,
		"singularize": Singularize,
		"loop":        Loop,
	}
}

func Pluralize(input string) string {
	if strings.HasSuffix(input, "s") || strings.HasSuffix(input, "ies") || strings.HasSuffix(input, "es") {
		return input
	}

	if strings.HasSuffix(input, "y") {
		return input + "ies"
	}

	if strings.HasSuffix(input, "ss") {
		return input + "es"
	}
	return input + "s"
}

func Singularize(input string) string {
	if !strings.HasSuffix(input, "s") && !strings.HasSuffix(input, "ies") && !strings.HasSuffix(input, "sses") {
		return input
	}

	if strings.HasSuffix(input, "sses") {
		return strings.TrimSuffix(input, "es")
	}

	if strings.HasSuffix(input, "ies") {
		return strings.TrimSuffix(input, "ies") + "y"
	}

	return strings.TrimSuffix(input, "s")
}

func Loop(start int64, length int64) []int64 {
	ret := []int64{}
	for i := start; int64(len(ret)) < length; i++ {
		ret = append(ret, i)
	}
	return ret
}
