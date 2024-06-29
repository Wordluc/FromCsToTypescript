package Parser

import (
	"GoFromCsToTypescript/Lexer"
	"errors"
)

type INode struct {
}

type Class struct {
	Fields []FieldNode
	Name   string
}

type FieldNode struct {
	Name     string
	Type     ITypeNode
	Nullable bool
}

type ITypeNode interface {
}

type SimpleTypeNode struct {
	Type Type
}

type CustomTypeNode struct {
	Type string
}

type GenericTypeNode struct { //parent<child1,child1>
	ParentName string
	ChildType  []ITypeNode
}

type Type int8

const (
	Unknown Type = iota
	Number  Type = iota
	String
	Boolean
)
const classIdentifier = "class"

func isVisibilitySetter(t string) bool {
	switch t {
	case "public", "private", "protected":
		return true
	default:
		return false
	}
}
func isBasicType(t string) Type {
	switch t {
	case "int":
		return Number
	case "string":
		return String
	case "bool":
		return Boolean
	default:
		return Unknown
	}
}

func Parse(input string) ([]INode, error) {
	_, err := Lexer.New(input)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func parseClass(l *Lexer.Lexer) (Class, error) {
	class := Class{}
	token := l.GetAndGoNext()
	if isVisibilitySetter(token.Val) {
		token = l.Pick()
	}
	if token.Val != classIdentifier {
		return class, errors.New("classes not supported yet")
	}
	l.Increse()
	token = l.GetAndGoNext()
	if token.Type != Lexer.Word {
		return class, errors.New("name not found")
	}
	class.Name = token.Val
	token = l.GetAndGoNext()
	if token.Type != Lexer.OpenCurly {
		return class, errors.New("{ not found")
	}
	parms, err := parseParamsList(l)
	if err != nil {
		return class, err
	}
	class.Fields = parms
	token = l.GetAndGoNext()
	if token.Type != Lexer.CloseCurly {
		return class, errors.New("} not found")
	}
	return class, nil
}
func parseParamsList(l *Lexer.Lexer) ([]FieldNode, error) {
	parms := []FieldNode{}
	token := l.Pick()
	for token.Type != Lexer.CloseCurly {
		parm, err := parseParam(l)
		if err != nil {
			return parms, err
		}
		parms = append(parms, parm)
		l.Increse()
		token = l.Pick()
	}
	return parms, nil
}
