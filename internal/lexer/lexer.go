package lexer

import (
	"unicode"

	"github.com/czernous/monkeylang-go/internal/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) getTokensMap() map[rune]token.Token {
	m := make(map[rune]token.Token, 14)

	m['='] = newToken(token.ASSIGN, l.ch)
	m[';'] = newToken(token.SEMICOLON, l.ch)
	m['('] = newToken(token.LPAREN, l.ch)
	m[')'] = newToken(token.RPAREN, l.ch)
	m['{'] = newToken(token.LBRACE, l.ch)
	m['}'] = newToken(token.RBRACE, l.ch)
	m['>'] = newToken(token.GT, l.ch)
	m['<'] = newToken(token.LT, l.ch)
	m[','] = newToken(token.COMMA, l.ch)
	m['+'] = newToken(token.PLUS, l.ch)
	m['-'] = newToken(token.MINUS, l.ch)
	m['/'] = newToken(token.SLASH, l.ch)
	m['!'] = newToken(token.BANG, l.ch)
	m['*'] = newToken(token.ASTERISK, l.ch)

	return m
}

func (l *Lexer) readChar() {
	inputRunes := []rune(l.input)
	if l.readPosition >= len(inputRunes) {
		l.ch = 0
	} else {

		l.ch = inputRunes[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	if isLetter(l.ch) {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)

		return tok
	}

	if isDigit(l.ch) {
		tok.Type = token.INT
		tok.Literal = l.readNumber()

		return tok
	}

	if l.ch == 0 {
		tok.Literal = ""
		tok.Type = token.EOF
	} else if _, ok := l.getTokensMap()[l.ch]; ok {
		tok = l.getTokensMap()[l.ch]
	} else {
		tok = newToken(token.ILLEGAL, l.ch)
	}

	if l.ch == '=' {
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()

			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	}

	if l.ch == '!' {
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()

			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	}

	l.readChar()

	return tok

}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return rune(l.input[l.readPosition])
	}
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func newToken(tokenType token.TokenType, ch rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
