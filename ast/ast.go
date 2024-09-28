package ast

import "MonkeyGo/token"

// node from the very start
type Node interface {
	TokenLiteral() string
}

// a statement based on Node
type Statement interface {
	Node
	statementNode()
}

// a expression based on Node
type Expression interface {
	Node
	expressionNode()
}

// a Statement
type Program struct {
	Statements []Statement // any item which impl with Statement can be in the struct
}

// as the root Node of the program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// let statement
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// variable as it said Identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode () {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

