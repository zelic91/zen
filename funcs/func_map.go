package funcs

import (
	"html/template"
	"strings"

	"github.com/zelic91/zen/config"
)

func FuncMap() template.FuncMap {
	return map[string]interface{}{
		"pluralize":       Pluralize,
		"singularize":     Singularize,
		"userProperties":  UserProperties,
		"loop":            Loop,
		"sqlType":         SQLType,
		"hasReferences":   HasReferences,
		"references":      References,
		"isLastInMap":     IsLastInMap,
		"structFieldName": StructFieldName,
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

func SQLType(goType string) string {
	var ret string

	switch goType {
	case "string":
		ret = "VARCHAR"
	case "int64":
		ret = "BIGSERIAL"
	}

	return ret
}

func HasReferences(properties map[string]config.ModelProperty) bool {
	for _, value := range properties {
		if value.References != "" {
			return true
		}
	}

	return false
}

func References(properties map[string]config.ModelProperty) map[string]config.ModelProperty {
	out := map[string]config.ModelProperty{}
	for key, value := range properties {
		if value.References != "" {
			out[key] = value
		}
	}
	return out
}

func IsLastInMap(key string, m map[string]config.ModelProperty) bool {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys[len(keys)-1] == key
}

func StructFieldName(fieldName string) string {
	if strings.HasSuffix(fieldName, "Id") {
		return strings.TrimRight(fieldName, "Id") + "ID"
	}
	return fieldName
}
