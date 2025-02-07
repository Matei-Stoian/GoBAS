package lexer

import (
	"testing"

	"github.com/Matei-Stoian/GoBAS/token"
)

func TestNextToken(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "empty input",
			input: "",
			expected: []token.Token{
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "simple print statement",
			input: `10 PRINT "Hello World"`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "10"},
				{Type: token.PRINT, Literal: "PRINT"},
				{Type: token.STRING, Literal: "Hello World"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "variable assignment",
			input: `20 LET answer = 42`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "20"},
				{Type: token.LET, Literal: "LET"},
				{Type: token.IDENT, Literal: "answer"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.INTEGER, Literal: "42"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "mathematical operations",
			input: `30 LET result = (5 + 3.14) * 2`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "30"},
				{Type: token.LET, Literal: "LET"},
				{Type: token.IDENT, Literal: "result"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.LBRACKET, Literal: "("},
				{Type: token.INTEGER, Literal: "5"},
				{Type: token.PLUS, Literal: "+"},
				{Type: token.FLOAT, Literal: "3.14"},
				{Type: token.RBRACKET, Literal: ")"},
				{Type: token.ASTERISK, Literal: "*"},
				{Type: token.INTEGER, Literal: "2"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "conditional statement",
			input: `40 IF x <= 10 THEN 50`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "40"},
				{Type: token.IF, Literal: "IF"},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.LTEQUALS, Literal: "<="},
				{Type: token.INTEGER, Literal: "10"},
				{Type: token.THEN, Literal: "THEN"},
				{Type: token.INTEGER, Literal: "50"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "string variable",
			input: `50 LET name$ = "Alice"`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "50"},
				{Type: token.LET, Literal: "LET"},
				{Type: token.IDENT, Literal: "name$"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.STRING, Literal: "Alice"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "multiple operators",
			input: `60 IF a <> b AND c >= d OR e < f THEN 100`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "60"},
				{Type: token.IF, Literal: "IF"},
				{Type: token.IDENT, Literal: "a"},
				{Type: token.NOTEQUALS, Literal: "<>"},
				{Type: token.IDENT, Literal: "b"},
				{Type: token.AND, Literal: "AND"},
				{Type: token.IDENT, Literal: "c"},
				{Type: token.GTEQUALS, Literal: ">="},
				{Type: token.IDENT, Literal: "d"},
				{Type: token.OR, Literal: "OR"},
				{Type: token.IDENT, Literal: "e"},
				{Type: token.LT, Literal: "<"},
				{Type: token.IDENT, Literal: "f"},
				{Type: token.THEN, Literal: "THEN"},
				{Type: token.INTEGER, Literal: "100"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "comment line",
			input: `100 REM This is a comment`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "100"},
				{Type: token.REM, Literal: "This is a comment"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "multiple lines",
			input: "10 PRINT \"Hello\"\n20 GOTO 10\n30 END",
			expected: []token.Token{
				{Type: token.LINENO, Literal: "10"},
				{Type: token.PRINT, Literal: "PRINT"},
				{Type: token.STRING, Literal: "Hello"},
				{Type: token.NEWLINE, Literal: "\n"},
				{Type: token.LINENO, Literal: "20"},
				{Type: token.GOTO, Literal: "GOTO"},
				{Type: token.INTEGER, Literal: "10"},
				{Type: token.NEWLINE, Literal: "\n"},
				{Type: token.LINENO, Literal: "30"},
				{Type: token.END, Literal: "END"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "illegal character",
			input: `110 LET x = @5`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "110"},
				{Type: token.LET, Literal: "LET"},
				{Type: token.IDENT, Literal: "x"},
				{Type: token.ASSIGN, Literal: "="},
				{Type: token.ILLEGAL, Literal: "@"},
				{Type: token.INTEGER, Literal: "5"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "unicode string",
			input: `70 PRINT "こんにちは世界"`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "70"},
				{Type: token.PRINT, Literal: "PRINT"},
				{Type: token.STRING, Literal: "こんにちは世界"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "case insensitive keyword",
			input: `90 pRiNt "Case Test"`,
			expected: []token.Token{
				{Type: token.LINENO, Literal: "90"},
				{Type: token.PRINT, Literal: "pRiNt"},
				{Type: token.STRING, Literal: "Case Test"},
				{Type: token.EOF, Literal: ""},
			},
		},
		{
			name:  "line without number",
			input: `PRINT "Emergency!"`,
			expected: []token.Token{
				{Type: token.PRINT, Literal: "PRINT"},
				{Type: token.STRING, Literal: "Emergency!"},
				{Type: token.EOF, Literal: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New(tt.input)

			for i, expectedTok := range tt.expected {
				tok := l.NextToken()

				if tok.Type != expectedTok.Type {
					t.Fatalf("test[%s] token %d - type mismatch\n\texpected: %+v\n\tgot: %+v",
						tt.name, i, expectedTok.Type, tok.Type)
				}

				if tok.Literal != expectedTok.Literal {
					t.Fatalf("test[%s] token %d - literal mismatch\n\texpected: %q\n\tgot: %q",
						tt.name, i, expectedTok.Literal, tok.Literal)
				}
			}
		})
	}
}
