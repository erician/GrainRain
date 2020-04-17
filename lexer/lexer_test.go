package lexer

import (
	"testing"

	"github.com/erician/grainrain/element"
)

func TestNextToken(t *testing.T) {
	input := `>>>>>=+-`
	l := NewLexer([]rune(input))
	want := []element.Token{
		{element.TRIPLEGREATER, element.TRIPLEGREATER, 0, 0},
		{element.DOUBLEGREATERASSIGN, element.DOUBLEGREATERASSIGN, 0, 0},
		{element.PLUS, element.PLUS, 0, 0},
		{element.MINUS, element.MINUS, 0, 0},
		{element.EOF, element.EOF, 0, 0},
	}
	for _, tok := range want {
		got := l.NextToken()
		if tok.Type != got.Type {
			t.Errorf("get token Type: %s, expected: %s", got.Type, tok.Type)
		}
		if tok.Literal != got.Literal && tok.Type != element.EOF {
			t.Errorf("get token Literal: %s, expected: %s", got.Literal, tok.Literal)
		}
	}
}

func TestNextTokenWithWhiteSpace(t *testing.T) {
	input := `>>>>>= +-`
	l := NewLexer([]rune(input))
	want := []element.Token{
		{element.TRIPLEGREATER, element.TRIPLEGREATER, 0, 0},
		{element.DOUBLEGREATERASSIGN, element.DOUBLEGREATERASSIGN, 0, 0},
		{element.PLUS, element.PLUS, 0, 0},
		{element.MINUS, element.MINUS, 0, 0},
		{element.EOF, element.EOF, 0, 0},
	}
	for _, tok := range want {
		got := l.NextToken()
		if tok.Type != got.Type {
			t.Errorf("get token Type: %s, expected: %s", got.Type, tok.Type)
		}
		if tok.Literal != got.Literal && tok.Type != element.EOF {
			t.Errorf("get token Literal: %s, expected: %s", got.Literal, tok.Literal)
		}
	}
}

func TestNextTokenWithComment(t *testing.T) {
	input := `>>>>>= 
	//haha  
	/*haha*/ 
	+-`
	l := NewLexer([]rune(input))
	want := []element.Token{
		{element.TRIPLEGREATER, element.TRIPLEGREATER, 0, 0},
		{element.DOUBLEGREATERASSIGN, element.DOUBLEGREATERASSIGN, 0, 0},
		{element.PLUS, element.PLUS, 0, 0},
		{element.MINUS, element.MINUS, 0, 0},
		{element.EOF, element.EOF, 0, 0},
	}
	for _, tok := range want {
		got := l.NextToken()
		if tok.Type != got.Type {
			t.Errorf("get token Type: %s, expected: %s", got.Type, tok.Type)
		}
		if tok.Literal != got.Literal && tok.Type != element.EOF {
			t.Errorf("get token Literal: %s, expected: %s", got.Literal, tok.Literal)
		}
	}
}

func TestNextTokenWithStringLiteral(t *testing.T) {
	input := `>>>>>= 
	//haha  
	/*haha*/ 
	+-
	'aa'
	'\0'
	'\141'
	'\0123'
	'\x617'
	'\u00618'
	'\u{0061}9'
	"\u{0061}10"`
	l := NewLexer([]rune(input))
	want := []element.Token{
		{element.TRIPLEGREATER, element.TRIPLEGREATER, 0, 0},
		{element.DOUBLEGREATERASSIGN, element.DOUBLEGREATERASSIGN, 0, 0},
		{element.PLUS, element.PLUS, 0, 0},
		{element.MINUS, element.MINUS, 0, 0},
		{element.STRINGLITERAL, "aa", 0, 0},
		{element.STRINGLITERAL, "\u0000", 0, 0},
		{element.STRINGLITERAL, "a", 0, 0},
		{element.STRINGLITERAL, "\u000a3", 0, 0},
		{element.STRINGLITERAL, "a7", 0, 0},
		{element.STRINGLITERAL, "a8", 0, 0},
		{element.STRINGLITERAL, "a9", 0, 0},
		{element.STRINGLITERAL, "a10", 0, 0},
		{element.EOF, element.EOF, 0, 0},
	}
	for _, tok := range want {
		got := l.NextToken()
		if tok.Type != got.Type {
			t.Errorf("get token Type: %s, expected: %s", got.Type, tok.Type)
		}
		if tok.Literal != got.Literal && tok.Type != element.EOF {
			t.Errorf("get token Literal: %s, expected: %s", got.Literal, tok.Literal)
		}
	}
}
