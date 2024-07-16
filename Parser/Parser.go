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

func isVisibilitySetter(t string) bool {
	switch t {
	case "public", "private", "protected":
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
		token := l.Pick()
    class, err = parse(token,l)
		if err != nil {
			token = l.PickNext()
			class, err = parse(token,l)
			if err != nil {
				return []Class{}, err
			}
		}
		classes = append(classes, class)
		token = l.Pick()
		if token.Type == Lexer.EOF {
			return classes, nil
		}
	}
}

func parse(token Lexer.Token,l *Lexer.Lexer) (Class, error) {
	  if token.Val == abstractIdentifier{
		  l.Increse()
			return parse(l.Pick(),l)
	  }else if token.Val == classIdentifier {
			return parseClass(l)
		} else if token.Val == recordIdentifier {
			return parseRecord(l)
		}
		return Class{}, errors.New("class not found")
}

func parseRecord(l *Lexer.Lexer) (Class, error) {
	class := Class{}
	token := l.Pick()
	if isVisibilitySetter(token.Val) {
		l.Increse()
	}
	l.Increse()
	token = l.Pick()
	if token.Type != Lexer.Word {
		return class, errors.New("name not found")
	}
	l.Increse()
	class.Name = token.Val
	token = l.Pick()
	if token.Type != Lexer.OpenCircle {
		l.Increse()
		//l.Increse()
	}
	l.Increse()
	if l.Pick().Type == Lexer.Colons {
		class.ExtendType, _ = parseExtendType(l)
		l.Increse()
	}
	if l.Pick().Type == Lexer.OpenCurly {
		l.Increse()
	}
	token=l.Pick()
	parms := []FieldNode{}
	for token.Type != Lexer.CloseCircle && token.Type != Lexer.CloseCurly && token.Type != Lexer.Semicolon {

		parm, err := parseParam(l)
		if err != nil {
			return class, err
		}
		l.Increse()
		parms = append(parms, parm)
		if l.Pick().Type == Lexer.Colons {
			class.ExtendType, err = parseExtendType(l)
		}
		token = l.Pick()
		if l.Pick().Type == Lexer.OpenCurly {
			l.Increse()
		}

	}
	class.Fields = parms
	token = l.GetAndGoNext()
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
	token := l.Pick()
	if isVisibilitySetter(token.Val) {
		l.Increse()
	}
	l.Increse()
	token = l.Pick()
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
