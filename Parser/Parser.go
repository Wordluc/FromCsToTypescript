package Parser

//gestire gli extends,implements
import (
	"GoFromCsToTypescript/Lexer"
	"errors"
	"regexp"
)

type INode interface {
}

type Class struct {
	Fields     []FieldNode
	Name       string
	ExtendType []INode
}

type FieldNode struct {
	Name     string
	Type     INode
	Nullable bool
}

type SimpleTypeNode struct {
	Type Type
}

type CustomTypeNode struct {
	Type string
}

type GenericTypeNode struct { //parent<child1,child1>
	ParentName string
	ChildType  []INode
}

type Type int8

const (
	Unknown Type = iota
	Number  Type = iota
	String
	Date
	Boolean
)
const classIdentifier = "class"
const abstractIdentifier = "abstract"
const recordIdentifier = "record"

func isModifier(t string) bool {
	switch t {
	case "public", "private", "protected","abstract":
		return true
	default:
		return false
	}
}
func isBasicType(t string) Type {
	if isNumber, _ := regexp.Match(`[i|I]nt\d*|[f|F]loat\d*`, []byte(t)); isNumber {
		return Number
	}
	switch t {
	case "short", "Byte", "byte":
		return Number
	case "string":
		return String
	case "bool":
		return Boolean
	case "DateTime":
		return Date
	default:
		return Unknown
	}
}
func ParseStr(str string) ([]Class, error) {
	l, e := Lexer.New(str)
	if e != nil {
		return nil, e
	}
	return Parse(l)
}
func Parse(l *Lexer.Lexer) ([]Class, error) {
	classes := []Class{}
	for {
		class := Class{}
		var err error
		class, err = parse(l)
		classes = append(classes, class)
		if l.Pick().Type == Lexer.EOF {
			return classes, nil
		}
		if err != nil{
			return nil,err
		}
	}
}

func parse(l *Lexer.Lexer) (Class, error) {
	for isModifier(l.Pick().Val) {
		l.Increse()
	}
	if l.Pick().Val == classIdentifier {
		return parseClass(l)
	} else if l.Pick().Val == recordIdentifier {
		return parseRecord(l)
	}
	return Class{}, errors.New("no record/class found")
}

func parseRecord(l *Lexer.Lexer) (Class, error) {
	class := Class{}
	token := l.Pick()
	l.Increse()
	token = l.Pick()
	if token.Type != Lexer.Word {
		return class, errors.New("name not found")
	}
	l.Increse()
	class.Name = token.Val
	token = l.Pick()
	parms := []FieldNode{}
	if l.Pick().Type == Lexer.OpenCircle {
		l.Increse()
		for l.Pick().Type != Lexer.CloseCircle {
			parm, err := parseParmRecord(l)
			if err != nil {
				return class, err
			}
			parms = append(parms, parm)
			if l.Pick().Type == Lexer.Comma{
				l.Increse()
			}
		}
		l.Increse()
	}
	if l.Pick().Type == Lexer.Colons {
		class.ExtendType, _ = parseExtendType(l)
	}
	if l.Pick().Type == Lexer.OpenCurly{
		l.Increse()
		tparm,e:=parseFieldsList(l)
		if e!=nil{
			return class,e
		}
		parms=append(parms,tparm...)
	}else{
		l.Increse()
	}
	class.Fields = parms
	return class, nil
}
func parseExtendType(l *Lexer.Lexer) ([]INode, error) {
	extendTypes := []INode{}
	l.Increse()
	extendType, err := parseType(l)
	if err != nil {
		return extendTypes, err
	}
	l.Increse()
	extendTypes = append(extendTypes, extendType)
	return extendTypes, nil
}
func parseClass(l *Lexer.Lexer) (Class, error) {
	class := Class{}
	l.Increse()
	token := l.Pick()
	if token.Type != Lexer.Word {
		return class, errors.New("name not found")
	}
	l.Increse()
	class.Name = token.Val
	if l.Pick().Type == Lexer.Colons {
		l.Increse()
		extendType, err := parseType(l)
		if err != nil {
			return class, err
		}
		class.ExtendType = append(class.ExtendType, extendType)
		l.Increse()
	}
	token = l.GetAndGoNext()
	if token.Type != Lexer.OpenCurly {
		return class, errors.New("{ not found")
	}
	parms, err := parseFieldsList(l)
	if err != nil {
		return class, err
	}
	class.Fields = parms
	return class, nil
}

func parseFieldsList(l *Lexer.Lexer) ([]FieldNode, error) {
	parms := []FieldNode{}
	for l.Pick().Type != Lexer.CloseCurly {
		parm, err := parseField(l)
		if err != nil {
			return parms, err
		}
		parms = append(parms, parm)
		if l.Pick().Type==Lexer.CloseCurly{
			break
		}
		l.Increse()
	}
	l.Increse()
	if l.Pick().Type == Lexer.Semicolon{
		l.Increse()
	}
	return parms, nil
}
