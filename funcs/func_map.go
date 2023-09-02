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

func singularize(input string) string {
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

func loop(start int64, length int64) []int64 {
	ret := []int64{}
	for i := start; int64(len(ret)) < length; i++ {
		ret = append(ret, i)
	}
	return ret
}
