package main

import "fmt"

type Ast struct {
	Global        map[string]any
	Main          MainFunc
	DeclaredFuncs []Func[any]
}
type MainFunc struct {
	Events []Token
}
type Func[R any] struct {
	Return R
	Events []Token
}

type Token struct {
	Type   int
	Values map[string]string
	Action string
}

func (t *Token) funcPrint() {
	fmt.Println("fn", t.Values["name"], t.Values["params"], t.Values["innards"])
}
func (t *Token) varPrint() {
	fmt.Println("var", t.Values["name"]+":", t.Values["type"], "=", t.Values["value"])
}

var tokenTypeHash = map[int]string{
	0:  "Control",
	1:  "Var Declaration",
	2:  "Function Declaration",
	3:  "Function Call",
	4:  "If Statement",
	5:  "Else Statement",
	6:  "Else If Statement",
	7:  "Return Statement",
	8:  "Var Set",
	9:  "Main Declaration",
	10: "",
}
