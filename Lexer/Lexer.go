package Lexer

import (
	"errors"
	"regexp"
)

type TokenKind int8

type Token struct {
	Type TokenKind
	Val  string
}
type Lexer struct {
	tokens []Token
	cur    int
}

const regex = `\w+|[{|}]|[<|>]|[(|)]|;|\?|,`
const (
	Error TokenKind = -2
	Unknow TokenKind = -1
  Word TokenKind = iota
	OpenCurly
	CloseCurly
	OpenCircle
	OpenAngle
	CloseAngle
	CloseCircle
	Semicolon
	QuestionMark
	Comma

)
func New(input string) (*Lexer,error) {
	reg,e:=regexp.Compile(regex)
	if e!=nil{
		return nil,e
	}
	var tokens []Token
	for _,v:=range reg.FindAllString(input,-1){
		tokens=append(tokens,Token{Val:v,Type:getType(v)})
	}
	lexer:=Lexer{tokens:tokens}
	return &lexer,nil
}

func (l *Lexer) GetAndGoNext ()Token {
	if l.cur >= len(l.tokens) {
		return Token{Type: Error, Val: ""}
	}
	t := l.tokens[l.cur]
	l.cur++
	return t
}

func (l *Lexer) Pick ()Token {
	t := l.tokens[l.cur]
	return t
}
func (l *Lexer) Increse () error {
	if l.cur+1 >= len(l.tokens) {
		return errors.New("out of range")
	}
	l.cur++
	return nil
}
func (l *Lexer) PickNext ()Token {
	if l.cur+1 >= len(l.tokens) {
		return Token{Type: Error, Val: ""}
	}
	t := l.tokens[l.cur+1]
	return t
}

func (l *Lexer) PickPre ()Token {
	if l.cur-1 < 0 {
		return Token{Type: Error, Val: ""}
	}
	t := l.tokens[l.cur-1]
	return t
}

func getType(t string) TokenKind {
	switch t {
	case "{":
		return OpenCurly
	case "}":
		return CloseCurly
	case "<":
		return OpenAngle
	case ">":
		return CloseAngle
	case "(":
		return OpenCircle
	case ")":
		return CloseCircle
	case ";":
		return Semicolon
	case "?":
		return QuestionMark
	case ",":
		return Comma
	default:
		if (t[0] >= 'a' && t[0] <= 'z')||(t[0] >= 'A' && t[0] <= 'Z'){
			return Word
		}
		return Unknow
	}
}
