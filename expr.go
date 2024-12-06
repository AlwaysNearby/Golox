package golox

type VisitorExpr interface {
	VisitBinary(expr Binary) any
	VisitGrouping(expr Grouping) any
	VisitLiteral(expr Literal) any
	VisitUnary(expr Unary) any
}

type Expr interface {
	Accept(visitor VisitorExpr) any
}

type Binary struct {
	Left Expr
	Operator Token
	Right Expr
}

func (b Binary) Accept(visitor VisitorExpr) any {
	return visitor.VisitBinary(b)
}

type Grouping struct {
	Expression Expr
}

func (g Grouping) Accept(visitor VisitorExpr) any {
	return visitor.VisitGrouping(g)
}

type Literal struct {
	Value interface{}
}

func (l Literal) Accept(visitor VisitorExpr) any {
	return visitor.VisitLiteral(l)
}

type Unary struct {
	Operator Token
	Right Expr
}

func (u Unary) Accept(visitor VisitorExpr) any {
	return visitor.VisitUnary(u)
}
