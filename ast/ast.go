package ast

import "monkey/token"

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

type LetStatement struct {
	Token token.Token // token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }