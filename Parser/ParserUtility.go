package Parser

import (
	"GoFromCsToTypescript/Lexer"
	"slices"
)

//increase the lexer until one of the tokens is met
func JumpUntilOne(l *Lexer.Lexer, tokens []Lexer.TokenKind)bool {
	for {
		if slices.Contains(tokens,l.Pick().Type) {
			return true
		}
		l.Increse()
		if l.Pick().Type==Lexer.EOF{
			break
		}
	}
	return false
}
