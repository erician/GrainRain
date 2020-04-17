package lexer

import "github.com/erician/grainrain/element"

func (l *Lexer) skipComment() (ret bool) {
	ret = false
	if l.r == '/' {
		if l.peekRune(0) == '/' {
			ret = true
			l.skipSingleLineComment()
		} else if l.peekRune(0) == '*' {
			ret = true
			l.skipMultiLineComment()
		}
	}
	return ret
}

func (l *Lexer) skipSingleLineComment() {
	for l.r != element.INVALIDRUNE && element.IsLineTerminator(l.r) == false {
		l.readRune()
	}
}

func (l *Lexer) skipMultiLineComment() {
	for l.r != element.INVALIDRUNE && (!(l.r == '*' && l.peekRune(0) == '/')) {
		l.readRune()
	}
	if l.r == '*' && l.peekRune(0) == '/' {
		l.readRune()
		l.readRune()
	} else {
		// we should print error
	}
}
