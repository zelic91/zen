package funcs

import (
	"html/template"
	"strings"

	"github.com/zelic91/zen/config"
)

func FuncMap() template.FuncMap {
	return map[string]interface{}{
		"pluralize":      Pluralize,
		"singularize":    Singularize,
		"userProperties": UserProperties,
		"loop":           Loop,
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

func UserProperties() map[string]config.ModelProperty {
	return map[string]config.ModelProperty{
		"first_name":      {Type: "string"},
		"last_name":       {Type: "string"},
		"username":        {Type: "string", NotNull: true},
		"email":           {Type: "string"},
		"password_hashed": {Type: "string"},
		"password_salt":   {Type: "string"},
		"status":          {Type: "string"},
	}
}

func Loop(start int64, length int64) []int64 {
	ret := []int64{}
	for i := start; int64(len(ret)) < length; i++ {
		ret = append(ret, i)
	}
	return ret
}
