package element

// TokenType defined in the following
type TokenType string

// Token are lexer return values
type Token struct {
	Type    TokenType
	Literal string // store with rune
	Line    uint32
	Column  uint32
}

// these are all tokens

// CommonToken:
// IdentifierName
// Punctuator
// NumericLiteral
// StringLiteral
// Template

// The DivPunctuator, RegularExpressionLiteral,
// RightBracePunctuator, and TemplateSubstitutionTail
// productions derive additional tokens that are not
// included in the CommonToken production.

// Identifier Names
const (
	IDENT = "IDENT"
)

// Reserved Words
// ReservedWord::
// 		Keyword
//		FutureReservedWord
//		NullLiteral
//		BooleanLiteral

// Keywords
const (
	AWAIT      = "await"
	BREAK      = "break"
	CASE       = "case"
	CATCH      = "catch"
	CLASS      = "class"
	CONST      = "const"
	CONTINUE   = "continue"
	DEBUGGER   = "debugger"
	DEFAULT    = "default"
	DELETE     = "delete"
	DO         = "do"
	ELSE       = "else"
	EXPORT     = "export"
	EXTENDS    = "extends"
	FINALLY    = "finally"
	FOR        = "for"
	FUNCTIONE  = "function"
	IF         = "if"
	IMPORT     = "import"
	IN         = "in"
	INSTANCEOF = "instanceof"
	NEW        = "new"
	RETURN     = "return"
	SUPER      = "super"
	SWITCH     = "switch"
	THIS       = "this"
	THROW      = "throw"
	TRY        = "try"
	TYPEOF     = "typeof"
	VAR        = "var"
	VOID       = "void"
	WHILE      = "while"
	WITH       = "with"
	YIELD      = "yield"
)

// Future Reserved Words
const (
	ENUM = "enum"
	// if use strict mode code, the followings are also reserved
	IMPLEMENTS = "implements"
	PACKAGE    = "package"
	PROTECTED  = "protected"
	INTERFACE  = "interface"
	PRIVATE    = "private"
	PUBLIC     = "public"
)

// Null Literals
const (
	NULL = "null"
)

// Boolean Literals
const (
	TRUE  = "true"
	FALSE = "false"
)

// end of ReservedWord

// Punctuators
// Punctuator::one of
//		{ ( ) [ ] . ... ; , < > <= >= == != === !==
//		+ - * % ** ++ -- << >> >>> & | ^ ! ~ && || ? :
//		= += -= *= %= **= <<= >>= >>>= &= |= ^= =>
// DivPunctuator::
//		/
//		/=
// RightBracePunctuator::
//		}
const (
	// all name for these punctuators:
	// https://wenku.baidu.com/view/db311465657d27284b73f242336c1eb91b373356.html
	// punctuator
	LBRACE              = "{"
	LPAREN              = "("
	RPAREN              = ")"
	LBRACKET            = "["
	RBRACKET            = "]"
	DOT                 = "."
	ELLIPSIS            = "..."
	SEMICOLON           = ";"
	COMMA               = ","
	LESS                = "<"
	GREATER             = ">"
	LESSOREQUAL         = "<="
	GREATEROREQUAL      = ">="
	EQUAL               = "=="
	NOTEQUAL            = "!="
	STRICTEQUAL         = "==="
	STRICTNOTEQUAL      = "!=="
	PLUS                = "+"
	MINUS               = "-"
	STAR                = "*"
	PERCENT             = "%"
	DOUBLESTAR          = "**" // pow(x, y)
	DOUBLEPLUS          = "++"
	DOUBLEMINUS         = "--"
	DOUBLELESS          = "<<"
	DOUBLEGREATER       = ">>"
	TRIPLEGREATER       = ">>>"
	AMP                 = "&"
	BAR                 = "|"
	CARET               = "^"
	EXCLAMATION         = "!"
	TILDE               = "~"
	AND                 = "&&"
	OR                  = "||"
	QUESTION            = "?"
	COLON               = ":"
	ASSIGN              = "="
	PLUSASSIGN          = "+="
	MINUSASSIGN         = "-="
	STARASSIGN          = "*="
	PERCENTASSIGN       = "%="
	DOUBLESTARASSIGN    = "**="
	DOUBLELESSASSIGN    = "<<="
	DOUBLEGREATERASSIGN = ">>="
	TRIPLEGREATERASSIGN = ">>>="
	AMPASSIGN           = "&="
	BARASSIGN           = "|="
	CARETASSIGN         = "^="
	ASSIGNGREATER       = "=>"
	// DivPunctuator
	SLASH       = "/"
	SLASHASSIGN = "/="
	// RightBracePunctuator
	RBRACE = "}"
)

// define EOP
const (
	EOF = "EOF"
	// Unicode code point values from U+0000 to U+10FFFF
	INVALIDRUNE = 0x7fffffff
)

// Numeric Literals
// NumericLiteral::
//		DecimalLiteral
//		BinaryIntegerLiteral
//		OctalIntegerLiteral
//		HexIntegerLiteral
const (
	NUMERICLITERAL    = "NUMERICLITERAL"
	SMALLESTCODEPOINT = 0x0
	BIGGESTCODEPOINT  = 0x10FFFF
)

// IsOctalDigit check r is a octal digit
func IsOctalDigit(r rune) bool {
	return (IsZeroToThree(r) || IsFourToSeven(r))
}

// IsZeroToThree just like the name
func IsZeroToThree(r rune) bool {
	return (r == '0' || r == '1' || r == '2' || r == '3')
}

// IsFourToSeven just like the name
func IsFourToSeven(r rune) bool {
	return (r == '4' || r == '5' || r == '6' || r == '7')
}

// IsHexDigit just like the name
func IsHexDigit(r rune) bool {
	return (r == '0' || r == '1' || r == '2' || r == '3' ||
		r == '4' || r == '5' || r == '6' || r == '7' ||
		r == '8' || r == '9' || r == 'a' || r == 'b' ||
		r == 'c' || r == 'd' || r == 'e' || r == 'f' ||
		r == 'A' || r == 'B' || r == 'C' || r == 'D' ||
		r == 'E' || r == 'F')
}

// HexDigit2Number transfer hex digit code point to number
func HexDigit2Number(r rune) rune {
	if r >= '0' && r <= '9' {
		return r - '0'
	} else if r >= 'a' && r <= 'f' {
		return r - 'a'
	} else if r >= 'A' && r <= 'F' {
		return r - 'A'
	}
	return INVALIDRUNE
}

// String Literals
// StringLiteral::
//		"DoubleStringCharactersopt"
//		'SingleStringCharactersopt'
const (
	STRINGLITERAL = "STRINGLITERAL"
)

// Regular Expression Literals
// RegularExpressionLiteral::
//		/RegularExpressionBody/RegularExpressionFlags
const (
	REGEXP = "REGEXP"
)

// Template Literal Lexical Components
// Template::
//		NoSubstitutionTemplate
//		TemplateHead
const (
	TEMPLATE = "TEMPLATE"
)
