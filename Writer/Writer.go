package Writer

import (
	"GoFromCsToTypescript/Lexer"
	"GoFromCsToTypescript/Parser"
	"errors"
)

func Convert(str string) (string, error) {
	l, e := Lexer.New(str)
	if e != nil {
		return "",errors.New("Lexer-Error:"+ e.Error())
	}
	node,e:=Parser.Parse(l)
	if e != nil {
		return "",errors.New("Parser-Error:"+ e.Error())
	}
	return ConvertClass(node)
}

type IConvert func (Parser.INode) (string, error)
func getConvertType(node Parser.INode) (IConvert, error) {
	switch node.(type) {
	case Parser.CustomTypeNode:
       return ConvertCustomType,nil
	case Parser.SimpleTypeNode:
			 return ConvertSimpleType,nil
	case Parser.GenericTypeNode:
      return ConvertGenericType,nil
	}
	return nil,errors.New("node not managed")
}
