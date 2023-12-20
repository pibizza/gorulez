package main

import (
	"fmt"

	"github.com/pibizza/gorulez/expr"
)

func main() {

	env := make(map[string]expr.Value)

	env["x"] = expr.IntValue(3)

	e1 := expr.IntPlus(expr.IntId("x"), expr.IntValue(4))

	fmt.Println(e1.Eval(env))
}
