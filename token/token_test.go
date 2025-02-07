package token_test

import (
	"testing"

	"github.com/Matei-Stoian/GoBAS/token"
)

func TestLookupOperator(t *testing.T) {
	tests := map[string]token.Type{
		"=":  token.ASSIGN,
		"+":  token.PLUS,
		"-":  token.MINUS,
		"*":  token.ASTERISK,
		"/":  token.SLASH,
		"%":  token.MOD,
		"^":  token.POW,
		">":  token.GT,
		">=": token.GTEQUALS,
		"<":  token.LT,
		"<=": token.LTEQUALS,
		"<>": token.NOTEQUALS,
	}

	for input, expected := range tests {
		if tok, ok := token.LookupOperator(input); !ok || tok != expected {
			t.Errorf("LookupOperator(%q) = %q, expected %q", input, tok, expected)
		}
	}
}

func TestLookupIdentifier(t *testing.T) {
	tests := map[string]token.Type{
		"if":     token.IF,
		"then":   token.THEN,
		"else":   token.ELSE,
		"goto":   token.GOTO,
		"input":  token.INPUT,
		"return": token.RETURN,
		"step":   token.STEP,
		"for":    token.FOR,
		"next":   token.NEXT,
		"custom": token.IDENT, // Not a keyword, should return IDENT
	}

	for input, expected := range tests {
		if tok := token.LookupIdentifier(input); tok != expected {
			t.Errorf("LookupIdentifier(%q) = %q, expected %q", input, tok, expected)
		}
	}
}

func TestTokenString(t *testing.T) {
	tests := []struct {
		input    token.Token
		expected string
	}{
		{token.Token{Type: token.INTEGER, Literal: "42"}, "Token{Type: INTEGER, Value: \"42\"}"},
		{token.Token{Type: token.STRING, Literal: "Hello"}, "Token{Type: STRING, Value: \"Hello\"}"},
		{token.Token{Type: token.NEWLINE, Literal: ""}, "Token{Type: NEWLINE, Value: \\n}"},
		{token.Token{Type: token.EOF, Literal: ""}, "Token{Type: EOF}"},
	}

	for _, test := range tests {
		if result := test.input.String(); result != test.expected {
			t.Errorf("Token.String() = %q, expected %q", result, test.expected)
		}
	}
}
