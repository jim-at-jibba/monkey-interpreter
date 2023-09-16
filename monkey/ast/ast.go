package ast

import "github.com/jim-at-jibba/monkey/token"

// let <identifier> = <expression>;
// Three parts to capture let Token, the identifier
// and the expression

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

// Expressions produce a value
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of our AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token // the token.IDENT Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
