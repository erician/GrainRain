package lexer

import (
	"math"

	"github.com/erician/grainrain/element"
)

func (l *Lexer) readStringLiteral(start rune) string {
	l.readRune()

	var literal []rune
	for l.r != element.INVALIDRUNE && l.r != '\'' {
		switch l.r {
		case '\\':
			switch l.peekRune(0) {
			// golang support escape sequence, so we directly use it.
			case 'b':
				literal = append(literal, '\b')
			case 't':
				literal = append(literal, '\t')
			case 'n':
				literal = append(literal, '\n')
			case 'v':
				literal = append(literal, '\v')
			case 'f':
				literal = append(literal, '\f')
			case 'r':
				literal = append(literal, '\r')
			case '"':
				literal = append(literal, '"')
			case '\'':
				literal = append(literal, '\'')
			case '\\':
				literal = append(literal, '\\')
			// to parse \x, \u, \0, ref:
			// https://mathiasbynens.be/notes/javascript-escapes
			case 'x':
				if element.IsHexDigit(l.peekRune(1)) && element.IsHexDigit(l.peekRune(2)) {
					literal = append(literal,
						element.HexDigit2Number(l.peekRune(1))*16+
							element.HexDigit2Number(l.peekRune(2)))
					l.readRune()
					l.readRune()
				} else {
					// we should print error info
				}
			case 'u':
				if l.peekRune(1) == '{' {
					// \u {CodePoint}
					endOff := 2
					for l.peekRune(endOff) != '}' && l.peekRune(endOff) != element.INVALIDRUNE {
						endOff++
					}
					if l.peekRune(endOff) == element.INVALIDRUNE {
						// we should print error info
					} else {
						var v rune
						v = 0
						i := 0
						endOff--
						for ; endOff != 1; endOff-- {
							v += element.HexDigit2Number(l.peekRune(endOff)) *
								(rune)(math.Pow(16, float64(i)))
							if v > element.BIGGESTCODEPOINT {
								// we should print error info
							}
							i++
						}
						literal = append(literal, v)
						for ; i > 0; i-- {
							l.readRune()
						}
						l.readRune()
						l.readRune()
					}
				} else {
					// \u Hex4Digits
					if element.IsHexDigit(l.peekRune(1)) && element.IsHexDigit(l.peekRune(2)) &&
						element.IsHexDigit(l.peekRune(3)) && element.IsHexDigit(l.peekRune(4)) {
						literal = append(literal,
							element.HexDigit2Number(l.peekRune(1))*4096+
								element.HexDigit2Number(l.peekRune(2))*256+
								element.HexDigit2Number(l.peekRune(3))*16+
								element.HexDigit2Number(l.peekRune(4)))
						l.readRune()
						l.readRune()
						l.readRune()
						l.readRune()
					} else {
						// we should print error info
					}
				}
			default:
				if element.IsOctalDigit(l.peekRune(0)) {
					// if strict mode code {
					// 		if l.peekRune(0) == '0' {
					//			literal = append(literal, l.peekRune(0)-'0')
					//      }
					//} else {
					// read LegacyOctalEscapeSequence, while is the following codes
					if element.IsOctalDigit(l.peekRune(1)) == false {
						literal = append(literal, l.peekRune(0)-'0')
					} else {
						if element.IsZeroToThree(l.peekRune(0)) {
							if element.IsOctalDigit(l.peekRune(2)) {
								literal = append(literal,
									(l.peekRune(0)-'0')*64+(l.peekRune(1)-'0')*8+l.peekRune(2)-'0')
								l.readRune()
								l.readRune()
							} else {
								literal = append(literal, (l.peekRune(0)-'0')*8+l.peekRune(1)-'0')
								l.readRune()
							}
						} else if element.IsFourToSeven(l.peekRune(0)) {
							literal = append(literal, (l.peekRune(0)-'0')*8+l.peekRune(1)-'0')
							l.readRune()
						}
					}
				} else if element.IsLineTerminator(l.peekRune(0)) {
					if l.peekRune(0) == element.CR && l.peekRune(1) == element.LF {
						// <CR><LF>
						l.readRune()
					}
				} else {
					// if not any escape character
					literal = append(literal, l.peekRune(0))
				}
			}
			l.readRune()
		case start:
			// we should print error info.
		default:
			if l.r == element.CR || l.r == element.LF {
				// we should print error info.
			} else {
				// <LS> <PS> are legal
				literal = append(literal, l.r)
			}
		}
		l.readRune()
	}
	if l.r == '\'' {
		l.readRune()
	} else if l.r == element.INVALIDRUNE {
		// we should print error info.
	}
	return string(literal)
}
