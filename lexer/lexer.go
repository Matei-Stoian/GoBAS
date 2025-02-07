package lexer

import (
	"strings"
	"unicode"

	"github.com/Matei-Stoian/GoBAS/token"
)

type Lexer struct {
	input          string
	position       int
	readPosition   int
	ch             rune
	atStartOfLine  bool
	lineNumberRead bool
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	l.atStartOfLine = true
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.readPosition])
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	// Handle newlines and line structure
	if l.ch == '\n' {
		tok = token.Token{Type: token.NEWLINE, Literal: "\n"}
		l.readChar()
		l.atStartOfLine = true
		l.lineNumberRead = false
		return tok
	}

	// Handle line numbers at start of line
	if l.atStartOfLine && !l.lineNumberRead {
		if unicode.IsDigit(l.ch) {
			return l.readLineNumberToken()
		}
		l.atStartOfLine = false
	}

	switch l.ch {
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		switch {
		case unicode.IsDigit(l.ch) || l.ch == '.':
			return l.readNumber()
		case unicode.IsLetter(l.ch) || l.ch == '_' || l.ch == '$':
			return l.readIdentifier()
		default:
			tok = l.readOperatorOrSymbol()
		}
	}

	if tok.Type == "" {
		l.readChar()
	}
	return tok
}

func (l *Lexer) readLineNumberToken() token.Token {
	lit := l.readLineNumber()
	l.lineNumberRead = true
	l.atStartOfLine = false
	return token.Token{Type: token.LINENO, Literal: lit}
}

func (l *Lexer) readLineNumber() string {
	start := l.position
	for unicode.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) readIdentifier() token.Token {
	start := l.position
	for isIdentifierChar(l.ch) {
		l.readChar()
	}
	lit := l.input[start:l.position]
	tokType := token.LookupIdentifier(lit)

	// Special handling for REM comments
	if tokType == token.REM {
		l.skipWhitespace()
		return token.Token{
			Type:    token.REM,
			Literal: l.readRemComment(),
		}
	}

	return token.Token{Type: tokType, Literal: lit}
}

func (l *Lexer) readRemComment() string {
	start := l.position
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	return strings.TrimSpace(l.input[start:l.position])
}

func (l *Lexer) readNumber() token.Token {
	start := l.position
	var tokType token.Type
	tokType = token.INTEGER
	decimalFound := false

	for {
		if unicode.IsDigit(l.ch) {
			l.readChar()
		} else if l.ch == '.' && !decimalFound {
			decimalFound = true
			tokType = token.FLOAT
			l.readChar()
		} else {
			break
		}
	}

	return token.Token{
		Type:    tokType,
		Literal: l.input[start:l.position],
	}
}

func (l *Lexer) readString() string {
	l.readChar() // Skip opening quote
	start := l.position

	for l.ch != '"' && l.ch != 0 {
		l.readChar()
	}

	lit := l.input[start:l.position]
	if l.ch == '"' {
		l.readChar() // Skip closing quote
	}
	return lit
}

func (l *Lexer) readOperatorOrSymbol() token.Token {
	// Handle multi-character operators first
	if l.ch == '<' || l.ch == '>' {
		op := string(l.ch)
		next := l.peekChar()

		if next == '=' || (l.ch == '<' && next == '>') {
			l.readChar()
			op += string(l.ch)
			if tokType, ok := token.LookupOperator(op); ok {
				l.readChar()
				return token.Token{Type: tokType, Literal: op}
			}
		}
	}

	// Handle single-character operators
	if tokType, ok := token.LookupOperator(string(l.ch)); ok {
		t := token.Token{Type: tokType, Literal: string(l.ch)}
		l.readChar()
		return t
	}

	// Handle illegal characters
	illegal := string(l.ch)
	l.readChar()
	return token.Token{Type: token.ILLEGAL, Literal: illegal}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

func isIdentifierChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '$'
}
