package ast

type Node interface {
	TokeLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statemnts) > 0 {
		return p.Statemnets[0].TokenLiteral()
	} else {
		return ""
	}
}
