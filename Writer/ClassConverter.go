package Writer

import (
	"GoFromCsToTypescript/Parser"
	"errors"
	"strings"
)

func ConvertClass(node Parser.INode) (string, error) {
	class, ok := node.(Parser.Class)
	if !ok {
		return "", errors.New("not a class")
	}
	var str strings.Builder
	str.WriteString("export interface ")
	if class.ExtendType != nil {
		strs, e := ConverExtends(class.ExtendType)
		if e != nil {
			return "", errors.New("Converter-Error:" + e.Error())
		}
		str.WriteString(class.Name)
		str.WriteString(strs + " {\n")
	} else {
		str.WriteString(class.Name + " {\n")
	}
	var e error
	var f IConvert
	for _, field := range class.Fields {
		str.WriteString(field.Name + " : ")
		f, e = getConvertType(field.Type)
		if e != nil {
			return "", errors.New("Converter-Error:" + e.Error())
		}
		strType, e := f(field.Type)
		if e != nil {
			return "", errors.New("Converter-Error:" + e.Error())
		}
		str.WriteString(strType)
		if field.Nullable {
			str.WriteString(" | null")
		}
		str.WriteString(";\n")
	}
	str.WriteString("}")
	return str.String(), nil
}
func ConverExtends(nodes []Parser.INode) (string, error) {
	var strs strings.Builder
	strs.WriteString(" extends ")
	for i, node := range nodes {
		f, e := getConvertType(node)
		str, e := f(node)
		if e != nil {
			return "", errors.New("Converter-Error:" + e.Error())
		}
		strs.WriteString(str)
		if i < len(nodes)-1 {
			strs.WriteString(", ")
		}
	}
	return strs.String(), nil
}
