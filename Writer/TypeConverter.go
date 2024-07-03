package Writer

import (
	"GoFromCsToTypescript/Parser"
	"errors"
	"strings"
)
func ConvertCustomType(node Parser.INode) (string, error) {
	customType, ok := node.(Parser.CustomTypeNode)
	if !ok {
		return "", errors.New("not a custom type")
	}
	return customType.Type, nil
}


func ConvertSimpleType(node Parser.INode) (string, error) {
	customType, ok := node.(Parser.SimpleTypeNode)
	if !ok {
		return "", errors.New("not a simple type")
	}
  switch customType.Type {
	case Parser.Number:
		return "number", nil
	case Parser.String:
		return "string", nil
	case Parser.Boolean:
		return "boolean", nil
	}
	return "", errors.New("not a simple type")
}
func isArray(str string) bool {
	switch str {
	case "List","IEnumerable":
		return true
	}
	return false
}
func ConvertGenericType(node Parser.INode) (string, error) {
	customType, ok := node.(Parser.GenericTypeNode)
	if !ok {
		return "", errors.New("not a generic type")
	}
  var str strings.Builder
	if isArray(customType.ParentName) {
		str.WriteString("Array")
	}else {
		str.WriteString(customType.ParentName)
	}
	str.WriteString("<")
	for i, typ := range customType.ChildType {
		if i > 0 {
			str.WriteString(",")
		}
		f, err := getConvertType(typ)
		if err != nil {
			return "", err
		}
		strType, err := f(typ)
		if err != nil {
			return "", err
		}
		str.WriteString(strType)
	}
	str.WriteString(">")
	return str.String(), nil
}
