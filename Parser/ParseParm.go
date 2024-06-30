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
	if token.Type==Lexer.QuestionMark {
		param.Nullable = true
		token = l.Pick()
		l.Increse()
	}
	if token.Type != Lexer.Word {
		return param, errors.New("param name not found")
	}
	param.Name = token.Val
	token=l.Pick()
	if token.Type!=Lexer.OpenCurly{
		return param, nil
	}
	l.Increse()
	token=l.Pick()
	for token.Type!=Lexer.CloseCurly{
		l.Increse()
		token=l.Pick()
	}
	if l.PickNext().Type!=Lexer.Assignment{
		return param, nil
	}
	l.Increse()
	token=l.Pick()
	if token.Type==Lexer.Semicolon{
		return param, nil
	}
	for token.Type!=Lexer.Semicolon{
		l.Increse()
		token=l.Pick()
	}
	return param, nil
}
func parseType(l *Lexer.Lexer) (ITypeNode, error) {
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
		return CustomTypeNode{Type: token.Val}, nil
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
