package Parser

import (
	"GoFromCsToTypescript/Lexer"
	"errors"
)

type INode interface {
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
const recordIdentifier = "record"

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

func Parse(input string) (Class, error) {
	l, err := Lexer.New(input)
	if err != nil {
		return Class{}, err
	}
	token := l.PickNext()
	if token.Val == classIdentifier {
		return parseClass(l)
	} else if token.Val == recordIdentifier {
		return parseRecord(l)
	}
	return Class{}, errors.New("identifier not managed")
}
func parseRecord(l *Lexer.Lexer) (Class, error) {
	class := Class{}
	token := l.GetAndGoNext()
	if isVisibilitySetter(token.Val) {
	}
	l.Increse()
	token = l.GetAndGoNext()
	if token.Type != Lexer.Word {
		return class, errors.New("name not found")
	}
	class.Name = token.Val
	token = l.GetAndGoNext()
	if token.Type != Lexer.OpenCircle {
		return class, errors.New("( not found")
	}
	token = l.Pick()
	parms := []FieldNode{}
	for token.Type != Lexer.CloseCircle && token.Type != Lexer.CloseCurly && token.Type != Lexer.Semicolon {
		parm, err := parseParam(l)
		if err != nil {
			return class, err
		}
		parms = append(parms, parm)
		l.Increse()
		if l.Pick().Type == Lexer.OpenCurly {
			l.Increse()
		}

		token = l.Pick()

	}
	class.Fields = parms
	token = l.GetAndGoNext()
	return class, nil
}
func parseClass(l *Lexer.Lexer) (Class, error) {
	class := Class{}
	token := l.GetAndGoNext()
	if isVisibilitySetter(token.Val) {
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
