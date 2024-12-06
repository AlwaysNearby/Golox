package golox

import "fmt"

type TokenType int

const (
	//single character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	//One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	//Literals.
	IDENTIFIER
	STRING
	NUMBER

	//Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

func (t TokenType) String() string {
	switch t {
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COMMA:
		return ","
	case DOT:
		return "."
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case SEMICOLON:
		return ";"
	case STAR:
		return "*"
	case BANG:
		return "!"
	case BANG_EQUAL:
		return "!="
	case EQUAL:
		return "="
	case EQUAL_EQUAL:
		return "=="
	case LESS:
		return "<"
	case LESS_EQUAL:
		return "<="
	case GREATER:
		return ">"
	case GREATER_EQUAL:
		return ">="
	case SLASH:
		return "/"
	case STRING:
		return "string"
	case NUMBER:
		return "number"
	case IDENTIFIER:
		return "identifier"
	case AND:
		return "and"
	case OR:
		return "or"
	case CLASS:
		return "class"
	case ELSE:
		return "else"
	case IF:
		return "if"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case PRINT:
		return "print"
	case RETURN:
		return "return"
	case SUPER:
		return "super"
	case THIS:
		return "this"
	case VAR:
		return "var"
	case NIL:
		return "nil"
	case FALSE:
		return "false"
	case FUN:
		return "fun"
	case TRUE:
		return "true"
	case EOF:
		return ""
	default:
		panic(fmt.Sprintf("new token added, process it %d", t))
	}
}
