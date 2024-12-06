package golox

import (
	"testing"
)

func TestScanToken(t *testing.T) {
	tests := []struct {
		source string
		want   []Token
	}{
		{"( ) { } , . - + ; * !  = < > /  \r \t 123 1234.55 >= <= == != \"string\" var while for number",
			[]Token{{Type: LEFT_PAREN, Lexeme: "(", Literal: nil, Line: 1},
				{Type: RIGHT_PAREN, Lexeme: ")", Literal: nil, Line: 1},
				{Type: LEFT_BRACE, Lexeme: "{", Literal: nil, Line: 1},
				{Type: RIGHT_BRACE, Lexeme: "}", Literal: nil, Line: 1},
				{Type: COMMA, Lexeme: ",", Literal: nil, Line: 1},
				{Type: DOT, Lexeme: ".", Literal: nil, Line: 1},
				{Type: MINUS, Lexeme: "-", Literal: nil, Line: 1},
				{Type: PLUS, Lexeme: "+", Literal: nil, Line: 1},
				{Type: SEMICOLON, Lexeme: ";", Literal: nil, Line: 1},
				{Type: STAR, Lexeme: "*", Literal: nil, Line: 1},
				{Type: BANG, Lexeme: "!", Literal: nil, Line: 1},
				{Type: EQUAL, Lexeme: "=", Literal: nil, Line: 1},
				{Type: LESS, Lexeme: "<", Literal: nil, Line: 1},
				{Type: GREATER, Lexeme: ">", Literal: nil, Line: 1},
				{Type: SLASH, Lexeme: "/", Literal: nil, Line: 1},
				{Type: NUMBER, Lexeme: "123", Literal: float64(123), Line: 1},
				{Type: NUMBER, Lexeme: "1234.55", Literal: 1234.55, Line: 1},
				{Type: GREATER_EQUAL, Lexeme: ">=", Literal: nil, Line: 1},
				{Type: LESS_EQUAL, Lexeme: "<=", Literal: nil, Line: 1},
				{Type: EQUAL_EQUAL, Lexeme: "==", Literal: nil, Line: 1},
				{Type: BANG_EQUAL, Lexeme: "!=", Literal: nil, Line: 1},
				{Type: STRING, Lexeme: "\"string\"", Literal: "string", Line: 1},
				{Type: VAR, Lexeme: "var", Literal: nil, Line: 1},
				{Type: WHILE, Lexeme: "while", Literal: nil, Line: 1},
				{Type: FOR, Lexeme: "for", Literal: nil, Line: 1},
				{Type: IDENTIFIER, Lexeme: "number", Literal: nil, Line: 1},
				{Type: EOF, Lexeme: "", Literal: nil, Line: 1}},
		},
		{
			"/* var privet //////////////// \n \n \n \n \n \n \n /* */ */ // var skip while for != == \n var",
			[]Token{
				{Type: VAR, Lexeme: "var", Literal: nil, Line: 9},
				{Type: EOF, Lexeme: "", Literal: nil, Line: 9},
			},
		},
		{
			"123.var/;for",
			[]Token{
				{Type: NUMBER, Lexeme: "123", Literal: float64(123), Line: 1},
				{Type: DOT, Lexeme: ".", Literal: nil, Line: 1},
				{Type: VAR, Lexeme: "var", Literal: nil, Line: 1},
				{Type: SLASH, Lexeme: "/", Literal: nil, Line: 1},
				{Type: SEMICOLON, Lexeme: ";", Literal: nil, Line: 1},
				{Type: FOR, Lexeme: "for", Literal: nil, Line: 1},
				{Type: EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
		{
			"var for @ #",
			[]Token{
				{Type: VAR, Lexeme: "var", Literal: nil, Line: 1},
				{Type: FOR, Lexeme: "for", Literal: nil, Line: 1},
				{Type: EOF, Lexeme: "", Literal: nil, Line: 1},
			},
		},
	}

	for numberTest, test := range tests {
		actualTokens := NewScanner(test.source).ScanTokens()

		if len(actualTokens) != len(test.want) {
			t.Errorf("expected amount tokens = %d, but actual amount = %d", len(actualTokens), len(test.want))
		}

		for i := 0; i < len(actualTokens); i++ {
			if actualTokens[i].Type != test.want[i].Type {
				t.Errorf("expected token type [view = %s, code = %d], but actual token type [view = %s, code = %d] at position test [index test = %d, index line test = %d]",
					actualTokens[i].Type.String(), actualTokens[i].Type, test.want[i].Type.String(), test.want[i].Type, numberTest, i)
			}

			if actualTokens[i].Lexeme != test.want[i].Lexeme {
				t.Errorf("expected lexeme = %s, but actual lexeme %s at position test [index test = %d, index line test = %d]", actualTokens[i].Lexeme, test.want[i].Lexeme, numberTest, i)
			}

			if actualTokens[i].Line != test.want[i].Line {
				t.Errorf("expected token at line = %d, but token line = %d at position test [index test = %d, index line test = %d]", actualTokens[i].Line, test.want[i].Line, numberTest, i)
			}

			if actualTokens[i].Type == STRING {
				if actualTokens[i].Literal.(string) != test.want[i].Literal.(string) {
					t.Errorf("literal expected %s, but actual literal = %s at position test [index test = %d, index line test = %d]", actualTokens[i].Literal.(string), test.want[i].Literal.(string), numberTest, i)
				}
			} else if actualTokens[i].Type == NUMBER {
				if actualTokens[i].Literal.(float64) != test.want[i].Literal.(float64) {
					t.Errorf("literal expected %f, but actual literal = %f at position test [index test = %d, index line test = %d]", actualTokens[i].Literal.(float64), test.want[i].Literal.(float64), numberTest, i)
				}
			}
		}
	}
}
