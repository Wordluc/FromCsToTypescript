package Parser

import (
	"GoFromCsToTypescript/Lexer"
	"slices"
)

//increase the lexer until one of the tokens is met
func JumpUntilOne(l *Lexer.Lexer, tokens []Lexer.TokenKind) {
	for {
		if slices.Contains(tokens,l.Pick().Type) {
			l.Increse()
			break
		}
		l.Increse()
	}
}
