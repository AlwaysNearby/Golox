package main

import (
	"fmt"
	"golox"
)

type AstPrinter struct {
	expr          golox.Expr
	currentIndent int
	indentStep    int
	levelsUsed    map[int]int
	currentLevel  int
}

func (a AstPrinter) Print() {
	if a.expr == nil {
		return
	}
	a.expr.Accept(a)
}

func (a AstPrinter) VisitBinary(expr golox.Binary) any {
	a.parenthesize(expr.Operator.Lexeme, exprsToSlice(expr.Left, expr.Right))
	return nil
}

func (a AstPrinter) VisitGrouping(expr golox.Grouping) any {
	a.parenthesize("group", exprsToSlice(expr.Expression))
	return nil
}

func (a AstPrinter) VisitLiteral(expr golox.Literal) any {
	a.parenthesize(fmt.Sprintf("%v", expr.Value), nil)
	return nil
}

func (a AstPrinter) VisitUnary(expr golox.Unary) any {
	a.parenthesize(expr.Operator.Lexeme, exprsToSlice(expr.Right))
	return nil
}

func exprsToSlice(exprs ...golox.Expr) (slice []golox.Expr) {
	for _, expr := range exprs {
		if expr != nil {
			slice = append(slice, expr)
		}
	}
	return
}

// TODO сделать поддержку любого отступа, так как символ может не достигать sпереноса
func (a AstPrinter) parenthesize(name string, exprs []golox.Expr) {
	spaces := ""
	for i := 0; i < a.currentIndent; i++ {
		if i%a.indentStep == 0 && a.needVisualizeDepth(i/a.indentStep) {
			spaces += " |"
		} else {
			spaces += " "
		}
	}

	var prefix string = ""
	if a.currentLevel > 0 {
		switch a.levelsUsed[a.currentLevel-1] {
		case 1:
			prefix = " └──"
			break
		case 2:
			prefix = " ├──"
			break
		}

		a.levelsUsed[a.currentLevel-1]--
	}

	a.levelsUsed[a.currentLevel] = len(exprs)
	a.currentLevel++
	fmt.Print(spaces + fmt.Sprintf("%s[%s]\n", prefix, name))
	for _, expr := range exprs {
		if expr == nil {
			continue
		}
		a.currentIndent += a.indentStep
		expr.Accept(a)
		a.currentIndent -= a.indentStep
	}

	a.currentLevel--
}

func (a AstPrinter) needVisualizeDepth(depth int) bool {
	return a.levelsUsed[depth] > 0
}

func main() {
	astPrinter := AstPrinter{currentIndent: -4, indentStep: 4, currentLevel: 0, levelsUsed: map[int]int{}}

	astPrinter.expr = golox.Binary{
		Left: golox.Binary{
			Left:     golox.Literal{Value: 1},
			Operator: golox.Token{Type: golox.PLUS, Lexeme: "+", Literal: nil, Line: 1},
			Right:    golox.Literal{Value: 2},
		},
		Operator: golox.Token{Type: golox.STAR, Lexeme: "*", Literal: nil, Line: 1},
		Right: golox.Binary{
			Left:     golox.Literal{Value: 4},
			Operator: golox.Token{Type: golox.PLUS, Lexeme: "-", Literal: nil, Line: 1},
			Right:    golox.Literal{Value: 3},
		},
	}

	astPrinter.Print()
}
