package element

// line terminator are included in the set of white space code points.
// these are all line terminators in JS:

// Code Point	Unicode Name			Abbreviation
// U+000A		LINE FEED (LF)			<LF>
// U+000D		CARRIAGE RETURN (CR)	<CR>
// U+2028		LINE SEPARATOR			<LS>
// U+2029		PARAGRAPH SEPARATOR		<PS>

// NOTE: The sequence <CR><LF> is commonly used as a line terminator, not two.
// there is more details, visit:
// https://www.ecma-international.org/ecma-262/10.0/index.html#
// sec-ecmascript-language-lexical-grammar
const (
	LF 	=	'\u000A'
	CR	= 	'\u000D'
	LS 	= 	'\u2028'
	PS 	= 	'\u2029'
)

// IsLineTerminator check if r is a line terminator
func IsLineTerminator(r rune) bool {
	return (r == LF || r == CR || r == LS || r == PS)
}