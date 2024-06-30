package Lexer

import "testing"

func confrontTokens(t *testing.T, l *Lexer, expTokens []Token) {
	var token Token
	for i := 0; i < len(expTokens); i++ {
		token = l.PickNext()
		if i < len(expTokens)-1 {
			if token.Val != expTokens[i+1].Val {
				t.Fatal("next-val-got:", token.Val, ",expected:", expTokens[i+1].Val)
				return
			}
		}
		token = l.PickPre()
		if i > 0 {
			if token.Val != expTokens[i-1].Val {
				t.Fatal("pre-val-got:", token.Val, ",expected:", expTokens[i-1].Val)
				return
			}
		}
		token = l.GetAndGoNext()
		if token.Type != token.Type {
			t.Fatal("type-got:", token.Type, ",expected:", token.Type)
			return
		}
		if token.Val != token.Val {
			t.Fatal("val-got:", l.tokens[i].Val, ",expected:", expTokens[i].Val)
			return
		}
	}
}

func Test_TokenizingSimpleClass(t *testing.T) {
	l, err := New(`public class prova{
		//commento fcdfr
		/*ciao 
		questo
		è un commento
		*/
		public string test;}`)
	if err != nil {
		t.Fatal(err)
	}
	expTokens := []Token{
		{Type: Word, Val: "public"},
		{Type: Word, Val: "class"},
		{Type: Word, Val: "prova"},
		{Type: OpenCurly, Val: "{"},
		{Type: Word, Val: "public"},
		{Type: Word, Val: "string"},
		{Type: Word, Val: "test"},
		{Type: Semicolon, Val: ";"},
		{Type: CloseCurly, Val: "}"},
	}
	confrontTokens(t, l, expTokens)
}

func Test_TokenizingClassWithGeneric(t *testing.T) {
	l, err := New("public class prova{public pippo<cia,pippo> test;}")
	if err != nil {
		t.Fatal(err)
	}
	expTokens := []Token{
		{Type: Word, Val: "public"},
		{Type: Word, Val: "class"},
		{Type: Word, Val: "prova"},
		{Type: OpenCurly, Val: "{"},
		{Type: Word, Val: "public"},
		{Type: Word, Val: "pippo"},
		{Type: OpenAngle, Val: "<"},
		{Type: Word, Val: "cia"},
		{Type: Comma, Val: ","},
		{Type: Word, Val: "pippo"},
		{Type: CloseAngle, Val: ">"},
		{Type: Word, Val: "test"},
		{Type: Semicolon, Val: ";"},
		{Type: CloseCurly, Val: "}"},
	}
	confrontTokens(t, l, expTokens)
}

func Test_TokenizingSimpleClassWithNullableProperty(t *testing.T) {
	l, err := New("public class prova{public string? test;}")
	if err != nil {
		t.Fatal(err)
	}
	expTokens := []Token{
		{Type: Word, Val: "public"},
		{Type: Word, Val: "class"},
		{Type: Word, Val: "prova"},
		{Type: OpenCurly, Val: "{"},
		{Type: Word, Val: "public"},
		{Type: Word, Val: "string"},
		{Type: QuestionMark, Val: "?"},
		{Type: Word, Val: "test"},
		{Type: Semicolon, Val: ";"},
		{Type: CloseCurly, Val: "}"},
	}
	confrontTokens(t, l, expTokens)
}

func Test_TokenizingSimpleRecord(t *testing.T) {
	l, err := New("public record prova(string test,string test1);")
	if err != nil {
		t.Fatal(err)
	}
	expTokens := []Token{
		{Type: Word, Val: "public"},
		{Type: Word, Val: "record"},
		{Type: Word, Val: "prova"},
		{Type: OpenCircle, Val: "("},
		{Type: Word, Val: "string"},
		{Type: Word, Val: "test"},
		{Type: Comma, Val: ","},
		{Type: Word, Val: "string"},
		{Type: Word, Val: "test1"},
		{Type: CloseCircle, Val: ")"},
		{Type: Semicolon, Val: ";"},
	}
	confrontTokens(t, l, expTokens)
}

