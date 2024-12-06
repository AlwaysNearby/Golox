package main

import (
	"fmt"
	"golox"
	"strings"
)

type RnpPrinter struct {
	expr golox.Expr
}

func (r RnpPrinter) Print() {
	if r.expr == nil {
		return
	}

	fmt.Println(r.expr.Accept(r))
}

func (r RnpPrinter) VisitBinary(expr golox.Binary) any {
	return r.rpn(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (r RnpPrinter) VisitGrouping(expr golox.Grouping) any {
	return r.rpn("groping", expr.Expression)
}

func (r RnpPrinter) VisitLiteral(expr golox.Literal) any {
	return fmt.Sprintf("%v", expr.Value)
}

func (r RnpPrinter) VisitUnary(expr golox.Unary) any {
	return r.rpn(expr.Operator.Lexeme, expr.Right)
}

func (r RnpPrinter) rpn(lexeme string, exprs ...golox.Expr) string {
	polish := []string{}

	for _, expr := range exprs {
		polish = append(polish, expr.Accept(r).(string))
	}

	polish = append(polish, lexeme)

	return strings.Join(polish, " ")
}

func main() {
	r := RnpPrinter{
		expr: golox.Binary{
			Left: golox.Binary{
				Left:     golox.Literal{Value: 2},
				Operator: golox.Token{Type: golox.PLUS, Lexeme: "+", Literal: nil, Line: 1},
				Right:    golox.Literal{Value: 2},
			},
			Operator: golox.Token{Type: golox.STAR, Lexeme: "*", Literal: nil, Line: 1},
			Right: golox.Binary{
				Left:     golox.Literal{Value: 4},
				Operator: golox.Token{Type: golox.PLUS, Lexeme: "-", Literal: nil, Line: 1},
				Right:    golox.Literal{Value: 3},
			},
		},
	}

	r.Print()
}
