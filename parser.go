package main

import "fmt"

type TokenType = string

type Token interface {
	fmt.Stringer
	TokenType() TokenType
}

type TokenCmp int
type TokenPrecedence func(TokenType, TokenType) (TokenCmp, error)

type CSTNode struct {
	IsTerminal bool
	Terminal   Token
	relPrev    TokenCmp
	Children   []CSTNode
}
type OpParser struct {
	stack []CSTNode
	prec  TokenPrecedence
}

const (
	NA = TokenCmp(0)
	LT = TokenCmp(1)
	EQ = TokenCmp(2)
	GT = TokenCmp(3)
)

func (p *OpParser) Next(tok Token) error {
start:
	topTerm := p.topTerminal()
	if topTerm < 0 { // No terminals
		p.appendTerm(tok, LT)
		return nil
	}
	topType := p.stack[topTerm].Terminal.TokenType()
	c, err := p.prec(topType, tok.TokenType())
	if err != nil {
		return err
	}
	if c == EQ || c == LT { // Shift
		p.appendTerm(tok, c)
		return nil
	}
	if c == GT {
		i := p.topTerminalLT()
		if i < 0 {
			p.appendTerm(tok, c)
			return nil
		}
		// Reduce
		piv := i
		if i > 0 && !p.stack[i-1].IsTerminal {
			piv = i - 1
		}
		l := len(p.stack)
		node := CSTNode{}
		node.Children = append(node.Children, p.stack[piv:l]...)
		p.stack = p.stack[0:piv]
		p.stack = append(p.stack, node)
		goto start
	}

	return nil
}

func (p *OpParser) appendTerm(tok Token, c TokenCmp) {
	p.stack = append(p.stack, CSTNode{
		IsTerminal: true,
		Terminal:   tok,
		relPrev:    c,
	})
}

func (p OpParser) topTerminal() int {
	for i := len(p.stack) - 1; i >= 0; i -= 1 {
		if p.stack[i].IsTerminal {
			return i
		}
	}
	return -1
}

func (p OpParser) topTerminalLT() int {
	for i := len(p.stack) - 1; i >= 0; i -= 1 {
		if p.stack[i].IsTerminal && p.stack[i].relPrev == LT {
			return i
		}
	}
	return -1
}

// Utils for testing:

type SimpleToken struct {
	Lit string
	T   TokenType
}

func (t SimpleToken) String() string {
	return t.Lit
}

func (t SimpleToken) TokenType() TokenType {
	return t.T
}

func OpHelper(precs map[TokenType]int) TokenPrecedence {
	return func(t1 TokenType, t2 TokenType) (TokenCmp, error) {
		if t1 == EOF && t2 == EOF {
			return 0, fmt.Errorf("Syntax error!")
		}
		if t1 == EOF {
			return LT, nil
		}
		if t2 == EOF {
			return GT, nil
		}
		p1, ok1 := precs[t1]
		p2, ok2 := precs[t2]
		if !ok1 || !ok2 {
			return EQ, fmt.Errorf("Missing!")
		}
		if p1 < p2 {
			return LT, nil
		}
		return GT, nil
	}
}

func (i TokenCmp) String() string {
	if i == GT {
		return "⋗"
	}
	if i == LT {
		return "⋖"
	}
	if i == EQ {
		return "≐"
	}
	return "?"
}
