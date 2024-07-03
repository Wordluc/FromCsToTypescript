package Writer

import (
	"GoFromCsToTypescript/Parser"
	"errors"
	"strings"
)
func ConvertClass(node Parser.INode) (string, error) {
  class,ok:=node.(Parser.Class)
	if !ok {
		return "",errors.New("not a class")
	}
	var str  strings.Builder
	str.WriteString("export interface ")
	str.WriteString(class.Name+" {\n")
	var e error
	var f IConvert
	for _,field:=range class.Fields {
		str.WriteString(field.Name+" : ")
		f,e=getConvertType(field.Type)
		if e != nil {
			return "",errors.New("Converter-Error:"+ e.Error())
		}
		strType,e:=f(field.Type)
		if e != nil {
			return "",errors.New("Converter-Error:"+ e.Error())
		}
		str.WriteString(strType)
		if field.Nullable {
			str.WriteString("?")
		}
		str.WriteString(";\n")
	}
	str.WriteString("}")
	return str.String(), nil
}
