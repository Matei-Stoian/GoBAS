package token

import (
	"fmt"
	"strings"
)

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	// Special tokens
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	NEWLINE = "NEWLINE"
	LINENO  = "LINENO"

	// Literals and identifiers
	IDENT   = "IDENT"
	INTEGER = "INTEGER"
	FLOAT   = "FLOAT"
	STRING  = "STRING"
	BOOLEAN = "BOOLEAN"
	BUILTIN = "BUILTIN"

	// Keywords
	END    = "END"
	GOSUB  = "GOSUB"
	GOTO   = "GOTO"
	INPUT  = "INPUT"
	LET    = "LET"
	PRINT  = "PRINT"
	REM    = "REM"
	RETURN = "RETURN"
	FOR    = "FOR"
	NEXT   = "NEXT"
	STEP   = "STEP"
	TO     = "TO"
	IF     = "IF"
	THEN   = "THEN"
	ELSE   = "ELSE"
	AND    = "AND"
	OR     = "OR"
	XOR    = "XOR"
	DEF    = "DEF"
	DIM    = "DIM"
	FN     = "FN"
	READ   = "READ"
	SWAP   = "SWAP"
	DATA   = "DATA"

	// Operators and symbols
	ASSIGN    = "="
	ASTERISK  = "*"
	COMMA     = ","
	MINUS     = "-"
	MOD       = "%"
	PLUS      = "+"
	SLASH     = "/"
	POW       = "^"
	COLON     = ":"
	SEMICOLON = ";"
	LBRACKET  = "("
	RBRACKET  = ")"
	LINDEX    = "["
	RINDEX    = "]"
	GT        = ">"
	GTEQUALS  = ">="
	LT        = "<"
	LTEQUALS  = "<="
	NOTEQUALS = "<>"
)

var operatorMap = map[string]Type{
	// Single-character operators
	"=": ASSIGN,
	"*": ASTERISK,
	",": COMMA,
	"-": MINUS,
	"%": MOD,
	"+": PLUS,
	"/": SLASH,
	"^": POW,
	":": COLON,
	";": SEMICOLON,
	"(": LBRACKET,
	")": RBRACKET,
	"[": LINDEX,
	"]": RINDEX,
	">": GT,
	"<": LT,

	// Multi-character operators
	">=": GTEQUALS,
	"<=": LTEQUALS,
	"<>": NOTEQUALS,
}

var keywordMap = map[string]Type{
	"and":    AND,
	"data":   DATA,
	"def":    DEF,
	"dim":    DIM,
	"else":   ELSE,
	"end":    END,
	"fn":     FN,
	"for":    FOR,
	"gosub":  GOSUB,
	"goto":   GOTO,
	"if":     IF,
	"input":  INPUT,
	"let":    LET,
	"next":   NEXT,
	"or":     OR,
	"read":   READ,
	"print":  PRINT,
	"rem":    REM,
	"return": RETURN,
	"step":   STEP,
	"swap":   SWAP,
	"then":   THEN,
	"to":     TO,
	"xor":    XOR,
}

// LookupOperator checks if a character sequence is a valid operator
func LookupOperator(op string) (Type, bool) {
	tok, ok := operatorMap[op]
	return tok, ok
}

// LookupIdentifier checks if an identifier is a reserved keyword
func LookupIdentifier(ident string) Type {
	lower := strings.ToLower(ident)
	if tok, ok := keywordMap[lower]; ok {
		return tok
	}
	return IDENT
}

// String returns a human-readable representation of the token
func (t Token) String() string {
	switch t.Type {
	case NEWLINE:
		return `Token{Type: NEWLINE, Value: \n}`
	case EOF:
		return `Token{Type: EOF}`
	default:
		return fmt.Sprintf("Token{Type: %s, Value: %q}", t.Type, t.Literal)
	}
}

// Predefined tokens for single-character operators
var (
	Assign    = Token{Type: ASSIGN, Literal: "="}
	Plus      = Token{Type: PLUS, Literal: "+"}
	Minus     = Token{Type: MINUS, Literal: "-"}
	Asterisk  = Token{Type: ASTERISK, Literal: "*"}
	Slash     = Token{Type: SLASH, Literal: "/"}
	Mod       = Token{Type: MOD, Literal: "%"}
	Pow       = Token{Type: POW, Literal: "^"}
	Comma     = Token{Type: COMMA, Literal: ","}
	Colon     = Token{Type: COLON, Literal: ":"}
	Semicolon = Token{Type: SEMICOLON, Literal: ";"}
)
