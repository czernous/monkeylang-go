package monkeylang_test

import (
	"testing"

	. "github.com/czernous/monkeylang-go/internal/ast"
	"github.com/czernous/monkeylang-go/internal/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "myVar",
					},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENT,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

    testVal := "let myVar = anotherVar;"

 
	if program.String() != testVal {
		t.Errorf("program.String() wrong, want=%q, got=%q", testVal, program.String())
	}

}
