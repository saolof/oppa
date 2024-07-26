package main

import (
	"fmt"
	"testing"
)

func TestBinop1(t *testing.T) {
	tokens := SimpleInput("id", "*", "id", "+", "id", "*", "id", "$")
	prec := OpHelper(map[TokenType]int{"+": 1, "*": 2, "id": 1000})
	reductions := map[string]string{"id": "expr", "expr + expr": "expr", "expr * expr": "expr"}
	parser := OpParser{prec: prec, productions: reductions}
	for _, tok := range tokens {
		err := parser.Next(tok)
		if err != nil {
			t.Fail()
			return
		}
	}
	if len(parser.stack) != 2 || parser.stack[1].Terminal.String() != "$" || parser.stack[0].Type != "expr" {
		t.Fail()
		return
	}
	ch1 := parser.stack[0].Children
	fmt.Println(ch1)
	if !ch1[1].IsTerminal || ch1[1].Terminal.TokenType() != "+" {
		t.Fail()
		return
	}
	ch2 := ch1[0].Children
	ch3 := ch1[0].Children
	if !ch2[1].IsTerminal || ch2[1].Terminal.TokenType() != "*" {
		t.Fail()
		return
	}
	if !ch3[1].IsTerminal || ch2[1].Terminal.TokenType() != "*" {
		t.Fail()
		return
	}

}
