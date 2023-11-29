package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	LET       = "LET"
	FUNCTION  = "FUNCTION"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	ELSE      = "ELSE"
	RETURN    = "RETURN"
	INT       = "INT"
	PLUS      = "+"
	MINUS     = "-"
	ASSIGN    = "="
	BANG      = "!"
	COMMA     = ","
	SEMICOLON = ";"
	ASTERISK  = "*"
	SLASH     = "/"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NOT_EQ    = "!="
	IDENT     = "IDENT"
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
