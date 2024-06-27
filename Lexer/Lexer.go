package Lexer

import "regexp"

type TokenKind int8

type Token struct {
	Type TokenKind
	Val  string
}
type Lexer struct {
	tokens []Token
}

const regex = `\w+|[{|}]|[<|>]|[(|)]|;`
const (
	Error TokenKind = -1
  Word TokenKind = iota
	OpenCurly
	CloseCurly
	OpenCircle
	OpenAngle
	CloseAngle
	CloseCircle
	Semicolon

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

func getType(t string) TokenKind {
	switch t {
	case "{":
		return OpenCurly
	case "}":
		return CloseCurly
	case "<":
		return OpenCircle
	case ">":
		return CloseCircle
	case "(":
		return OpenAngle
	case ")":
		return CloseAngle
	case ";":
		return Semicolon
	default:
		if (t[0] >= 'a' && t[0] <= 'z')||(t[0] >= 'A' && t[0] <= 'Z'){
			return Word
		}
		return Error
	}
}
