package main

import "fmt"

const (
	EOF   = "$"
	ID    = "id"
	PLUS  = "+"
	TIMES = "*"
)

func main() {
	tokens := []Token{
		SimpleToken{"id", "id"},
		SimpleToken{"+", "+"},
		SimpleToken{"id", "id"},
		SimpleToken{"+", "+"},
		SimpleToken{"id", "id"},
		SimpleToken{"*", "*"},
		SimpleToken{"id", "id"},
		SimpleToken{"$", "$"},
	}
	prec := OpHelper(map[TokenType]int{"+": 1, "*": 2, "id": 1000})
	var prev Token
	prev = SimpleToken{"$", "$"}
	for _, t := range tokens {
		comp, _ := prec(prev.TokenType(), t.TokenType())
		fmt.Print(comp)
		fmt.Print(" ")
		fmt.Print(t)
		fmt.Print(" ")
		prev = t
	}
	fmt.Println("Parsing!")
	parser := OpParser{prec: prec}
	//	parser.Next(SimpleToken{"$", "$"})
	for _, t := range tokens {
		err := parser.Next(t)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(len(parser.stack))
		fmt.Println(parser.stack)
	}

	//fmt.Println(OpParse(tokens, exprparser))

	fmt.Println("hello work!")
}
