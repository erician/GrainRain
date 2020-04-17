package lexer

import (
	"github.com/erician/grainrain/element"
)

// Lexer just like the name, the input should be []rune
type Lexer struct {
	input        []rune
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	r            rune // current code point examination
}

// NewLexer create a new Lexer
func NewLexer(input []rune) *Lexer {
	l := &Lexer{input: input}
	l.readRune()
	return l
}

func (l *Lexer) readRune() {
	if l.readPosition >= len(l.input) {
		// assign that is never met in source codes.
		l.r = element.INVALIDRUNE
	} else {
		l.r = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// the start pos is readPosition, so if not overflow,
// peekRune will return l.input[l.readPosition+offset]
func (l *Lexer) peekRune(offset int) rune {
	if l.readPosition+offset >= len(l.input) {
		return element.INVALIDRUNE
	}
	return l.input[l.readPosition+offset]
}

// if we really skip a/some white space, return true, or false
func (l *Lexer) skipWhitespace() (ret bool) {
	ret = element.IsWhiteSpace(l.r)
	for element.IsWhiteSpace(l.r) {
		l.readRune()
	}
	return ret
}

func newToken(tokenType element.TokenType, literal string) element.Token {
	return element.Token{Type: tokenType, Literal: literal}
}

// NextToken return next token
func (l *Lexer) NextToken() element.Token {
	var tok element.Token
	for {
		if l.skipWhitespace() == false && l.skipComment() == false {
			break
		}
	}

	switch l.r {
	// IdentifierName
	// Punctuator
	case '{':
		tok = newToken(element.LBRACE, element.LBRACE)
	case '(':
		tok = newToken(element.LPAREN, element.LPAREN)
	case ')':
		tok = newToken(element.RPAREN, element.RPAREN)
	case '[':
		tok = newToken(element.LBRACKET, element.LBRACKET)
	case ']':
		tok = newToken(element.RBRACKET, element.RBRACKET)
	case '.':
		if l.peekRune(0) == '.' && l.peekRune(1) == '.' {
			l.readRune()
			l.readRune()
			tok = newToken(element.ELLIPSIS, element.ELLIPSIS)
		} else {
			tok = newToken(element.DOT, element.DOT)
		}
	case ';':
		tok = newToken(element.SEMICOLON, element.SEMICOLON)
	case ',':
		tok = newToken(element.COMMA, element.COMMA)
	case '<':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.LESSOREQUAL, element.LESSOREQUAL)
		} else if l.peekRune(0) == '<' {
			if l.peekRune(1) == '=' {
				l.readRune()
				l.readRune()
				tok = newToken(element.DOUBLELESSASSIGN, element.DOUBLELESSASSIGN)
			} else {
				l.readRune()
				tok = newToken(element.DOUBLELESS, element.DOUBLELESS)
			}
		} else {
			tok = newToken(element.LESS, element.LESS)
		}
	case '>':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.GREATEROREQUAL, element.GREATEROREQUAL)
		} else if l.peekRune(0) == '>' {
			if l.peekRune(1) == '>' {
				if (l.peekRune(2)) == '=' {
					l.readRune()
					l.readRune()
					l.readRune()
					tok = newToken(element.TRIPLEGREATERASSIGN, element.TRIPLEGREATERASSIGN)
				} else {
					l.readRune()
					l.readRune()
					tok = newToken(element.TRIPLEGREATER, element.TRIPLEGREATER)
				}
			} else if l.peekRune(1) == '=' {
				l.readRune()
				l.readRune()
				tok = newToken(element.DOUBLEGREATERASSIGN, element.DOUBLEGREATERASSIGN)
			} else {
				l.readRune()
				tok = newToken(element.DOUBLEGREATER, element.DOUBLEGREATER)
			}
		} else {
			tok = newToken(element.GREATER, element.GREATER)
		}
	case '=':
		if l.peekRune(0) == '>' {
			l.readRune()
			tok = newToken(element.ASSIGNGREATER, element.ASSIGNGREATER)
		} else if l.peekRune(0) == '=' {
			if l.peekRune(1) == '=' {
				l.readRune()
				l.readRune()
				tok = newToken(element.STRICTEQUAL, element.STRICTEQUAL)
			} else {
				l.readRune()
				tok = newToken(element.EQUAL, element.EQUAL)
			}
		} else {
			tok = newToken(element.ASSIGN, element.ASSIGN)
		}
	case '!':
		if l.peekRune(0) == '=' {
			if l.peekRune(1) == '=' {
				l.readRune()
				l.readRune()
				tok = newToken(element.STRICTNOTEQUAL, element.STRICTNOTEQUAL)
			} else {
				l.readRune()
				tok = newToken(element.NOTEQUAL, element.NOTEQUAL)
			}
		} else {
			tok = newToken(element.EXCLAMATION, element.EXCLAMATION)
		}
	case '+':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.PLUSASSIGN, element.PLUSASSIGN)
		} else if l.peekRune(0) == '+' {
			l.readRune()
			tok = newToken(element.DOUBLEPLUS, element.DOUBLEPLUS)
		} else {
			tok = newToken(element.PLUS, element.PLUS)
		}
	case '-':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.MINUSASSIGN, element.MINUSASSIGN)
		} else if l.peekRune(0) == '-' {
			l.readRune()
			tok = newToken(element.DOUBLEMINUS, element.DOUBLEMINUS)
		} else {
			tok = newToken(element.MINUS, element.MINUS)
		}
	case '*':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.STARASSIGN, element.STARASSIGN)
		} else if l.peekRune(0) == '*' {
			if l.peekRune(1) == '=' {
				l.readRune()
				l.readRune()
				tok = newToken(element.DOUBLESTARASSIGN, element.DOUBLESTARASSIGN)
			} else {
				l.readRune()
				tok = newToken(element.DOUBLESTAR, element.DOUBLESTAR)
			}
		} else {
			tok = newToken(element.STAR, element.STAR)
		}
	case '%':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.PERCENTASSIGN, element.PERCENTASSIGN)
		} else {
			tok = newToken(element.PERCENT, element.PERCENT)
		}
	case '&':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.AMPASSIGN, element.AMPASSIGN)
		} else if l.peekRune(0) == '&' {
			l.readRune()
			tok = newToken(element.AND, element.AND)
		} else {
			tok = newToken(element.AMP, element.AMP)
		}
	case '|':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.BARASSIGN, element.BARASSIGN)
		} else if l.peekRune(0) == '|' {
			l.readRune()
			tok = newToken(element.OR, element.OR)
		} else {
			tok = newToken(element.BAR, element.BAR)
		}
	case '^':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.CARETASSIGN, element.CARETASSIGN)
		} else {
			tok = newToken(element.CARET, element.CARET)
		}
	case '~':
		tok = newToken(element.TILDE, element.TILDE)
	case '?':
		tok = newToken(element.QUESTION, element.QUESTION)
	case ':':
		tok = newToken(element.COLON, element.COLON)
	case '/':
		if l.peekRune(0) == '=' {
			l.readRune()
			tok = newToken(element.SLASHASSIGN, element.SLASHASSIGN)
		} else {
			tok = newToken(element.SLASH, element.SLASH)
		}
	case '}':
		tok = newToken(element.RBRACE, element.RBRACE)
	// NumericLiteral
	// StringLiteral
	case '\'':
		tok.Type = element.STRINGLITERAL
		tok.Literal = l.readStringLiteral('\'')
	case '"':
		tok.Type = element.STRINGLITERAL
		tok.Literal = l.readStringLiteral('"')
	// Template
	case element.INVALIDRUNE:
		tok.Literal = ""
		tok.Type = element.EOF
	}
	l.readRune()
	return tok
}
