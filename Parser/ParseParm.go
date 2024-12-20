package Parser

import (
	"GoFromCsToTypescript/Lexer"
	"errors"
)

func parseParam(l *Lexer.Lexer) (FieldNode, error) {
	token := l.Pick()
	param := FieldNode{}
	if isVisibilitySetter(token.Val) {
		l.Increse()
		token = l.Pick()
	}
	Ptype, e := parseType(l)
	if e != nil {
		return param, e
	}
	param.Type = Ptype
	l.Increse()
	token = l.GetAndGoNext()
	if token.Type == Lexer.QuestionMark {
		param.Nullable = true
		token = l.Pick()
		l.Increse()
	}
	if token.Type != Lexer.Word {
		return param, errors.New("param name not found")
	}
	param.Name = token.Val
	println("name"+param.Name)
	println("first"+ l.Pick().Val)
	jumpSetGetter(l)
	jumpUntilNextField(l)
	println("after"+l.Pick().Val)
	return param, nil
}
func parseType(l *Lexer.Lexer) (INode, error) {
	token := l.Pick()
	Ptype := isBasicType(token.Val)
	if Ptype > 0 {
		return SimpleTypeNode{Type: Ptype}, nil
	} else {
		if l.PickNext().Type == Lexer.OpenAngle {
			gType, err := parseGenericType(l)
			if err != nil {
				return nil, err
			}
			return gType, nil
		}
		for l.PickNext().Type == Lexer.Dot {
			l.Increse()
			l.Increse()
		}
		return CustomTypeNode{Type: l.Pick().Val}, nil
	}
}
func parseGenericType(l *Lexer.Lexer) (GenericTypeNode, error) {
	gType := GenericTypeNode{}
	token := l.Pick()
	if token.Type != Lexer.Word {
		return gType, errors.New("generic parent type name not found")
	}
	gType.ParentName = token.Val
	l.Increse()
	token = l.GetAndGoNext()
	if token.Type != Lexer.OpenAngle {
		return gType, errors.New("< not found")
	}
	token = l.Pick()
	typeParam, e := parseType(l)
	if e != nil {
		return gType, e
	}
	gType.ChildType = append(gType.ChildType, typeParam)
	l.Increse()
	token = l.Pick()
	if token.Type == Lexer.CloseAngle {
		return gType, nil
	} else if token.Type == Lexer.Comma {
		l.Increse()
		for token.Type != Lexer.CloseAngle {
			typeParam, e := parseType(l)
			if e != nil {
				return gType, e
			}
			gType.ChildType = append(gType.ChildType, typeParam)
			l.Increse()
			token = l.Pick()
		}
	}
	return gType, nil
}

func jumpSetGetter(l *Lexer.Lexer) {
	token := l.Pick()
	if token.Type != Lexer.OpenCurly {
		return
	}
	token = l.Pick()
	for token.Type != Lexer.CloseCurly {
		l.Increse()
		token = l.Pick()
	}
}
func jumpUntilNextField(l *Lexer.Lexer) {
	token := l.Pick()
	if l.PickNext().Type != Lexer.Assignment && token.Type != Lexer.Assignment {
		return
	}
	for token.Type != Lexer.Comma && token.Type != Lexer.Semicolon && token.Type != Lexer.CloseCircle {
		l.Increse()
		token = l.Pick()
	}
}
