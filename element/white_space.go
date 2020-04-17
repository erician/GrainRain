package element

import "unicode"

// these are all white spaces in JS:

// Code Point	Name					Abbreviation
// U+0009		CHARACTER TABULATION	<TAB>
// U+000B		LINE TABULATION			<VT>
// U+000C		FORM FEED (FF)			<FF>
// U+0020		SPACE					<SP>
// U+00A0		NO-BREAK SPACE			<NBSP>
// U+FEFF		ZERO WIDTH NO-BREAK SPACE	<ZWNBSP>
// Other category “Zs”	Any other Unicode “Space_Separator” code point	<USP>

// NOTE: we can use IsSpace in golang to check if it is a USP,
// the webset is: https://golang.org/pkg/unicode/
const (
	TAB 	=	'\u0009'
	VT  	= 	'\u000B'
	FF  	= 	'\u000C'
	SP  	= 	'\u0020'
	NBSP 	= 	'\u00A0'
	ZWNBSP 	= 	'\uFEFF'
	// USP, we will use IsSpace to check this
)

// IsWhiteSpace check if code_point is a white space in JS
// rune is int32
func IsWhiteSpace(r rune) bool {
	return (r == TAB || r == VT   || r == FF ||
	   		r == SP  || r == NBSP || r == ZWNBSP ||
	   		unicode.IsSpace(r) == true)
}