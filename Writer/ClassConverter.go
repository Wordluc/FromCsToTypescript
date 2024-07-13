package Writer

import (
	"GoFromCsToTypescript/Parser"
	"errors"
	"strings"
)
func ConvertClass(class Parser.Class) (string, error) {
	var str strings.Builder
	str.WriteString("export interface ")
	if len(class.ExtendType) != 0 {
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
		str.WriteString("  ")
		if field.Nullable {
			str.WriteString(field.Name + "? : ")
		}else{
			str.WriteString(field.Name + " : ")
		}
		f, e = getConvertType(field.Type)
		if e != nil {
			return "", errors.New("Converter-Error:" + e.Error())
		}
		strType, e := f(field.Type)
		if e != nil {
			return "", errors.New("Converter-Error:" + e.Error())
		}
		str.WriteString(strType)
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
